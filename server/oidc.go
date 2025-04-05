package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const (
	// Session keys
	initialRouteKey = "initial_route"
	rawIDTokenKey   = "id_token"
	refreshTokenKey = "refresh_token"
)

var oauth2Config *oauth2.Config
var oauth2Verifier *oidc.IDTokenVerifier

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

func InitOIDC(ctx context.Context) error {
	issuerURL := viper.GetString(issuerURLKey)
	provider, err := oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		return err
	}
	clientID := viper.GetString(clientIDKey)
	clientSecret := viper.GetString(clientSecretKey)
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	oauth2Verifier = provider.Verifier(oidcConfig)

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  viper.GetString(hostnameKey) + callbackRoute,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", oidc.ScopeOfflineAccess},
	}
	return nil
}

func refreshTokens(ctx context.Context, config *oauth2.Config, refreshToken string) (*oauth2.Token, error) {
	tokenSource := config.TokenSource(ctx, &oauth2.Token{RefreshToken: refreshToken})
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func redirectToOIDCLogin(c *gin.Context) {
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
	c.Redirect(http.StatusFound, oauth2Config.AuthCodeURL(state, oidc.Nonce(nonce)))
}

func handleOAuth2Callback(c *gin.Context) {
	logger.Debug("Entering handleOAuth2Callback")

	// Retrieve and validate state
	state, err := c.Cookie("state")
	if err != nil {
		logger.Errorf("State cookie not found: %w", err)
		c.String(http.StatusBadRequest, "State not found")
		return
	}
	if c.Query("state") != state {
		logger.Error("State mismatch", "expected", state, "got", c.Query("state"))
		c.String(http.StatusBadRequest, "State did not match")
		return
	}

	// Exchange code for token
	oauth2Token, err := oauth2Config.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		logger.Errorf("Failed to exchange token: %w", err)
		c.String(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
		return
	}

	// Retrieve and validate nonce
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		logger.Error("No id_token field in oauth2 token")
		c.String(http.StatusInternalServerError, "No id_token field in oauth2 token")
		return
	}
	idToken, err := oauth2Verifier.Verify(c.Request.Context(), rawIDToken)
	if err != nil {
		logger.Errorf("Failed to verify ID Token: %w", err)
		c.String(http.StatusInternalServerError, "Failed to verify ID Token: "+err.Error())
		return
	}

	nonce, err := c.Cookie("nonce")
	if err != nil {
		logger.Errorf("Nonce cookie not found: %w", err)
		c.String(http.StatusBadRequest, "Nonce not found")
		return
	}
	if idToken.Nonce != nonce {
		logger.Error("Nonce mismatch", "expected", nonce, "got", idToken.Nonce)
		c.String(http.StatusBadRequest, "Nonce did not match")
		return
	}

	session := sessions.Default(c)

	session.Set(rawIDTokenKey, rawIDToken)
	session.Set(refreshTokenKey, oauth2Token.RefreshToken)
	err = session.Save()
	if err != nil {
		logger.Errorf("Cannot save session: %w", err)
		c.String(http.StatusInternalServerError, "Cannot save session")
	}
	redirectURI := session.Get(initialRouteKey)
	c.Redirect(http.StatusFound, redirectURI.(string))
}

func AuthMiddleware(c *gin.Context) {
	logger.Debug("Entering AuthMiddleware")

	// Skip authentication for callback route
	if c.Request.URL.Path == callbackRoute {
		logger.Debug("Skip auth because user hitting callback route")
		c.Next()
	}

	session := sessions.Default(c)

	session.Set(initialRouteKey, c.Request.RequestURI)
	err := session.Save()
	if err != nil {
		logger.Errorf("Could not save session: %w", err)
		c.String(http.StatusInternalServerError, "Cannot save session")
	}

	// Retrieve ID token from session
	rawIDToken := session.Get(rawIDTokenKey)
	if rawIDToken == nil {
		logger.Debug("ID token missing, redirecting to login")
		redirectToOIDCLogin(c)
		c.Abort()
		return
	}

	// Verify ID token
	idToken, err := oauth2Verifier.Verify(c.Request.Context(), rawIDToken.(string))
	if err != nil || idToken.Expiry.Before(time.Now()) {
		logger.Info("ID token expired or invalid, attempting to refresh")

		// Retrieve refresh token
		refreshToken := session.Get(refreshTokenKey)
		if refreshToken == nil {
			logger.Info("Refresh token cookie missing, redirecting to login")
			redirectToOIDCLogin(c)
			return
		}

		// Refresh tokens
		newToken, err := refreshTokens(c.Request.Context(), oauth2Config, refreshToken.(string))
		if err != nil {
			logger.Errorf("Failed to refresh token, redirecting to login: %w", err)
			redirectToOIDCLogin(c)
			return
		}

		// Verify new ID token
		_, err = oauth2Verifier.Verify(c.Request.Context(), newToken.Extra("id_token").(string))
		if err != nil {
			logger.Errorf("Failed to verify refreshed ID token, redirecting to login: %w", err)
			redirectToOIDCLogin(c)
			return
		}

		// Update session with new tokens
		session.Set(rawIDTokenKey, newToken.Extra("id_token").(string))
		session.Set(refreshTokenKey, newToken.RefreshToken)
		err = session.Save()
		if err != nil {
			logger.Errorf("Cannot save session: %w", err)
			c.String(http.StatusInternalServerError, "Cannot save session")
		}
	}

	// Token is valid, proceed with the request
	logger.Debug("Token is valid, proceed with the request")
	c.Next()
}
