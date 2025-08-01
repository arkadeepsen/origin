// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	http "net/http"

	securityv1 "github.com/openshift/api/security/v1"
	scheme "github.com/openshift/client-go/security/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SecurityV1Interface interface {
	RESTClient() rest.Interface
	PodSecurityPolicyReviewsGetter
	PodSecurityPolicySelfSubjectReviewsGetter
	PodSecurityPolicySubjectReviewsGetter
	RangeAllocationsGetter
	SecurityContextConstraintsGetter
}

// SecurityV1Client is used to interact with features provided by the security.openshift.io group.
type SecurityV1Client struct {
	restClient rest.Interface
}

func (c *SecurityV1Client) PodSecurityPolicyReviews(namespace string) PodSecurityPolicyReviewInterface {
	return newPodSecurityPolicyReviews(c, namespace)
}

func (c *SecurityV1Client) PodSecurityPolicySelfSubjectReviews(namespace string) PodSecurityPolicySelfSubjectReviewInterface {
	return newPodSecurityPolicySelfSubjectReviews(c, namespace)
}

func (c *SecurityV1Client) PodSecurityPolicySubjectReviews(namespace string) PodSecurityPolicySubjectReviewInterface {
	return newPodSecurityPolicySubjectReviews(c, namespace)
}

func (c *SecurityV1Client) RangeAllocations() RangeAllocationInterface {
	return newRangeAllocations(c)
}

func (c *SecurityV1Client) SecurityContextConstraints() SecurityContextConstraintsInterface {
	return newSecurityContextConstraints(c)
}

// NewForConfig creates a new SecurityV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*SecurityV1Client, error) {
	config := *c
	setConfigDefaults(&config)
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new SecurityV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*SecurityV1Client, error) {
	config := *c
	setConfigDefaults(&config)
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &SecurityV1Client{client}, nil
}

// NewForConfigOrDie creates a new SecurityV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SecurityV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SecurityV1Client for the given RESTClient.
func New(c rest.Interface) *SecurityV1Client {
	return &SecurityV1Client{c}
}

func setConfigDefaults(config *rest.Config) {
	gv := securityv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SecurityV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
