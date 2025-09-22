## Index
- [EC2 (Elastic Compute Cloud)](#ec2-elastic-compute-cloud)
- [Spot Instance](#spot-instance)
- [Security Groups](#security-groups)
- [AWS Lambda](#aws-lambda)
- [AWS S3](#aws-s3-simple-storage-service)
- [Amazon RDS](#amazon-rds-relational-database-service)
- [Amazon Aurora](#amazon-aurora)
- [Amazon RDS vs Amazon Aurora](#amazon-rds-vs-amazon-aurora)

<br>
<br>
<br>
<br>
<br>

### EC2 (Elastic Compute Cloud):

What is EC2?
- EC2 stands for Elastic Compute Cloud.
- EC2 is a core AWS service that provides general purpose cloud server (instances).
- These servers can be used to run backend services, APIs, crown jobs, etc.
- There are various types of servers available, we can choose the OS, instance size, network, storage, etc.

<br>

Key features:
- We can use **Amazon Machine Image (AMI)** to launch a preconfigured server.
- Use **Security Groups** to control inbound and outbound traffic of the instance.
- Use **Auto Scaling Groups (ASG)** to automatically scale EC2 instances based on demand metrics.
- Use **Elastic Load Balancer (ELB)** to distribute traffic to multiple EC2 instances.
- Use **Elastic Block Store (EBS)** to attach persistent storage volumes to EC2 instance.

<br>
<br>
<br>

### Spot Instance
- A **Discounted version of EC2** that costs upto 90% lower than on-demand prices.
- AWS can reclaim (terminate) the instance with 2 minutes notice.
- Ideal for non-realtime, fault safe tasks.
- When working with spot instances, use **Elastic Block Store (EBS)** to save data on persistent volumes.

<br>
<br>
<br>

### Security Groups:
- Security groups are **Virtual Firewalls**.
- They control the inbound & outbound traffic to and from an AWS service.
- They define rules using:
    - Protocol (TCP/UDP)
    - Port (80/443)
    - IP Address

<br>

- Default security group allows all outbound traffic and no inbound traffic.
- We can attach multiple security groups to an instance.
- E.g.
    - Allow traffic from all IP addresses to 8080 port.
    - Allow SSH only from this IP address.

<br>
<br>
<br>

### AWS Lambda
- It is a **Serverless Compute Service** that runs our code in response to an event.
- It scales automatically based on demand. There is no need to manage servers & infrastructure.
- It is stateless in nature and does not gaurantee data persistance across function calls.

<br>

Key Features:
- Supports multiple languages: Python, Java, JavaScript, etc.
- We have to pay only for the compute time consumed (billing in milliseconds).
- It can be triggered manually or by other AWS services.

<br>

Limits:
- Time limit: **15 minutes (900 seconds)**. Default time limit is 3 seconds.
- Memory limit: **128 MB to 10 GB**. Default memory is 128 MB. More memory means more per millisecond cost.
- CPU limit: CPU limit cannot be set explicitely. It is tied to memory limit.
- Disk limit: **512 MB to 10 GB**. Default is 512 MB.

<br>
<br>
<br>

### AWS S3 (Simple Storage Service)
- S3 stands for **Simple Storage Service**.
- It is used to store objects like images, videos, files, etc.

<br>

Key Features:
- It stores data in **Buckets** (folders) and each file is an **Object**.
- It is highly scalable and fully managed service. We don't have to worry about hitting storage limit.
- It has a "Pay-as-you-go" model that means we only have to pay for the storage we use.
- S3 also support versioning & object encryption.
- It has different storage options depending upon how frequently the data is accessed:
    - **S3 Standard**: General-purpose storage with low latency and high throughput for frequently accessed data.
    - **S3 Infrequent Access**: For infrequently accessed data that needs quick retrieval — lower cost, but retrieval fees apply.
    - **S3 Glacier: Low-cost**. For long-term archiving where data is accessed very rarely. Retrieval times from minutes to hours.
    - **S3 Intelligent-Tiering**: Automatically moves data between access tiers based on usage patterns to optimize costs.

<br>
<br>
<br>

### Amazon RDS (Relational Database Service)
- RDS stands for **Relational Database Service**.
- It is a managed service that takes care of:
    - Database setup
    - Scaling
    - Backup & Snapshots
    - Patching
    - Security
- Supports 6 database engines: MySQL, PostgreSQL, Oracle, SQL Server, MariaDB, AWS Arora.

<br>

Key Features:
- It provides automated backups, patching, and updates.
- It provides read-replicas for horizontal scaling.
- Has Multi-AZ support for high availability.
- Provides support for automated / manual database snapshots.
- It can be easily integrated with **AWS CloudWatch** to analyze key metrics like CPU utilization, Memory usage, I/O activity, etc.

<br>

Multi-AZ:
- **Multi-AZ (Availability Zone)** is a high-availability feature that automatically replicates your RDS database to a standby instance in a different AZ (within the same region).
- It is not for read scaling — it’s for failover and disaster recovery.

<br>
<br>
<br>

### Amazon Aurora
- It is a **Fully Managed Relational Database** service offered by AWS.
- It provides performance of high-end commercial databases (Oracle, SQL Server), but with the simplicity and cost-effictiveness of open source databases (MySQL, PostgreSQL).
- Supports only 2 database engines: MySQL & PostgreSQL.

<br>

Key features:
- Upto 5x faster that MySQL & 3x faster faster than PostgreSQL.
- Supports auto-scaling for read replicas (upto 15).
- Storage automatically scales upto 128 TB.
- High availability: replicates 6 copies of data across 3 Availability zones (AZs).
- Supports multi-master deployment to provide even higher availability.

<br>
<br>
<br>

### Amazon RDS vs Amazon Aurora

| Feature / Criteria           | **Amazon RDS**                                                             | **Amazon Aurora**                                                       |
| ---------------------------- | -------------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| **Database Engines**         | Supports 6 engines: MySQL, PostgreSQL, Oracle, SQL Server, MariaDB, Aurora | Only Aurora MySQL & Aurora PostgreSQL (compatible with open-source)     |
| **Performance**              | Standard performance based on engine                                       | Up to **5x faster than MySQL**, **3x faster than PostgreSQL**           |
| **Storage Scaling**          | Manual or auto (for some engines), but slower                              | **Auto-scales storage** (10 GB to 128 TB) seamlessly                    |
| **Availability**             | Multi-AZ available with failover                                           | **Built-in high availability** (6 copies across 3 AZs), faster failover |
| **Read Replicas**            | Up to 5 (depending on engine)                                              | Up to **15 low-latency replicas**                                       |
| **Serverless Option**        | Only Aurora (not available in standard RDS)                                | **Aurora Serverless v2** supports auto-scaling compute                  |
| **Backups & Snapshots**      | Automated & manual                                                         | Same, with faster recovery times                                        |
| **Pricing**                  | Generally **lower cost**                                                   | Higher cost but better performance and scalability                      |
| **Multi-Master Support**     | Not available                                                              | Available (Aurora Multi-Master – limited)                               |
| **Use of Custom Extensions** | More flexible (especially in PostgreSQL)                                   | Some limitations (fewer extensions)                                     |

<br>

When to Use Amazon RDS
- We need standard performance and lower cost.
- We are using commercial databases like Oracle or SQL Server.
- We don’t need massive scaling or extreme fault tolerance.
- We're running small-to-medium applications or traditional workloads.

<br>

When to Use Amazon Aurora
- You need high performance and high availability.
- You want automatic scalability of storage and compute.
- You need fast failover and high resilience.
- You're building high-throughput or globally distributed applications.
- You want to stay close to MySQL or PostgreSQL compatibility, but with better performance.

<br>
<br>
<br>

### 

<br>
<br>
<br>

Auto Scaling:
- AWS Auto Scaling is a service that automatically adjusts the number of instances (EC2) based on the load.
- It scales up the resources under high load and scales down the resources under low load. Eventually optimising the resource cost.
- It scales based on metrics like CPU usage, network traffic or custom metrics.
- Used with services like EC2, ECS, Aurora replicas, DynamoDB, etc.
- It also performs health checks and replaces unhealthy instances.

<br>
<br>
<br>

ECS (Elastic Container Service):
- ECS is a fully managed container orchestration service.
- It manages running, scheduling and scaling of docker containers.
- Just like kubernetes it monitors container health and handles auto recovery.
- ECS is AWS native solution just like kubernetes.

<br>
<br>
<br>
<br>
<br>

- ELB (Elastic Load Balancer):
- Lambda
- Elastic IP
- AMI (Amazon machine image)
- EBS (Elastic Block Storage)
