
EC2 (Elastic Compute Cloud):
- It is a general purpose cloud server that can be used to run web applications, crown jobs, etc.
- There are various types of servers available:
    - based on size (small, medium, large)
    - CPU, memory & disk requirements.
    - with / without GPU. Also various GPU options are available.



Spot Instance
- A discounted version of EC2. Cost upto 90% lower than on-demand prices.
- AWS can reclaim (terminate) the instance with 2 minutes notice.
- Ideal for non-realtime, fault safe tasks.



Security Groups:
- Security groups are virtual firewalls.
- They control the inbound & outbound traffic to and from a service.
- E.g.
    - Allow traffic from all IP addresses to 8080 port.
    - Allow SSH only from this IP address.



Auto Scaling:
- AWS Auto Scaling is a service that automatically adjusts the number of instances (EC2) based on the load.
- It scales up the resources under high load and scales down the resources under low load. Eventually optimising the resource cost.
- It scales based on metrics like CPU usage, network traffic or custom metrics.
- Used with services like EC2, ECS, Aurora replicas, DynamoDB, etc.
- It also performs health checks and replaces unhealthy instances.



ECS (Elastic Container Service):
- ECS is a fully managed container orchestration service.
- It manages running, scheduling and scaling of docker containers.
- Just like kubernetes it monitors container health and handles auto recovery.
- ECS is AWS native solution just like kubernetes.



ELB (Elastic Load Balancer):
 - 



Lambda:
 - 



Elastic IP:
 - 



EBS (Elastic Block Storage):
 - 
