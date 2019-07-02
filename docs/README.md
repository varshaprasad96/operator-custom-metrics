# Operator-custom metrics
This library assists in collecting operator custom-metrics from operators and exposing them to prometheus instance.

## Prerequisites
In order to use the functions provided by this library download the package inside the `vendor` directory of your operator code.
The following command can be used:

```bash
go get https://github.com/openshift/operator-custom-metrics.git
```

## Using operator-custom metrics library

### Description
The library is designed to register the metrics provided by the user to prometheus. The user can provide the **metrics configuration** based on which the metrics are collected.
The following functions are performed by the library:
1. Register custom metrics provided by the user to prometheus.
2. Publish metrics at the endpoint and path based on the metrics configuration.
3. Create a service resource, which selects the pods deployed by the operator based on the labels.
4. Create an external route so that the metrics can be accessed.

The following are the parameters of the metrics configuration (`metricsConfig`) which the user can provide:

- Metrics Path:
This is the endpoint at which the metrics would be exposed. The default metrics endpoint is `/customMetrics`.

- Metrics Port:
This is the port at which the metrics would be published. The default metrics port is `:8089`

- List of collectors:
The user can provide the collectors which are to be registered with prometheus. The library will make a list and register them with prometheus. If no collectors
are passed, the default GoLang metrics are exposed.

- Record metrics function:
This is a user-defined function which describes the process in which the custom metrics are to be collected. This can be passed to the library or can be executed within 
the operator code based on the desired implementation.

## Using the library
The following functions of the library can be called by the user to create a metrics configuration which is to be passed to the library:
1. `NewBuilder()` - Sets the parameters of the metricsConfig Object to default values.
2. `WithPort(port string)` - Updates the default value of port in the metricsConfig object, with the value provided by the user.
3. `WithPath(path string)` - Updates the default value of path in the metricsConfig object, with the value provided by the user. 
4. `WithCollectors(collector prometheus.Collector)` - Creates a list of prometheus collectors which are to be registered.
5. `WithMetricsFunction(recordMetricsFunction convertMetrics)` - User defined function which updates the metric parameters given to the library.
6. `GetConfig()` - Returns the reference to metricsConfig which is built with the configuration set by the user.

## Example Usage
```go
    
import (
	"context"
	"testing"
	"time"
	metricslib "github.com/operator-custom-metrics"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// Metrics endpoint and path which is to be used to expose metrics.
const (
	metricsEndPoint = "8080"
	metricsPath     = "/metrics"
)

// Metric variables to be collected.
var (
	opsProcessed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events Test",
	})
)

// RecordMetrics updates the values of the metrics which are to be collected.
func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

// TestConfigMetrics calls the library functions to build a configuration, and 
// start a metrics server.
func TestConfigMetrics(t *testing.T) {
    prTest := metricslib.NewBuilder().
            WithPort(metricsEndPoint).
            WithPath(metricsPath).
            WithCollectors(opsProcessed).
            WithMetricsFunction(RecordMetrics).
            GetConfig()
        
    // Start metrics server with the exposed metrics.
	if err := metricslib.ConfigureMetrics(context.TODO(), *prTest); err != nil {
		log.Error(err, "Fail")
	}

}

```