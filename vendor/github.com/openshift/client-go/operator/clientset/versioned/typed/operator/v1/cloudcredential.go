// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// CloudCredentialsGetter has a method to return a CloudCredentialInterface.
// A group's client should implement this interface.
type CloudCredentialsGetter interface {
	CloudCredentials() CloudCredentialInterface
}

// CloudCredentialInterface has methods to work with CloudCredential resources.
type CloudCredentialInterface interface {
	Create(ctx context.Context, cloudCredential *v1.CloudCredential, opts metav1.CreateOptions) (*v1.CloudCredential, error)
	Update(ctx context.Context, cloudCredential *v1.CloudCredential, opts metav1.UpdateOptions) (*v1.CloudCredential, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, cloudCredential *v1.CloudCredential, opts metav1.UpdateOptions) (*v1.CloudCredential, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CloudCredential, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CloudCredentialList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CloudCredential, err error)
	Apply(ctx context.Context, cloudCredential *operatorv1.CloudCredentialApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CloudCredential, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, cloudCredential *operatorv1.CloudCredentialApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CloudCredential, err error)
	CloudCredentialExpansion
}

// cloudCredentials implements CloudCredentialInterface
type cloudCredentials struct {
	*gentype.ClientWithListAndApply[*v1.CloudCredential, *v1.CloudCredentialList, *operatorv1.CloudCredentialApplyConfiguration]
}

// newCloudCredentials returns a CloudCredentials
func newCloudCredentials(c *OperatorV1Client) *cloudCredentials {
	return &cloudCredentials{
		gentype.NewClientWithListAndApply[*v1.CloudCredential, *v1.CloudCredentialList, *operatorv1.CloudCredentialApplyConfiguration](
			"cloudcredentials",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.CloudCredential { return &v1.CloudCredential{} },
			func() *v1.CloudCredentialList { return &v1.CloudCredentialList{} }),
	}
}
