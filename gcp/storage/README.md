# Storage

3 types of storage service

1. Object - Cloud Storage
2. File - Cloud Filestore
3. Block - Persistent disk

## Cloud Storage

- Object storage for images, videos or any format
- Stored in containers called _buckets_
- Can choose geographical location
- Encryption with google or own keys


## Cloud Filestore

- Managed file storage service for applications that require a file system interface
- Fully managed NFS file servers for the management of application data of VMs
- Good for migrating applications to GCP without having to rewrite
- Automatically scale up and down
- 99.99% regional availability
- Pay only for what is used
- Encryption at rest

## Persistent Disk

- Durable network storage device that compute instances access like physical disks
- Shared data between instances
- SSD (faster, higher cost) or HDD (slower, lower cost)
- Three types:
  - Zonal - default
  - Regional - replicated between 2 zones
  - Local SSD - very high IOPS, low latency, highest performance

