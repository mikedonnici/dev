# Certified Kubernetes Application Developer

- [Core Concepts](#core-concepts)
    - [Kubernetes Architecture](#kubernetes-architecture)
    - [Create and Configure PODs](#create-and-configure-pods)
- [Configuration](#configuration)
    - [Commands and Arguments](#commands-and-arguments)
    - [Environment Vars](#environment-vars)
    - [ConfigMaps](#configmaps)
    - [Secrets](#secrets)
    - [SecurityContexts](#securitycontexts)
    - [ServiceAccounts](#service-accounts)

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

### Environment Vars

- Env vars can be set using a plain key-value pair, ConfigMap or Secrets

_set an env var with a key-value pair_:

```yaml
spec:
  containers:
    - name: webapp
      image: webapp
      env:
        - name: APP_COLOUR
          value: blue
```

#### ConfigMaps

- ConfigMaps are used to centrally manage environment data and to pass key-value
  data to pods, and hence containers.
- ConfigMaps are created, then injected into a POD definition
- Stored in plain text, so not good for sensitive information

_Imperative creation of a configMap_:

```shell
# using literals
kubectl create configmap \
   <config-name> --from-literal=ENV_VAR_1=value \
             --from-literal=ENV_VAR_2=value
             
# from a file
kubectl create configmap \
   <config-name> --from-file=<path-to-file> 
```

_Declarative creation of a ConfigMap_:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  APP_COLOUR: blue
  APP_MODE: prod
```

- Can then inject a ConfigMap into a POD definition:

```yaml
# ... # 
spec:
  containers:
    - # ... #
      envFrom:
        - configMapRef:
            name: app-config
```

- Can use a single env var:

```yaml
# ... # 
spec:
  containers:
    - # ... #
      env:
        - name: APP_COLOUR
          valueFrom:
            configMapKeyRef:
              name: app-config
              value: APP_COLOUR    
```

- Inject as a volume:

```yaml
# ... # 
volumes:
  - name: app-config-volume
    configMapKeyRef:
      name: app-config    
```

#### Secrets

- Used to store sensitive information
- Similar to `ConfigMap` except stored in a hashed format

_Imperative creation of a secret from literals_:

```shell
kubectl create secret generic <secrete-name> \
  --from-literal=<key1>=<value1> \
  --from-literal=<key2>=<value2>
```

_Imperative creation of a secret from a file_:

```shell
kubectl create secret generic <secrete-name> --fromt-file=<path-to-file>
```

_Declarative creation of a secret_:

- Values should be base64 encoded, ie `echo -n 'value' | base64`

```yaml
# secret-data.yml
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
data:
  DB_HOST: ZGJob3N0
  DB_USER: ZGJ1c2Vy
  DB_PASSWORD: c2VjcmV0cGFzc3dvcmQ=
```

- Create with `kubectl create -f secret-data.yml`
- To view:
    - `kubectl get secrets`
    - `kubectl describe secret app-secrets`
    - `kubectl get secret app-secrets -o yaml` (shows hashed values)
- To decode a secret `echo -n 'c2VjcmV0cGFzc3dvcmQ=' | base64 --decode`


- Inject Secret into a POD and values will be available as env vars:

_Inject all of the Secret data_:

```yaml
# ... # 
spec:
  containers:
    - # ... #
      envFrom:
        - secretRef:
            name: app-secrets
```

_Inject a single value from the Secret_:

```yaml
# ... # 
spec:
  containers:
    - # ... #
      env:
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: app-secrets # Secrete object name
              key: DB_PASSWORD
```

_Inject as a volume_:

```yaml
# ... # 
volumes:
  - name: app-secrets-volume
    secret:
      secretName: app-secrets    
```

- When injected as a volume, each secret is created as a file, with the value as
  its content.

```shell
ls /opt/app-secrets-volume
DB_HOST  DB_USER  DB_PASSWORD
```

#### SecurityContexts

- Docker container processes are run directly on the host machine, but in their
  own namespace.
- The pid for the same running process will be different when viewed on the host
  vs the container - processes isolation.
- By default, Docker runs processes as `root`, however the root user in the
  container hs fewer privileges (capabilities).
  See `/usr/include/linux/capabilities.h`
- Can run a containers as a different user with `--user=1000`, for example.
- Can also specify the user in the `Dockerfile`:

```dockerfile
FROM ubuntu
USER 1000
```

- The `root` user in the container can have capabilities adjusted
  with `--cap-add`, `--cap-drop` or `--privileged` to add all.

**In Kubernetes:**

- The same Docker security privileges can be managed in k8s as well.
- Security can be set at a container or POD level
- If set at POD level will carry over to container
- If set at both POD and container, container settings will override POD

_Set run-as-user at the POD level_:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: some-pod
spec:
  securityContext:
    runAsUser: 1000
  containers:
    - name: ubuntu
      image: ubuntu
      command: [ "sleep", "3600" ]
```

_Set run-as-user at the container level_:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: some-pod
spec:
  containers:
    - name: ubuntu
      image: ubuntu
      command: [ "sleep", "3600" ]
      # just move into the container
      securityContext:
        runAsUser: 1000
        capabilities:
          add: [ "MAC_ADMIN" ]
```

- Note that capabilities are only supported at the container level.

#### Service Accounts

- Two types of accounts in K8s:
    - User Accounts - used by humans, eg admin, dev
    - Service Accounts - used by machines, eg promethius, jenkins
- A `ServiceAccount` is required for an application to access to K8s api.
- Create: `kubectl create serviceaccount account-name`
- View: `kubctl get serviceaccunt`
- When a `ServiceAccount` is created a token is automatically generated, and is
  stored as a secret object
- The token can be viewed by describing the secret:
  `kubectl describe secrete secret-name`
- This token can be used as an authentication bearer token when contacting the
  K8s api, eg:
  `curl https:k8s-hostname:6443/api -insecure --header "Athorization: Bearer <token>"`
- A `ServiceAccount` can be givens access-based roles and permissions

- When the third-party application is hosted on the same K8s cluster, the
  process is easier
- The service token secret can be automatically mounted as a volume on the pod
  running the third-party application so it can be accessed easily.

- Each namespace has its own `default` service account
- Whenever a pod is created, the default service account and its token, are
  automatically mounted to that pod as a volume mount
- This default service account has restricted permissions and can only run basic
  api queries
- A `serviceAccount:` field can be added to the pod `spec:` to use a different
  service account
- You must delete and re-create the pod to change the service account, or update
  the deployment which will trigger a new rollout
- Can choose _not_ to mount the default service account with `spec`:
  `automountServiceAccountToken: false`




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
