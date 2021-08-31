# Certified Kubernetes Application Developer

- [Core Concepts](#core-concepts)
    - [Kubernetes Architecture](#kubernetes-architecture)
    - [Create and Configure PODs](#create-and-configure-pods)
- [Configuration](#configuration)
    - [Commands and Arguments](#commands-and-arguments)
    - ConfigMaps
    -
- Multi-Container PODs
- Observability
- POD Design
- Services & Networking
- State Persistence

- Appendix 1: [`kubectl` commands](#kubectl-commands)

## Core Concepts

### Kubernetes Architecture

- **Node**: Physical / virtual machine which is a k8s worker
- **Cluster**: set of nodes grouped together
- **Master**: K8s node that manages containers on workers

#### K8s Components

- **API Server**: front end for k8s (master)
- **etcd key store**: distributed key-value store (master)
- **Scheduler**: distributes work amongst nodes *master)
- **Controller**: make decisions to bring up new containers as required (master)
- **Container runtime**: Eg Docker, rkt, CRI-O (worker)
- **kubelet**: Agent on each node that communicates with master (worker)

### Create and Configure PODs

#### PODs

- Smallest object that can be created in K8s
- Containers are encapsulated into a POD
- POD is a single instance of an application
- Generally 1-1 POD to container
- Can have more that one container in a POD, eg a helper container, but rare use
  case

_Pod definition file_:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
    - name: nginx
      image: nginx
```

#### Replication Controller

- One of the controllers that manages the number of PODs in response to certain
  circumstances
- `ReplicationController` is the older technology, now use `ReplicaSet`
- `ReplicaSet` _requires_ a `selector:` because `ReplicaSet` _can_ manage PODs
  that were _not_ created by the `ReplicaSet`
- `ReplicationController` can have `selector:` but can be omitted, in which case
  is assumed to be same as the template.
- `ReplicaSet` is a process that monitors the PODs, which is why labels come in
  handy

_ReplicationController definition file_:

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: myapp-rc
  labels:
    app: myapp
    type: front-end
spec:
  template:
    # Pod definition here
    metadata:
      name: nginx
      labels:
        app: myapp
        type: front-end
    spec:
      containers:
        - name: nginx
          image: nginx
  replicas: 3
```

_ReplicaSet definition file_:

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: myapp-rs
  labels:
    app: myapp
    type: front-end
spec:
  template:
    metadata:
      name: nginx
      labels:
        app: myapp
        type: front-end
    spec:
      containers:
        - name: nginx
          image: nginx
  replicas: 3
  selector:
    matchLabels: # matches Pod label
      type: front-end
```

#### Deployments

- `Deployment` objects are used to manage a `ReplicaSets`.
- Creation of a `Deployment` with automatically create a `ReplicaSet`
- A new `Deployment` triggers a _rollout_, changes to underlying containers
  trigger a new _rollout_ and _revisions_ are tracked to allow for _rollback_.

_Deployment definition file_:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-rs
  labels:
    app: myapp
    type: front-end
spec:
  template:
    metadata:
      name: nginx
      labels:
        app: myapp
        type: front-end
    spec:
      containers:
        - name: nginx
          image: nginx
  replicas: 3
  selector:
    matchLabels: # matches Pod label
      type: front-end
```

#### Namespace

- `Default` namespace is created automatically when K8s is set up
- `kube-system` and `kube-public` are also created
- Namespaces isolate resources from other namespaces
- Small environments generally don't require additional namespaces
- Larger environments may require additional namespace, eg `DEV`, `PROD`
- A `Namespace` object can be created with cli:

```shell
kubectl create namespace dev
```

- ...or with a definition file:

_Namespace definition file_:

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: dev
```

- Each namespace can have quotas and policies to ensure resources
- Inside a namespace, services can be referenced by name, eg `db-service`
- Can communicate between namespaces by using internal DNS entries qualified
  name, eg `db-service.dev.svc.cluster.local`, where:
    - `db-service` - service name
    - `dev` - namespace
    - `svc` - service subdomain
    - `cluster.local` - default local domain
- By default, `kubectl` commands refer to `default` namespace
- use `--namespace` to query other namespaces, eg:
    - `kubectl get pods --namespace=kube-system`
- Same applies when creating resources, eg
    - `kubectl create -f def.yaml --namespace=dev`
- `namespace` can also be applied in a definition file:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  namespace: dev
spec:
  containers:
    - name: nginx
      image: nginx
```

- `ResourceQuota` object is created to limit resources for a namespace

_ResourceQuote definition file_:

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: compute-quota
  namespace: dev
spec:
  hard:
    pods: "10"
    requests.cpu: "4"
    requests.memory: 5Gi
    limits.cpu: "10"
    limits.memory: 10Gi
```

---

## Configuration

### Commands and Arguments

#### Commands, Arguments and Entrypoint in Docker

- Containers will exit when their task is complete, or the task crashes
- `CMD` in Dockerfile specifies the default task to run
- By default, docker does not attach a terminal to a container when it is run
  which is why bash exits when used as a `CMD`
- Default task can be overridden when the container is started, eg:
    - `docker run ubuntu sleep 10`
- Can also pass args to the default command by using `ENTRYPOINT`, so if the
  container specified `ENTRYPOINT ["sleep"]`, then starting
  with `docker run <image-name> 10` would start the container with the
  command `sleep 10`
- So, passing args with `CMD` replaces the command entirely, passing args
  with `ENTRYPOINT` appends the args to the command.
- To specify a _default_ arg to an `ENTRYPOINT` use both `ENTRYPOINT` and `CMD`:

```dockerfile
FROM ubuntu
ENTRYPOINT ["sleep"]
CMD ["5"]
```

- Then:
  - `docker run custom-ubuntu sleep` will run `sleep 5`
  - `docker run custom-ubuntu sleep 10` will run `sleep 10` 

- `--entrypoint` can also be used to override the default `ENTRYPOINT` command
  - `docker run --entrypoint someothercommand custom-ubuntu`










---
UP TO HERE
---

## Service - NodePort

```yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp-service
spec:
  type: NodePort
  ports:
    - targetPort: 80
      port: 80 # if omitted defaults to targetPort
      nodePort: 30008 #  if ommitted, randomly selected in range
  selector: # all labels from pod metadata, or deployment spec
    app: myapp
    type: front-end
```

## Service - ClusterIP

```yaml
apiVersion: v1
kind: Service
metadata:
  name: back-end
spec:
  type: ClusterIP # this is the default type
  ports:
    - targetPort: 80 # port where backend is exposed
      port: 80       # port where service is exposed
  selector: # all labels from pod metadata, or deployment spec
    app: myapp
    type: back-end
```

---

### `kubectl` commands

```shell
kubectl run hello-minicube
kubectl cluster-info
kubectl get nodes

# extract POD definition to yaml file
kubectl get pod <pod-name> -o yaml > pod-definition.yml

kubectl create -f def.yml
kubectl replace -f def.yml


kubectl get <obj>
kubectl delete <obj> <name>

# scaling
kubectl scale --replicas=<n> -f def.yml


# deployment
kubectl rollout status deployment/name
kubectl rollout history deployment/name


# output formats
kubectl get all -o json  # Output a JSON formatted API object.
kubectl get all -o name  # Print only the resource name and nothing else.
kubectl get all -o wide  # Output in the plain-text format with any additional information.
kubectl get all -o yaml  # Output a YAML formatted API object

# Use non-default namespace
kubectl get pods --namespace=kube-system
kubectl create -f def.yaml --namespace=dev

# Switch namespace context
kubectl config set-context $(kubectl config current-context) --namespace=dev

# all namespaces
kubectl get pods --all-namespaces

# dry runs
kubectl create -f def.yml --dry-run=client
# create definition files
kubectl run custom-nginx --image=nginx --port=8080 --dry-run=client -o yaml
kubectl create deployment --image=nginx nginx --replicas=4 --dry-run -o yaml
kubectl expose pod redis --port=6379 --name redis-service --dry-run=client -o yaml
kubectl create service clusterip redis --tcp=6379:6379 --dry-run=client -o yaml
kubectl create service nodeport nginx --tcp=80:80 --node-port=30080 --dry-run=client -o yaml
```
