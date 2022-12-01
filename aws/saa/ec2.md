/[aws](../README.md)/[aws-soluctions-architect-associate](index.md)
# aws-saa-ec2
Created: 2022-11-29 04:56

## Hands On
 - AWS Console
 - EC2 > Launch Instance
 - Advanced Details > User Data
 ```sh 
#!/bin/bash
# Use this for your user data (script from top to bottom)
# install httpd (Linux 2 version)
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
echo "<h1>Hello World from $(hostname -f)</h1>" > /var/www/html/index.html
```
 - Lauch Instace
 - See All intances
 - Wait pass the state from $Pending$ to $Running$ 
 - Get Public IP Address - may change if you restart instance
## Theory
### Instance Types
 - m5.2xlarge
	 - m: instance type
	 - 5: generation (improves over time)
	 - 2xlarge: size within the instance class
 - General Purpose
	 - balance between
		 - compute
		 - memory
		 - networking
	 - t2.micro is general purpose 
 - Compute Optimized
	 - Bacth processing workloads
	 - Media transcoding
	 - High performance web servers
	 - High performance computing (HPC)
	 - Scientific modeling & machine learning
	 - Dedicated gaming servers
 - Memory Optimized
	 - High performance, relational/non-relational db
	 - Distributed web scale cache stores
	 - In-memory databases optimized for BI
	 - Applications performing real-time processing
 - Storage Optimized
	 - High frequency online transaction processing (OLTP) systems
	 - Relational & NoSQL databases
	 - Cache for in-memory databases (Redis)
	 - Data warehousing apps
	 - Distributed file systems
	 - 
### Security Groups
 - "firewalls" ec2
 - network security
 - control traffic access throw ec2 instances
 - contain allow rules
 - can reference by IP or security groups
![[Pasted image 20221130185640.png]]
 - regulates
	 - access to ports
	 - authorised ipv4/ipv6 ranges
	 - control network traffic:
		 - inbound (from other to instance)
		 - outbound (from the instance to other)
![[Pasted image 20221130190006.png]]
 - Can be attached to multiple instances
 - Locked down to a region/ VPC combination
 - Does live "outside" the EC2
	 - if traffic is blocked the EC2
	 - instance won't see it
 - It's good to maintain one separate security group for ssh access
 - if application is not accessible (time out)
	 - then it's security group issue
 - if your application gives a "connection refused"error
	 - then it's an application error 
	 - or it's not lauched
 - inbound is blocked by default
 - outbound is authorized by default
 - Referencing other security groups Diagram
![[Pasted image 20221130191310.png]]
 - ports
	 - 22: ssh
	 - 21: ftp
	 - 22 sftp
	 - 80 http
	 - 443 https
	 - 3389 rdp (windows)
### SSH
 - Mac, Linux and Windows +10 use ssh
![[Pasted image 20221201101749.png]]
 - ssh ec2-user
 - ensure that ssh is enable
	 - Instance > Security > Port Range = 22
 - connect with ipv4 public
```sh
chmod 0400 ec2-tutorial.pem
ssh -i ec2-tutorial.pem ec2-user@<IPV4_EC2_INSTANCE>
```
 - aws following naming convention
 - manage IAM of instances
	 - Instances > Actions > Security > Modify IAM Role
### Questions
 - Which EC2 Purchasing Option can provide you the biggest discount, but it is not suitable for critical jobs or databases?
	 - Spot Instance
 - Spot instace can have 90% discount
 - *How long can you reserve an EC2 Reserved Instance?
	 - You must reserve the instance only for 1 or 3 years, not any time in between.
 - You would like to deploy a High-Performance Computing (HPC) application on EC2 instances. Which EC2 instance type should you choose?
	 - Compute Optimized EC2 instances are great for compute-intensive workloads requiring high-performance processors (e.g., batch processing, media transcoding, high-performance computing, scientific modeling & machine learning, and dedicated gaming servers).
 - Which EC2 Purchasing Option should you use for an application you plan to run on a server continuously for 1 year?
	 - Reserverd Instance
 - You are preparing to launch an application that will be hosted on a set of EC2 instances. This application needs some software installation and some OS packages need to be updated during the first launch. What is the best way to achieve this when you launch the EC2 instances?
	 - Write a bash and put on user data of EC2 instance
	 - EC2 User Data is used to bootstrap your EC2 instances using a bash script. This script can contain commands such as installing software/packages, download files from the Internet, or anything you want.
 - Which EC2 Instance Type should you choose for a critical application that uses an in-memory database?
	 - Memory Optimized EC2 instances are great for workloads requiring large data sets in memory.
 - You have an e-commerce application with an OLTP database hosted on-premises. This application has popularity which results in its database has thousands of requests per second. You want to migrate the database to an EC2 instance. Which EC2 Instance Type should you choose to handle this high-frequency OLTP database?
	 - Storage Optimized EC2 instances are great for workloads requiring high, sequential read/write access to large data sets on local storage.
 - Security Groups can be attached to only one EC2 instance.
	 - False
	 - Security Groups can be attached to multiple EC2 instances within the same AWS Region/VPC.
 - You're planning to migrate on-premises applications to AWS. Your company has strict compliance requirements that require your applications to run on dedicated servers. You also need to use your own server-bound software license to reduce costs. Which EC2 Purchasing Option is suitable for you?
	 - Dedicated Hosts are good for companies with strong compliance needs or for software that have complicated licensing models. This is the most expensive EC2 Purchasing Option available.
 - *You would like to deploy a database technology on an EC2 instance and the vendor license bills you based on the physical cores and underlying network socket visibility. Which EC2 Purchasing Option allows you to get visibility into them?
	 - Dedicated Hosts
