package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const (
	// Viper keys
	clientIDKey     = "oauth2_client_id"
	clientSecretKey = "oauth2_client_secret"
	issuerURLKey    = "oauth2_issuer_url"
	// Session keys
	initialRouteKey = "initial_route"
	rawIDTokenKey   = "id_token"
	refreshTokenKey = "refresh_token"
)

func setCallbackCookie(c *gin.Context, name, value string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		name,
		value,
		int(time.Hour.Seconds()),              // MaxAge in seconds
		"/",                                   // Path
		strings.Split(c.Request.Host, ":")[0], // Domain
		c.Request.TLS != nil,                  // Secure (set to false for local development)
		true,                                  // HttpOnly
	)
	logger.Debugw("Callback cookie is set", "name", name)
}

func newOIDCConfig(ctx context.Context, clientID string, clientSecret string) (oauth2.Config, *oidc.IDTokenVerifier, error) {
	provider, err := oidc.NewProvider(ctx, viper.GetString(issuerURLKey))
	if err != nil {
		return oauth2.Config{}, nil, err
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  viper.GetString(hostnameKey) + callbackRoute,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", oidc.ScopeOfflineAccess},
	}
	return config, verifier, nil
}

func refreshTokens(ctx context.Context, config oauth2.Config, refreshToken string) (*oauth2.Token, error) {
	tokenSource := config.TokenSource(ctx, &oauth2.Token{RefreshToken: refreshToken})
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func redirectToOIDCLogin(c *gin.Context) {
	ec := extractFromContext(c)
	logger.Debug("Entering redirectToOIDCLogin")

	state, err := randString(16)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal error")
		c.Abort()
		return
	}
	nonce, err := randString(16)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal error")
		c.Abort()
		return
	}
	setCallbackCookie(c, "state", state)
	setCallbackCookie(c, "nonce", nonce)

	// Redirect to OIDC login
	c.Redirect(http.StatusFound, ec.oauth2Config.AuthCodeURL(state, oidc.Nonce(nonce)))
}

func handleOAuth2Callback(c *gin.Context) {
	ec := extractFromContext(c)

	logger.Debug("Entering handleOAuth2Callback")

	// Retrieve and validate state
	state, err := c.Cookie("state")
	if err != nil {
		logger.Error(err, "State cookie not found")
		c.String(http.StatusBadRequest, "State not found")
		return
	}
	if c.Query("state") != state {
		logger.Error(nil, "State mismatch", "expected", state, "got", c.Query("state"))
		c.String(http.StatusBadRequest, "State did not match")
		return
	}

	// Exchange code for token
	oauth2Token, err := ec.oauth2Config.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		logger.Error(err, "Failed to exchange token")
		c.String(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
		return
	}

	// Retrieve and validate nonce
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		logger.Error(nil, "No id_token field in oauth2 token")
		c.String(http.StatusInternalServerError, "No id_token field in oauth2 token")
		return
	}
	idToken, err := ec.oauth2Verifier.Verify(c.Request.Context(), rawIDToken)
	if err != nil {
		logger.Error(err, "Failed to verify ID Token")
		c.String(http.StatusInternalServerError, "Failed to verify ID Token: "+err.Error())
		return
	}

	nonce, err := c.Cookie("nonce")
	if err != nil {
		logger.Error(err, "Nonce cookie not found")
		c.String(http.StatusBadRequest, "Nonce not found")
		return
	}
	if idToken.Nonce != nonce {
		logger.Error(nil, "Nonce mismatch", "expected", nonce, "got", idToken.Nonce)
		c.String(http.StatusBadRequest, "Nonce did not match")
		return
	}

	ec.session.Set(rawIDTokenKey, rawIDToken)
	ec.session.Set(refreshTokenKey, oauth2Token.RefreshToken)
	err = ec.session.Save()
	if err != nil {
		logger.Error(err, "Cannot save session")
		c.String(http.StatusInternalServerError, "Cannot save session")
	}
	redirectURI := ec.session.Get(initialRouteKey)
	c.Redirect(http.StatusFound, redirectURI.(string))
}

func AuthMiddleware(c *gin.Context) {
	ec := extractFromContext(c)
	logger.Debug("Entering AuthMiddleware")

	// Skip authentication for callback route
	if c.Request.URL.Path == callbackRoute {
		logger.Debug("Skip auth because user hitting callback route")
		c.Next()
	}

	ec.session.Set(initialRouteKey, c.Request.RequestURI)
	err := ec.session.Save()
	if err != nil {
		logger.Error(err, "Cannot save session")
		c.String(http.StatusInternalServerError, "Cannot save session")
	}

	// Retrieve ID token from session
	rawIDToken := ec.session.Get(rawIDTokenKey)
	if rawIDToken == nil {
		logger.Debug("ID token missing, redirecting to login")
		redirectToOIDCLogin(c)
		c.Abort()
		return
	}

	// Verify ID token
	idToken, err := ec.oauth2Verifier.Verify(c.Request.Context(), rawIDToken.(string))
	if err != nil || idToken.Expiry.Before(time.Now()) {
		logger.Info("ID token expired or invalid, attempting to refresh")

		// Retrieve refresh token
		refreshToken := ec.session.Get(refreshTokenKey)
		if refreshToken == nil {
			logger.Info("Refresh token cookie missing, redirecting to login")
			redirectToOIDCLogin(c)
			return
		}

		// Refresh tokens
		newToken, err := refreshTokens(c.Request.Context(), ec.oauth2Config, refreshToken.(string))
		if err != nil {
			logger.Error(err, "Failed to refresh token, redirecting to login")
			redirectToOIDCLogin(c)
			return
		}

		// Verify new ID token
		_, err = ec.oauth2Verifier.Verify(c.Request.Context(), newToken.Extra("id_token").(string))
		if err != nil {
			logger.Error(err, "Failed to verify refreshed ID token, redirecting to login")
			redirectToOIDCLogin(c)
			return
		}

		// Update session with new tokens
		ec.session.Set(rawIDTokenKey, newToken.Extra("id_token").(string))
		ec.session.Set(refreshTokenKey, newToken.RefreshToken)
		err = ec.session.Save()
		if err != nil {
			logger.Error(err, "Cannot save session")
			c.String(http.StatusInternalServerError, "Cannot save session")
		}
	}

	// Token is valid, proceed with the request
	logger.Debug("Token is valid, proceed with the request")
	c.Next()
}
