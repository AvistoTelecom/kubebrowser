package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/brawdunoir/kubebrowser/pkg/signals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/labels"
)

type contextKey string

var static = os.Getenv("KO_DATA_PATH")

const (
	loggerKey           contextKey = "logger"
	oauth2ConfigKey     contextKey = "oauth2_config"
	oauth2VerifierKey   contextKey = "oauth2_verifier"
	kubeconfigListerKey contextKey = "kubeconfig_lister"
	callbackRoute       string     = "/auth/callback"
	sessionSecret       string     = "secret"
)

func main() {
	// Set up logger
	l, _ := zap.NewProduction()
	defer l.Sync()
	logger := l.Sugar()

	// Set up signals so we handle the shutdown signal gracefully
	ctx := signals.SetupSignalHandler()

	// Create controller lister for Kubeconfigs CRD
	kubeconfigLister, err := setupKubeconfigLister(ctx)
	if err != nil {
		logger.Error(err, "Cannot setup kubeconfig lister")
		os.Exit(1)
	}

	// Create OIDC related config and verifier
	config, verifier, err := setupOidc(ctx, clientID, clientSecret)
	if err != nil {
		logger.Error(err, "Failed to setup Oidc")
		os.Exit(1)
	}

	// Populate context
	ctx = context.WithValue(ctx, loggerKey, logger)
	ctx = context.WithValue(ctx, oauth2ConfigKey, config)
	ctx = context.WithValue(ctx, oauth2VerifierKey, verifier)
	ctx = context.WithValue(ctx, kubeconfigListerKey, kubeconfigLister)

	// Create session store
	store := memstore.NewStore([]byte(sessionSecret))

	router := gin.New()
	router.Use(sessions.Sessions("kubebrowser_session", store))
	router.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	router.Use(AuthMiddleware)

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.RequestURI
		if path == "/" || strings.HasSuffix(path, ".svg") || strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".css") || strings.HasSuffix(path, ".ico") || strings.HasSuffix(path, ".html") {
			gin.WrapH(http.FileServer(gin.Dir(static, false)))(c)
		} else {
			c.File(static + "/index.html")
		}
	})

	router.GET(callbackRoute, handleOAuth2Callback)
	router.GET("/api/kubeconfigs", handleGetKubeconfigs)
	router.GET("/api/me", handleGetMe)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	// Run the server in a goroutine
	go func() {
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
	ec := extractFromContext(c)

	ec.logger.Debug("Getting kubeconfigs")
	kubeconfigs, err := ec.kubeconfigLister.Kubeconfigs(namespace).List(labels.Everything())
	if err != nil {
		ec.logger.Errorf("Error listing kubeconfigs: %s", err)
		c.String(http.StatusInternalServerError, "Error listing kubeconfigs")
	}

	k, err := preprareKubeconfigs(c, kubeconfigs)
	if err != nil {
		ec.logger.Error(err, "Error preparing kubeconfigs")
		c.String(http.StatusInternalServerError, "Error preparing kubeconfigs")
	}

	c.JSON(http.StatusOK, k)
}

func handleGetMe(c *gin.Context) {
	ec := extractFromContext(c)

	ec.logger.Debug("Entering handleGetMe")
	rawIDToken, _ := extractTokens(ec.session)

	idToken, err := ec.oauth2Verifier.Verify(c.Request.Context(), rawIDToken)
	if err != nil {
		ec.logger.Errorf("Error verifying ID Token: %s", err)
		c.String(http.StatusInternalServerError, "Error verifying ID Token")
	}
	claims := struct {
		Name string
	}{}

	if err := idToken.Claims(&claims); err != nil {
		ec.logger.Errorf("Error extracting claims: %s", err)
		c.String(http.StatusInternalServerError, "Error extracting claims")
	}

	c.JSON(http.StatusOK, claims.Name)
}
