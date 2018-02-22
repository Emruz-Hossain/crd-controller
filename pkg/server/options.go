package server

import (
	"io"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"crd-controller/pkg/apis/crd.emruz.com/v1alpha1"
	"fmt"
	"net"
)

const(
	defaultEtcdPathPrefix = "/registry/crd.emruz.com"
)

type CrdServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions

	StdOut io.Writer
	StdErr io.Writer
}

func NewCrdServerOptions(out, errOut io.Writer) *CrdServerOptions {
	opt := &CrdServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(defaultEtcdPathPrefix, Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion)),

		StdOut:             out,
		StdErr:             errOut,
	}
	opt.RecommendedOptions.Etcd = nil

	return opt
}

func (opt CrdServerOptions) Validate(args []string) error {
	return nil
}

func (opt *CrdServerOptions) Complete() error {
	return nil
}

func (opt CrdServerOptions) Config() (*CrdServerConfig, error) {
	// TODO have a "real" external address
	if err := opt.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(Codecs)
	if err := opt.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &CrdServerConfig{
		GenericConfig:    serverConfig,
	}
	return config, nil
}

func (opt CrdServerOptions) Run(stopCh <-chan struct{}) error {
	config, err := opt.Config()
	if err != nil {
		return err
	}

	s, err := config.Complete().New()
	if err != nil {
		return err
	}

	return s.Run(stopCh)
}
