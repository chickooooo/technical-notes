## Index
- [Service Discovery](#service-discovery)

<br>
<br>
<br>

### Service Discovery

- It helps services find and talk to each other.
- It is the process of automatically locating and connecting to services in a distributed system.
- This elimintes the need for hard coding the IP address in other services.
- When a new service starts, it registers itself in the Service Registry.
- When using Kubernetes, it automatically registers the pod behind the service.
- No separate registration is required, requests are automatically load balanced to healthy pods.
