# Developer Associate

(Notes from Udemy course and other sources)

- https://www.udemy.com/course/aws-certified-developer-associate-dva-c01
- https://courses.datacumulus.com/certified-developer-k92

## AWS Regions and Availability Zones (AZ)

- AWS has _regions_ around the world, eg `us-east-1`
- Each region has _availability zones_, eg `us-east-1a`, `us-east-1b`
- Each availability zone is a distinct, physical data centre in the region
- AWS consoles are scoped to a region, _except_ IAM and S3
- AWS Region Sydney: `ap-southeast-2` has three availbility zones:
    - `ap-southeast-1a`
    - `ap-southeast-1b`
    - `ap-southeast-1c`

## IAM - Identity Access Management

- IAM has a GLOBAL scope
- IAM is at the center of AWS management for security and access:
    - Users
    - Groups
    - Roles
    - Policies
- Root account should not be used for day-to-day management
- Should use _least privilege principle_

![img](../cloud-practitioner/IAM-overview.png)

**IAM Federation** is available for large enterprises to integrate with their
own repository of users using SAML (Active Directory).

Important tips:

- One IAM _User_ per person
- One IAM _Role_ per application
- IAM creds should never be shared
- IAM creds should _never_ be committed or even written into code
- Root account and IAM creds should not be used except for initial setup

## EC2

_Elastic Compute Cloud_

- Encompasses:
    - Virtual machines (EC2)
    - Virtual drives (EBS)
    - Distributing load (ELB)
    - Scaling services (ASG)

- Launching an EC2 Instance:
    - Select AMI (Amazon Machine Image)
    - Configure options, including security group
    - Create and download security key
    - Launch

- Connect:
    - Change `.pem` file permissions from `0644` to `0400`
    - `$ ssh -i ~/path/to-my.pem ec2-user@public-ip`

### Security Groups

- Firewall for EC2 instances
- Regulates access to ports to / from IP addresses / ranges
- Can be attached to multiple instances
- An instance can belong to multiple security groups
- Locked down to a region / VPC combination
- Is _outside_ EC2
- TIP: Maintain one separate security group for SSH access
- _Application timeout_ is generally security group issue
- _Connection refused_ is application level issue, not security group
- By default all inbound traffic is blocked
- By default all outbound traffic is allowed
- Security groups can reference other security groups

### Elastic IPs

- Restarting EC2 instance can change its public IP (private IP remains
  unchanged)
- Elastic IPs are reserved public IPv4 addresses which can be attached to an EC2
  instance
- Can only have 5 elastic IPs per account - can request more
- TIP: Avoid elastic IP - bad architecture. Prefer random IP with DNS assignment
- Best pattern is to use a load balancer

### EC2 User Data

- Bootstrap EC2 instance using _EC2 User Data_ script
- Script runs once only on the instance's first start up
- Used to automate tasks such as installing updates, software and other tasks
- Run with root privileges
- Is on step 3 when configuring EC2 instance:

![ec2 user data](./ec2_user_data.png)

Example User Data script:

```shell script
#!/bin/bash
yum update -y
yum install -y httpd.x86_64
systemctl start httpd.service
systemctl enable httpd.service
echo "Hello from $(hostname -f)" > /var/www/html/index.html
```

### EC2 Instance Launch Types

- On Demand Instances: short workload, predictable pricing
- Reserved: MINIMUM 1 year
    - Reserved Instances: long workloads with same instance type
    - Convertible Reserved Instances: long workloads with flexible instance type
    - Scheduled Reserved Instances: eg every thursday from 3-6pm for one year
- Spot Instances: short workloads, cheap, can lose instances, less reliable
- Dedicated Instances: No other customers will share your hardware
- Dedicated Hosts: entire physical server, control instance placement

#### EC2 On Demand

- Pay per second after first minute
- Highest cost, no upfront payment
- No long-term commitment
- Recommended for: short-term, uninterrupted workloads, unpredictable app
  behaviour

#### EC2 Reserved Instance

- Reservation period can be 1 or 3 years
- Pay upfront with long-term commitment
- Up to 75% discount
- Reserve a specific instance type
- Recommended for steady state usage applications, eg database
- **Convertible Reserved Instances**
    - Can change EC2 instance type
    - Up to 54% discount
- **Scheduled Reserved Instances**
    - Launch in a reserved time window
    - When you require a known fraction of a month, week, day

#### EC2 Spot Instances

- Most cost-efficient instances - up to 90% discount
- Instances can be lost at any time if your **max price** is less than current
  spot price
- _Only useful for workloads that are resilient to failure_, eg:
    - Batch jobs
    - Data analysis
    - Image processing
- Single spot instance can be requested when launching an instance, on the _
  configure instance_ page
- Can also request a _fleet_ of spot instance via _Spot Request_ on left menu

A way to combine instances would be to have a reserved instance for baseline
capacity, plus on demand instance(s) and / or spot instances for peaks,
depending on if workload resilience is required.

#### EC2 Dedicated Instances

