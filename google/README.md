# Google Cloud Platform

Notes
from [Developing Applications with Google Cloud](https://www.coursera.org/specializations/developing-apps-gcp)

## Google Cloud Fundamentals: Core Infrastructure

### GCP Regions and Zones

- **Multi-Region** (eg _Europe_) -> **Region** (eg _europe-west2_) -> **Zone** (
  eg _europe_west2_a_)
- Zones within Regions have fast connectivity between them (< 5ms)
- Zone is not necessarily a single data centre
- Resources can be placed in different regions for fault-tolerance and to reduce
  latency for users by location
- Some resource support placement in _Multi-Region_, for example Google Cloud
  Storage
- Regions within a Multi_Region are seperated by at least 160km.

### GCP Resource Hierarchy

- Organisation -> Folders (optional) -> Project -> GCP Resources
- Resources are always placed in a single project
- Projects can optionally be organised with folders
- Folders can be placed beneath a Organisation Nodes
- Policies can be defined at any level in the hierarchy and are applied top-down
  to all items below it in the hierarchy - this means lower-level policies take
  precedence

- Project ID is assigned by user, used frequently and is immutable
- Project Name is for convenience and can be changed
- Project Number is assigned by Google and used less often

- Folders provide a convenient way of grouping Projects so that policies may be
  applied to multiple projects at once - less error-prone
- To user Folders must have an Organisation node at the top of the hierarchy
- If account is part of a Workspace account then there will be one automatically
  created
- If not, can use Google Cloud Identity to create one

### Identity and Access Management