## Advanced Topics
### Private vs Public vs Elastic IP
 - networking has two sorts of IPS (ipv4/ipv6)
 - public ip
	 - public ip means the machine can be identified on the internet
	 - must be unique accross the web
	 - can be geo-located easily
 - private ip
	 - means machine can only be indentified on private network only
	 - unique accross private net
	 - two different private net can have same ip
	 - machine connect www use internet gateway/proxy
	 - only specified ip range can be used as private ip
 - elastic ip
	 - problem: when stop and then start the ec2 instance, it can change public ip
	 - solve: if you need to have a fixed public ip, use elastic ip
	 - public ipv4 owned as long it was not deleted
	 - can attach it to one instance at a time
	 - mask failure of an intance rapidly remapping the address to another instance
	 - can only have 5 elastic ip in a same account (more need ask AWS)
	 - overall try to avoid using elastic ip
		 - reflect poor architectural decisions
		 - use random public ip and register a dns name to it
		 - use route 53/load balancers
### Placement Groups
 - console: EC2 > Placement Groups > Create placement group
 - EC2 > Instance > Launch Instance > Advanced > Placement group name
 - sometimes you want control over the EC2 Instance placement strategy
 - That strategy can be defined using placement groups
 - When you create a placement group, you specify one of the following strategies:
	 - Cluster
		 - clusters instances
		 - into low-latency group
		 - in a single Availability Zone
	 - Spread
		 - spreads instances
		 - across underlying hardware
		 - max 7 instances per AZ
		 - critical applications
	 - Partition
		 - spreads instances
		 - across many different partions
		 - which rely on different sets of racks
		 - within an AZ
		 - Scales 100s of EC2 instances per group
		 - allow runs Hadoop, Cassandra, Kafka
![[Pasted image 20221201173327.png]]
 - Cluster
	 - Pros
		 - Great network 
		 - 10gbps bandwidth between instances with enhanced networking enabled - recommended
	 - Cons 
		 - if the rack fails
		 - then all instances fails at the same time
	 - Use case
		 - Big Data job that needs to complete fast
		 - Application that needs 
			 - extremely low latency
			 - and high network throughput
![[Pasted image 20221201173631.png]]
 - Spread 
	 - Pros:
		 - can span across Availability Zones (AZ)
		 - Reduced risk is simultaneous failure
		 - EC2 Instances are on different physical hardware
	 - Cons:
		 - Limited to 7 instances per AZ per placement group
	 - Use case:
		 - Application that needs to maximize high availability
		 - Critical Applications
			 - each instance must be reliable
			 - if one fail, not affect others
![[Pasted image 20221201174202.png]]
 - Partition
	 - Pros
		 - Up to 7 partitions per AZ
		 - Can span across multiple AZs in the same region
		 - Up to 100s of EC2 instances
		 - The instances in a partition do not share racks with the instances in the other partitions
		 - A partition failure can affect many EC2, but won't affect other partitions
		 - EC2 instances get access to the partition information as metadata
	 - Cons:
		 - Cost
	 - Use Cases
		 - HDFS
		 - HBASE
		 - Cassandra
		 - Kafka
### Elastic Network Interfaces (ENI)
 - Logical component in a VPC that represents a virtual network card
 - The ENI can have the followeing attributes:
	 - Primary private IPv4, one or more secondary IPv4
	 - One Elastic IP (IPv4) per private IPv4
	 - One Public IPv4
	 - One or more security groups
	 - A MAC address
 - You can create ENI independently and attach them on the fly (move them) on EC2 instances for failover
 - Found to a specific AZ
![[Pasted image 20221201174745.png]]

### EC2 Hibernate
 - We know we can stop, terminate instances
	 - Stop - the data on disk (EBS) is kept intact in the next start
	 - Terminate - any EBS volumes (root) also set-up to be destroyed is lost
 - On start, the following happens
	 - First start: the OS boot & the EC2 User Data script is run
	 - Following starts: the OS boots up
	 - Then your application starts, cache  get warmed up, and that  can take time!
 - Hibernate
	 - The in-memory (RAM state is preserved)
	 - The instance boot is much faster ( the OS is not  stopped/restarted)
	 - Under the hood: the RAM state is written to a file in the root EBS volume
	 - The root EBS volume must be encrypted
 - Use cases
	 - Long-running processing
	 - Saving the RAM state
	 - Services that take time to initialize

## References
1. [EC2 Chapter Udemy](https://www.udemy.com/course/aws-certified-solutions-architect-associate-saa-c03/learn/lecture/26098122?start=60#overview)