# Compute Engine

- Types of CE instances:
   - Pre-defined machine type to quickly launch a VM
   - Custom machine types can specify number of vCPUs and memory
   - Spot machines when state is not important, up to 91% cheaper

- Machine types:
  - General purpose - balance of compute, memory and network, E2, N2, N2D, N1 machine types
  - Memory optimized - high memory, M2 and M1 machine types
  - Compute optimized - high CPU, C2 machine types
  - Accelerator optimized - parallel processing, GPU, A2 machine types

- Managed instance groups:
  - group of VMs that are managed as a single entity
  - used for autoscaling
  - multi-zone deployed instances
  - auto-healing and updating
  - requires the use of an instance template

## Managing Compute Resources

- Three main areas of concern:
   - Security - controlling access to instances, ssh keys
   - Snapshots - point-in-time images for disaster recovery across zones
   - Monitoring and logging - agents that give key insights to health and performance of VMs


