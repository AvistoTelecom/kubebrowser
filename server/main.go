package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/brawdunoir/kubebrowser/pkg/signals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/labels"
)

type contextKey string

var static = os.Getenv("KO_DATA_PATH")

const (
	// Viper const
	hostnameKey      string = "hostname"
	podNamespaceKey  string = "pod_namespace"
	sessionSecretKey string = "session_secret"
	devKey           string = "dev"
	logLevelKey      string = "log_level"
	// Context keys
	oauth2ConfigKey     contextKey = "oauth2_config"
	oauth2VerifierKey   contextKey = "oauth2_verifier"
	// Normal const
	callbackRoute string = "/auth/callback"
	defaultPort   string = "8080"
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
	isDev := viper.GetBool(devKey)
	logLevel := viper.GetString(logLevelKey)

	if err := InitLogger(logLevel, isDev); err != nil {
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
	config, verifier, err := newOIDCConfig(ctx, viper.GetString(clientIDKey), viper.GetString(clientSecretKey))
	if err != nil {
		logger.Errorf("Failed to setup OIDC: %s", err)
		os.Exit(1)
	}

	// Populate context
	ctx = context.WithValue(ctx, oauth2ConfigKey, config)
	ctx = context.WithValue(ctx, oauth2VerifierKey, verifier)

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
	kubeconfigs, err := kubecfg.lister.Kubeconfigs(viper.GetString(podNamespaceKey)).List(labels.Everything())
	if err != nil {
		logger.Errorf("Error listing kubeconfigs: %s", err)
		c.String(http.StatusInternalServerError, "Error listing kubeconfigs")
	}

	k, err := preprareKubeconfigs(c, kubeconfigs)
	if err != nil {
		logger.Error(err, "Error preparing kubeconfigs")
		c.String(http.StatusInternalServerError, "Error preparing kubeconfigs")
	}

	c.JSON(http.StatusOK, k)
}

func handleGetMe(c *gin.Context) {
	ec := extractFromContext(c)

	logger.Debug("Entering handleGetMe")
	rawIDToken, _ := extractTokens(ec.session)

	idToken, err := ec.oauth2Verifier.Verify(c.Request.Context(), rawIDToken)
	if err != nil {
		logger.Errorf("Error verifying ID Token: %s", err)
		c.String(http.StatusInternalServerError, "Error verifying ID Token")
	}
	claims := struct {
		Name string
	}{}

	if err := idToken.Claims(&claims); err != nil {
		logger.Errorf("Error extracting claims: %s", err)
		c.String(http.StatusInternalServerError, "Error extracting claims")
	}

	c.JSON(http.StatusOK, claims.Name)
}