- Instances running on hardware that is dedicated to your account
- May share hardware with other instances in the _same_ account
- No control over instance placement (can move hardware after stop / start?)

#### EC2 Dedicated Hosts

- Physical, dedicated EC2 server
- Full control of instance placement
- Visibility into the underlying sockets, physical cores of the hardware
- 3 year period reservation
- Expensive
- Used for software that has complicated / restricted _Bring Your Own Licence_ (
  BYOL) model or regulatory requirements

### Elastic Network Interfaces (ENI)

- Logical component in a VPC that represents a virtual network card
- Bound to a specific AZ
- An ENI can have the following attributes:
    - Primary private IPv4, one or more secondary IPv4
    - One Elastic IP (IPv4) per private IPv4
    - One public IPv4
    - One or more security groups
    - A MAC address
- ENI can be created independently and can be attached / moved to EC2 instances
  for failover

![eni](./eni.png)

- Access ENIs via left menu on EC2 page:

![eni](./eni_link1.png)

- Or via `eth0` link on an EC2 instance:

![eni](./eni_link2.png)

### EC2 Pricing

- Priced per hour, depending on:
    - Region
    - Instance type
    - Launch type (On-Demand, Spot etc)
    - OS
- Billed by the second with a min of 60 seconds
- Additional charges for storage, data transfer, elastic IP, load balancing etc
- **YOU DO NOT PAY IF THE INSTANCE IS STOPPED**
- https://aws.amazon.com/ec2/pricing/on-demand/

### Custom AMI

- Can create custom AMI for more efficient deployments
- **AMI ARE BUILT FOR A SPECIFIC AWS REGION**

### EC2 Instance Type

- Five distinct characteristics: RAM, CPU, I/O, Network, GPU (Graphical
  Processing Unit)
- Over 50 types: https://aws.amazon.com/ec2/instance-types/
- Comparison tool: https://ec2instances.info/
- Main categories:
    - R/C/P/G/H/X/I/F/Z/CR are specialised for RAM, CPU, I/O, Network, GPU
    - M instances are _balanced_ combination of all
    - T2/T3 instances are _burstable_
        - Can burst when load becomes high, this uses _burst credits_
        - Once credits are used reverts back to standard
        - Credits are re-accumulated over time
    - T2 Unlimited allows unlimited bursts but pay for extra credit

### EC2 Checklist

- SSH into EC2 instance, change `.pem` permissions
- Use security groups
- Know difference between private, public and elastic IP
- Use User Data to customise instance at boot time
- Know that can build custom AMI to enhance your OS
- EC2 instances are billed by the second and can be easily created and discarded

---

## High Availability and Scalability for EC2

- Vertical Scaling - instance size up / down
- Horizontal Scaling - number of instances increase / decrease (scale out / in)
    - Auto scaling group
    - Load balancer
- High Availability
    - Auto scaling group multi AZ
    - Load balancer multi AZ

### ELB (Elastic Load Balancer)

- Servers that fwd internet traffic to multiple backend instances
- Spreads load across multiple downstream instances
- Exposes a single point of access to the service
- Seamlessly handles failure of downstream instances
- Does regular health checks of instances
- Provides SSL termination
- Enforce stickiness with cookies
- High availability across zones
- Separate public traffic from private traffic

Advantages:

- Completely managed by AWS and guaranteed to work
- A lot less effort to configure
- Integrated with a lot of existing AWS services

#### ELB Health Checks

- Crucial for load balancers to know if instances are able to reply to requests
- Health check done on a port and route, eg /health:80
- Response 200 OK is healthy
- Happen every 5 seconds (configurable)

#### ELB Types

- **Classic Load Balancer** - CLB (v1 - old generation)
    - Supports HTTP & HTTPS (layer 7) and TCP (layer 4)
    - Health checks are TCP or HTTP-based
    - Fixed hostname `xxx.region.elb.amazonaws.com`, but _not_ a fixed IP

- **Application Load Balancer** - ALB (v2 - new generation)
    - supports HTTP, HTTP/2, HTTPS, WebSocket (Layer 7 only)
    - load balancing to multiple http applications across machines (_target
      groups_)
    - load balancing to multiple applications on the same machine (eg
      containers)
    - supports redirects (eg HTTP -> HTTPS)
    - supports routing tables to different target groups base on:
        - URL paths
        - host names
        - query strings and headers
    - Can route to multiple target groups
    - Health checks are at the target group level
    - Target groups can include:
        - EC2 Instances - can be managed by an Auto Scaling Group - ECS Tasks -
          Lambda functions - HTTP request translated to JSON event - IP
          addresses - must be private
    - Great for micro-services and container-based applications - Docker and
      Amazon ECS
    - Supports port mapping to redirect to dynamic ports in ECS
    - A single ALB can replace need to multiple CLBs
    - ALB has a fixed host name `xxx.region.elb.amazonaws.com`, but _not_ a
      fixed IP
    - App servers don't see the client IP directly, get the following headers
      added to the request":
        - `X-Fowarded-For` - client IP
        - `X-Forwarded-Port` - request port number
        - `X-Forwarded-Proto` - request protocol

