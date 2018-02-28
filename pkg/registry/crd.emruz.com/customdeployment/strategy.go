package customdeployment

import (
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"crd-controller/pkg/apis/crd.emruz.com/v1alpha1"
)

func NewStrategy(typer runtime.ObjectTyper) customDeploymentStrategy {
	return customDeploymentStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	apiserver, ok := obj.(*v1alpha1.CustomDeployment)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a Flunder.")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), FlunderToSelectableFields(apiserver), apiserver.Initializers != nil, nil
}

// MatchCustomDeployment is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchCustomDeployment(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// CustomDeploymentToSelectableFields returns a field set that represents the object.
func FlunderToSelectableFields(obj *v1alpha1.CustomDeployment) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type customDeploymentStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (customDeploymentStrategy) NamespaceScoped() bool {
	return true
}

func (customDeploymentStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
}

func (customDeploymentStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (customDeploymentStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (customDeploymentStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (customDeploymentStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (customDeploymentStrategy) Canonicalize(obj runtime.Object) {
}

func (customDeploymentStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
