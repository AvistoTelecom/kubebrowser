package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"slices"

	v1alpha1 "github.com/brawdunoir/kubebrowser/pkg/apis/kubeconfig/v1alpha1"
	"github.com/spf13/viper"
)

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// Returns a subset of initial Kubeconfigs depending on the whitelist in each Kubeconfig and the
// claims (user and groups) in the idToken
func filterKubeConfigs(kubeconfigs []*v1alpha1.Kubeconfig, claims EmailAndGroups) []*v1alpha1.Kubeconfig {
	logger.Debug("Entering filterKubeconfig")
	filtered := make([]*v1alpha1.Kubeconfig, 0, len(kubeconfigs))
	for _, kubeconfig := range kubeconfigs {
		whitelist := kubeconfig.Spec.Whitelist
		logger.Debugw("Processing kubeconfig", "name", kubeconfig.Name, "whitelist", whitelist)

		if whitelist == nil {
			logger.Debugw("Whitelist is empty, adding kubeconfig", "name", kubeconfig.Name)
			filtered = append(filtered, kubeconfig)
			continue
		}

		// Check if user email is in whitelist
		if slices.Contains(whitelist.Users, claims.Email) {
			logger.Debugw("User match found, adding kubeconfig", "name", kubeconfig.Name, "email", claims.Email)
			filtered = append(filtered, kubeconfig)
			continue
		}

		// Check if any group matches the whitelist
		for _, group := range claims.Groups {
			if slices.Contains(whitelist.Groups, group) {
				logger.Debugw("Group match found, adding kubeconfig", "name", kubeconfig.Name, "group", group)
				filtered = append(filtered, kubeconfig)
				break
			}
		}
	}
	return filtered
}

func kubeConfigUser(rawIDToken, refreshToken string) v1alpha1.User {
	return v1alpha1.User{Name: "oidc", User: v1alpha1.UserSpec{
		AuthProvider: v1alpha1.AuthProviderSpec{Name: "oidc", Config: v1alpha1.AuthProviderConfig{
			ClientID:     oauth2Config.ClientID,
			ClientSecret: oauth2Config.ClientSecret,
			IDPIssuerURL: viper.GetString("oauth2_issuer_url"),
			IDToken:      rawIDToken,
			RefreshToken: refreshToken,
		}},
	}}
}

func toKubeConfigSpecs(filteredKubeconfigs []*v1alpha1.Kubeconfig, user v1alpha1.User) []*v1alpha1.KubeconfigSpec {
	copiedKubeconfig := make([]*v1alpha1.KubeconfigSpec, 0, len(filteredKubeconfigs))
	for _, kubeconfig := range filteredKubeconfigs {
		k := kubeconfig.DeepCopy()
		ks := k.Spec
		ks.Whitelist = nil                                      // Remove whitelist information
		ks.Kubeconfig.Users = nil                               // Remove all users
		ks.Kubeconfig.Users = append(ks.Kubeconfig.Users, user) // Put user created before
		ks.Kubeconfig.Contexts = ks.Kubeconfig.Contexts[:1]     // Keep first context only
		ks.Kubeconfig.Contexts[0].Context.User = user.Name      // Put same name as user
		copiedKubeconfig = append(copiedKubeconfig, &ks)
	}
	return copiedKubeconfig
}
