module github.com/openshift/operator-custom-metrics

go 1.12

require (
	github.com/coreos/prometheus-operator v0.31.1
	github.com/go-logr/logr v0.1.0 // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/openshift/api v3.9.0+incompatible
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery v0.0.0-20190313205120-d7deff9243b1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v0.3.3 // indirect
	sigs.k8s.io/controller-runtime v0.1.12
)
