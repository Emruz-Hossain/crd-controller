package crd_emruz_com

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	GroupName = "crd.emruz.com"
)

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
