// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetK8sNamespaceName sets provided value as "k8s.namespace.name" attribute.
func (rb *ResourceBuilder) SetK8sNamespaceName(val string) {
	if rb.config.K8sNamespaceName.Enabled {
		rb.res.Attributes().PutStr("k8s.namespace.name", val)
	}
}

// SetK8sReplicasetName sets provided value as "k8s.replicaset.name" attribute.
func (rb *ResourceBuilder) SetK8sReplicasetName(val string) {
	if rb.config.K8sReplicasetName.Enabled {
		rb.res.Attributes().PutStr("k8s.replicaset.name", val)
	}
}

// SetK8sReplicasetUID sets provided value as "k8s.replicaset.uid" attribute.
func (rb *ResourceBuilder) SetK8sReplicasetUID(val string) {
	if rb.config.K8sReplicasetUID.Enabled {
		rb.res.Attributes().PutStr("k8s.replicaset.uid", val)
	}
}

// SetOpencensusResourcetype sets provided value as "opencensus.resourcetype" attribute.
func (rb *ResourceBuilder) SetOpencensusResourcetype(val string) {
	if rb.config.OpencensusResourcetype.Enabled {
		rb.res.Attributes().PutStr("opencensus.resourcetype", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
