package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeconfigList contains a list of Kubeconfig objects
type KubeconfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kubeconfig `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kubeconfig is the Schema for the kubeconfigs API
type Kubeconfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeconfigSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen=true

// KubeconfigData defines the structure of the kubeconfig field
type KubeconfigData struct {
	APIVersion     string    `json:"apiVersion"`
	Kind           string    `json:"kind"`
	Clusters       []Cluster `json:"clusters"`
	Contexts       []Context `json:"contexts,omitempty"`
	CurrentContext string    `json:"current-context,omitempty"`
	Users          []User    `json:"users,omitempty"`
}

// +k8s:deepcopy-gen=true

// KubeconfigSpec defines the desired state of Kubeconfig
type KubeconfigSpec struct {
	Name       string         `json:"name"`
	Kubeconfig KubeconfigData `json:"kubeconfig"`
	Whitelist  *Whitelist     `json:"whitelist,omitempty"`
}

// Cluster represents a Kubernetes cluster entry
type Cluster struct {
	Name    string  `json:"name"`
	Cluster Details `json:"cluster"`
}

// Details holds cluster connection details
type Details struct {
	Server                   string `json:"server"`
	CertificateAuthorityData string `json:"certificate-authority-data,omitempty"`
	InsecureSkipTLSVerify    bool   `json:"insecure-skip-tls-verify,omitempty"`
}

// Context represents a context entry
type Context struct {
	Name    string      `json:"name"`
	Context ContextSpec `json:"context"`
}

// Context represents a user entry
type User struct {
	Name string   `json:"name"`
	User UserSpec `json:"user"`
}

// ContextSpec defines the details of a context
type ContextSpec struct {
	Cluster string `json:"cluster"`
	User    string `json:"user"`
}

// UserSpec defines the details of a user
type UserSpec struct {
	AuthProvider AuthProviderSpec `json:"auth-provider"`
}

// AuthProviderSpec defines the authentication provider details
type AuthProviderSpec struct {
	Name   string             `json:"name"`
	Config AuthProviderConfig `json:"config"`
}

// AuthProviderConfig holds the configuration for the authentication provider
type AuthProviderConfig struct {
	ClientID     string `json:"client-id"`
	ClientSecret string `json:"client-secret"`
	IDToken      string `json:"id-token"`
	IDPIssuerURL string `json:"idp-issuer-url"`
	RefreshToken string `json:"refresh-token"`
}

// +k8s:deepcopy-gen=true

// Whitelist contains allowed users/groups
type Whitelist struct {
	Users  []string `json:"users,omitempty"`
	Groups []string `json:"groups,omitempty"`
}

// Resource returns the GroupResource for the Kubeconfig resource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
