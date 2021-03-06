package server

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"crd-controller/pkg/apis/crd.emruz.com/v1alpha1"
	"crd-controller/pkg/apis/crd.emruz.com"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apiserver/pkg/registry/rest"
	cdpregistry "crd-controller/pkg/registry"
	cdpstorage "crd-controller/pkg/registry/crd.emruz.com/customdeployment"
)

var (
	registry = registered.NewOrDie("")
	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

type CrdServerConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
}

// CrdServer contains state for a Kubernetes cluster master/api server.
type CrdServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *CrdServerConfig) Complete() CompletedConfig {
	completedCfg := completedConfig{
		c.GenericConfig.Complete(),
	}

	completedCfg.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&completedCfg}
}

// New returns a new instance of CrdServer from the given config.
func (c *completedConfig) New() (*CrdServer, error) {
	genricServer, err := c.GenericConfig.New("crd-server", genericapiserver.EmptyDelegate)
	if err != nil {
		return nil, err
	}

	s := &CrdServer{
		GenericAPIServer: genricServer,
	}
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(crd_emruz_com.GroupName, registry, Scheme, metav1.ParameterCodec, Codecs)
	apiGroupInfo.GroupMeta.GroupVersion = v1alpha1.SchemeGroupVersion
	v1alpha1storage := map[string]rest.Storage{}
	v1alpha1storage["customdeployment"] = cdpregistry.RESTInPeace(cdpstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))


	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *CrdServer) Run(stopCh <-chan struct{}) error {
	return s.GenericAPIServer.PrepareRun().Run(stopCh)
}