- **Network Load Balancer** - NLB (v2 - new generation)
    - supports TCP, TLS (Secure TCP) & UDP (Layer 4 - lower level)
    - Forwards TCP / UDP traffic to instances
    - Much higher performance (lower latency ~ 100 ms vs ~400 ms for ALB)
    - Is not part of a security group and traffic is passed through as is - so
      security group containing target instances specifies where traffic is
      allowed from
    - Handles millions of requests per second
    - Has _one static IP per AZ_ and supports assigning Elastic IP
    - Not included in FREE tier

Can set up **internal** (private) load balancers or **external** (public) load
balancers.

#### Load Balancer Security Groups

Typically set up all traffic allowed to the load balancer, and traffic to EC2
instance only from ELB:

![ELB Security](./elb_security_groups.png)

#### Load Balancer Stickiness

- CLB and ALB only
- Maintains client connection to same target using a cookie
- Cookie has a controllable expiration date
- Use case eg, to maintain session data
- May create a load imbalance to EC2 targets
- Enabled at the target group level (ALB)

#### Cross-Zone Load Balancing

- Normally, each LB distributes load to instances in the same AZ
- This ensures LB in each AZ distributed requests evenly amongst all registered
  instances in each AZ
- Classic Load Balancer - disabled by default, no additional charge
- Application Load Balancer - always on, no additional charges
- Network Load Balancer - disabled by default, additional charges

#### Load Balancer Checklist

- ELB can scale but not instantaneously - can contact AWS for 'warm up'
- Troubleshooting:
    - 4xx errors are client-induced
    - 5xx errors and application-induced
    - Load Balancer 503 means LB is _at capacity_ or _no registered target_
    - If LB can't connect to app, _check security groups_
- Monitoring:
    - ELB access logs will log _all_ requests so can debug at request level
    - CloudWatch Metrics provides aggregate statistics

#### SSL / TLS

- SSL cert allows traffic between load balancer and clients to be encrypted -
  _in-flight_ encryption.
- SSL - Secure Sockets Layer
- TLS - Transport Layer Security is newer version
- TLS certs are mainly used these days but are still referred to as _SSL_ certs.
- Public SSL certs are issued by a Certificate Authority (CA), eg LetsEncrypt
- SSL certs have an expiration date that you set, must be renewed regularly.
- Traffic from client to LB is HTTPS, then from LB to server is HTTP but is on
  private VPC so is secure.
- The LB uses an X.509 certificate (SSL/TLS certificate)
- Can manage certs in ACM (AWS Certificate Manager)
- Can also create and upload own certs
- HTTPS listener:
    - Must specify a default cert
    - Can add optional list of certs to support multiple domains
    - Client can use SNI (Server Name Indication) to specify the hostname they
      reach
    - Can also set specific security policy to support older versions of SSL /
      TLS (legacy clients)

##### SNI - Server Name Indication

- Solves problem of multiple SSL certs, ie for multiple web sites, on a single
  server.
- Problem was that TLS handshake did not include target hostname which means
  server had no way of knowing the correct cert.
- Newer protocol that requires the client to indicate the hostname of the target
  server in the initial SSL handshake.
- SNI works when using Application Load Balancer (ALB), Network Load Balancer
  (NLB) and CloudFront. Does NOT work for Classic Load Balancer (CLB).

In summary:

- Classic Load Balancer (v1)
    - supports one cert (CLB)
    - need multiple load balancers and certs for multiple hostnames
- Application Load Balancer and Network Load Balancer (v2)
    - supports multiple listeners with multiple SSL certs
    - uses SNI to make it work

#### ELB Connection Draining

- For CLB is called _connection draining_
- For ALB and NLB is called _target group deregistration delay_
- Specifies the time to complete _in flight_ requests when a target EC2 instance
  is de-registering or is unhealthy.
- Will stop sending _new_ requests to de-registering / unhealthy instance.
- By default is 300 seconds, can set 1-3600 seconds, or disable by setting to 0.
- Set to a low value if requests are short, eg simple web app
- Set higher for longer requests, eg complex db lookups

### Auto-scaling Groups

- Allows increase (_scale out_) and decrease (_scale in_) of EC2 instances in
  response to changing loads.
- Ensures a minimum and maximum number of instances.
- Automatically attaches instances to a load balancer.
- Settings for an auto-scaling group include the usual EC2 parameters and well
  as scaling policies and load balancer info.
- Auto Scaling Alarms use CloudWatch alarms to monitors metrics and trigger
  scaling, eg target average CPU usage, number of requests, network load etc.
- Can also create CloudWatch custom metrics.
- ASGs use _Launch Configurations_ or _Launch Templates_ (newer)
- To update an ASG you provide a new launch config / template.
- IAM roles attached to an ASG will also be assigned to an EC2 instance.
- ASG is free, pay for underlying resources.
- Instances under an ASG that are accidentally terminated will automatically be
  re-created - extra safety!
- For example, an instance might be marked as unhealthy by a load balancer, the
  ASG would then terminate and replace that unhealthy instance.
