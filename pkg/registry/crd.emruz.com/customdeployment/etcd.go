package customdeployment

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"crd-controller/pkg/apis/crd.emruz.com/v1alpha1"
	"crd-controller/pkg/registry"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &v1alpha1.CustomDeployment{} },
		NewListFunc:              func() runtime.Object { return &v1alpha1.CustomDeploymentList{} },
		PredicateFunc:            MatchCustomDeployment,
		DefaultQualifiedResource: v1alpha1.Resource("customdeployment"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{store}, nil
}

