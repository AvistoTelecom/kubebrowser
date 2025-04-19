package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/AvistoTelecom/kubebrowser/pkg/signals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/labels"
)

var static = os.Getenv("KO_DATA_PATH")

// Viper keys
const (
	hostnameKey      = "hostname"
	podNamespaceKey  = "pod_namespace"
	sessionSecretKey = "session_secret"
	devKey           = "dev"
	logLevelKey      = "log_level"
	clientIDKey      = "oauth2_client_id"
	clientSecretKey  = "oauth2_client_secret"
	issuerURLKey     = "oauth2_issuer_url"
)

const (
	callbackRoute = "/auth/callback"
	defaultPort   = "8080"
)

func init() {
	viper.SetEnvPrefix("kubebrowser")
	viper.AutomaticEnv()
	viper.SetDefault(hostnameKey, "http://localhost:"+defaultPort)
	viper.SetDefault(sessionSecretKey, "changeme")
	viper.SetDefault(devKey, false)
	viper.SetDefault(logLevelKey, "INFO")
}

func main() {
	if err := InitLogger(); err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Set up signals so we handle the shutdown signal gracefully
	ctx := signals.SetupSignalHandler()

	// Create controller lister for Kubeconfigs CRD
	if err := kubecfg.Init(ctx); err != nil {
		logger.Errorf("Cannot setup kubeconfig lister: %s", err)
		os.Exit(1)
	}

	// Create OIDC related config and verifier
	err := InitOIDC(ctx)
	if err != nil {
		logger.Errorf("Failed to setup OIDC: %s", err)
		os.Exit(1)
	}

	// Create session store
	store := memstore.NewStore([]byte(viper.GetString(sessionSecretKey)))

	router := gin.New()

	router.Use(sessions.Sessions("kubebrowser_session", store))
	router.Use(ginzap.GinzapWithConfig(logger.Desugar(), &ginzap.Config{
		TimeFormat:      time.RFC3339,
		UTC:             true,
		SkipPaths:       []string{"/healthz", callbackRoute},
		SkipPathRegexps: []*regexp.Regexp{regexp.MustCompile(`^/home.*`)},
	}))
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	authorized := router.Group("/", AuthMiddleware)
	authorized.StaticFS("/home", http.Dir(static))
	authorized.GET(callbackRoute, handleOAuth2Callback)
	authorized.GET("/api/kubeconfigs", handleGetKubeconfigs)
	authorized.GET("/api/me", handleGetMe)

	srv := &http.Server{
		Addr:    ":" + defaultPort,
		Handler: router,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	// Run the server in a goroutine
	go func() {
		logger.Warn("Start to listen and serve")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: %s\n", err)
		}
	}()

	// Wait for the shutdown signal
	<-ctx.Done()

	// Gracefully shut down the server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error(err, "Server forced to shutdown")
	}

	logger.Warn("Server exiting")
}

func handleGetKubeconfigs(c *gin.Context) {
	logger.Debug("Getting kubeconfig")

	session := sessions.Default(c)
	rawIDToken := session.Get(rawIDTokenKey).(string)
	refreshToken := session.Get(refreshTokenKey).(string)

	// NOTE: verification has been done in AuthMiddleware already
	idToken, err := oauth2Verifier.Verify(c.Request.Context(), rawIDToken)
	if err != nil {
		logger.Error(err, "Error preparing kubeconfigs")
		c.String(http.StatusInternalServerError, "Error preparing kubeconfigs")
		return
	}

	var claims EmailAndGroups
	if err := idToken.Claims(&claims); err != nil {
		logger.Error(err, "Error preparing kubeconfigs")
		c.String(http.StatusInternalServerError, "Error preparing kubeconfigs")
		return
	}
	logger.Debugw("Extracted claims", "claims", claims)

	logger.Debug("Getting list of all kube configs")
	configs, err := kubecfg.lister.Kubeconfigs(viper.GetString(podNamespaceKey)).List(labels.Everything())
	if err != nil {
		logger.Errorf("Error listing kubeconfigs: %s", err)
		c.String(http.StatusInternalServerError, "Error listing kubeconfigs")
	}

	filtered := filterKubeConfigs(configs, claims)
	user := kubeConfigUser(rawIDToken, refreshToken)
	specs := toKubeConfigSpecs(filtered, user)

	c.JSON(http.StatusOK, specs)
}

func handleGetMe(c *gin.Context) {
	logger.Debug("Entering handleGetMe")

	session := sessions.Default(c)
	rawIDToken := session.Get(rawIDTokenKey).(string)

	// NOTE: verification has been done in AuthMiddleware already
	idToken, err := oauth2Verifier.Verify(c.Request.Context(), rawIDToken)
	if err != nil {
		logger.Errorf("Error verifying ID Token: %s", err)
		c.String(http.StatusInternalServerError, "Error verifying ID Token")
	}

	var claims NameOnly
	if err := idToken.Claims(&claims); err != nil {
		logger.Errorf("Error extracting claims: %s", err)
		c.String(http.StatusInternalServerError, "Error extracting claims")
	}

	c.JSON(http.StatusOK, claims.Name)
}
