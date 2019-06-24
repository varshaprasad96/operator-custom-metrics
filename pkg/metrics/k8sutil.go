package metrics

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	discovery "k8s.io/client-go/discovery"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// KubeConfigEnvVar defines the env variable KUBECONFIG which
	// contains the kubeconfig file path.
	KubeConfigEnvVar = "KUBECONFIG"

	// WatchNamespaceEnvVar is the constant for env variable WATCH_NAMESPACE
	// which is the namespace where the watch activity happens.
	// this value is empty if the operator is running with clusterScope.
	WatchNamespaceEnvVar = "WATCH_NAMESPACE"

	// OperatorNameEnvVar is the constant for env variable OPERATOR_NAME
	// which is the name of the current operator
	OperatorNameEnvVar = "OPERATOR_NAME"

	// PodNameEnvVar is the constant for env variable POD_NAME
	// which is the name of the current pod.
	PodNameEnvVar = "POD_NAME"
)

// GetWatchNamespace returns the namespace the operator should be watching for changes
func GetWatchNamespace() (string, error) {
	ns, found := os.LookupEnv(WatchNamespaceEnvVar)
	if !found {
		return "", fmt.Errorf("%s must be set", WatchNamespaceEnvVar)
	}
	return ns, nil
}

// errNoNS indicates that a namespace could not be found for the current
// environment
var ErrNoNamespace = fmt.Errorf("namespace not found for current environment")

// GetOperatorNamespace returns the namespace the operator should be running in.
func GetOperatorNamespace() (string, error) {
	nsBytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		if os.IsNotExist(err) {
			return "", ErrNoNamespace
		}
		return "", err
	}
	ns := strings.TrimSpace(string(nsBytes))
	log.V(1).Info("Found namespace", "Namespace", ns)
	return ns, nil
}

// GetOperatorName return the operator name
func GetOperatorName() (string, error) {
	operatorName, found := os.LookupEnv(OperatorNameEnvVar)
	if !found {
		return "", fmt.Errorf("%s must be set", OperatorNameEnvVar)
	}
	if len(operatorName) == 0 {
		return "", fmt.Errorf("%s must not be empty", OperatorNameEnvVar)
	}
	return operatorName, nil
}

// ResourceExists returns true if the given resource kind exists
// in the given api groupversion
func ResourceExists(dc discovery.DiscoveryInterface, apiGroupVersion, kind string) (bool, error) {
	apiLists, err := dc.ServerResources()
	if err != nil {
		return false, err
	}
	for _, apiList := range apiLists {
		if apiList.GroupVersion == apiGroupVersion {
			for _, r := range apiList.APIResources {
				if r.Kind == kind {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

// GetPod returns a Pod object that corresponds to the pod in which the code
// is currently running.
// It expects the environment variable POD_NAME to be set by the downwards API.
func GetPod(ctx context.Context, client crclient.Client, ns string) (*corev1.Pod, error) {
	podName := os.Getenv(PodNameEnvVar)
	if podName == "" {
		return nil, fmt.Errorf("required env %s not set, please configure downward API", PodNameEnvVar)
	}

	log.V(1).Info("Found podname", "Pod.Name", podName)

	pod := &corev1.Pod{}
	key := crclient.ObjectKey{Namespace: ns, Name: podName}
	err := client.Get(ctx, key, pod)
	if err != nil {
		log.Error(err, "Failed to get Pod", "Pod.Namespace", ns, "Pod.Name", podName)
		return nil, err
	}

	// .Get() clears the APIVersion and Kind,
	// so we need to set them before returning the object.
	pod.TypeMeta.APIVersion = "v1"
	pod.TypeMeta.Kind = "Pod"

	log.V(1).Info("Found Pod", "Pod.Namespace", ns, "Pod.Name", pod.Name)

	return pod, nil
}
