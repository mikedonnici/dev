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

**IAM Federation** is available for large enterprises to integrate with their own repository of users using SAML (Active Directory).

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
- Restarting EC2 instance can change its public IP (private IP remains unchanged)
- Elastic IPs are reserved public IPv4 addresses which can be attached to an EC2 instance
- Can only have 5 elastic IPs per account - can request more
- TIP: Avoid elastic IP - bad architecture. Prefer random IP with DNS assignment
- Best pattern is to use a load balancer

### EC2 User Data

- Bootstrap EC2 instance using _EC2 User Data_ script
- Script runs once only on the instance's first start up
- Used to automate tasks such as installing updates, software and other tasks
- Run with root privileges
- Is on step 3 when configuring EC2 instance:

![ec2 user data](ec2_user_data.png)

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
- Recommended for: short-term, uninterrupted workloads, unpredictable app behaviour

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
- Instances can be lost at any time if your **max price** is less than current spot price
- _Only useful for workloads that are resilient to failure_, eg:
    - Batch jobs
    - Data analysis
    - Image processing
- Single spot instance can be requested when launching an instance, on the _configure instance_ page
- Can also request a _fleet_ of spot instance via _Spot Request_ on left menu

A way to combine instances would be to have a reserved instance for baseline capacity, plus on demand 
instance(s) and / or spot instances for peaks, depending on if workload resilience is required.


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
- Used for software that has complicated / restricted _Bring Your Own Licence_ (BYOL) model or regulatory 
requirements

### Elastic Network Interfaces (ENI)

- Logical component in a VPC that represents a virtual network card
- Bound to a specific AZ
- An ENI can have the following attributes:
    - Primary private IPv4, one or more secondary IPv4
    - One Elastic IP (IPv4) per private IPv4
    - One public IPv4
    - One or more security groups
    - A MAC address
- ENI can be created independently and can be attached / moved to EC2 instances for failover  

![eni](eni.png)

- Access ENIs via left menu on EC2 page:
  
![eni](eni_link1.png)

- Or via `eth0` link on an EC2 instance:

![eni](eni_link2.png)






 
    


   
   
   


