The library provided in this package allows users to register and expose, custom as well as generic metrics to the Prometheus instance for monitoring. 
The default path and port at which the metrics would be exposed are "/customMetrics" and "8089" respectively.

The user can create a custom configuration for exposing the metrics by using the NewBuilder() function. 
The metric configurations, like the port and path can be modified using WithPath() and WithPort() functions, provided by the library.
Additionaly, the metric collectors which are to be registered with prometheus can be passed to WithCollectors() function. If no value is provided, the generic Go metrics are exposed. 

The functions in the library perform the following:
1. StartMetrics - Starts the metrics server at default port(8089) and default path ("/customMetrics").
2. RegisterMetrics - Registers the collectors provided by the user to prometheus.
3. GenerateService - Generates service object with provided specifications
4. GenerateRoute - Creates external route for the specified service.
5. CreateOrUpdateService - Updates the service if it already exsists, or creates a new one.
6. CreateOrUpdateRoute - Updates the route if it already exsists, or creates a new one.
7. ConfigureMetrics - Starts the metrics server, creates service and route to the endpoint where the metrics are exposed.
