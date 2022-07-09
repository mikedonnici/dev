# Docker Swarm

Check if swarm enabled:

```bash
docker info
...
Swarm: inactive
...
```

Activate swarm:

```bash
docker swarm init
```

- Creates Root Signing Cert for swarm
- Issues cert for first manager node
- Creates join tokens
- Raft database is created to store root CA, configs and secrets
    - Encrypted on disk
    - No need for another key/value system to hold orchestration secrets
    - Replicates log amongst managers via mutual TLS in _control plane_

`docker swarm` is a narrow-scope command, for joining / leaving swarms,
promoting managers etc.

```bash
docker swarm --help
```

`docker node` command uses to manage nodes

Eg, show nodes:

```bash
docker node ls
ID                            HOSTNAME   STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
se34h87576uup25vtwu4fo3c3 *   gsnsw      Ready     Active         Leader           20.10.6
```

`docker service` is the main command for doing things with swarm

Create a service:

```bash
docker service create alpine ping 8.8.8.8
rn1yfs2802a7tiynil1or8eqq <--- this is the service id
```

List services:

```bash
docker service ls
ID             NAME            MODE         REPLICAS   IMAGE           PORTS
rn1yfs2802a7   elastic_nobel   replicated   1/1        alpine:latest 
```

List tasks (service processes):

```bash
docker service ps elastic_nobel # <-- name or service id
ID             NAME              IMAGE           NODE      DESIRED STATE   CURRENT STATE           ERROR     PORTS
bunbot099yq2   elastic_nobel.1   alpine:latest   gsnsw     Running         Running 3 minutes ago
```

Scale service up:

```bash
docker service update elastic_nobel --replicas 3
...
docker service ls
ID             NAME            MODE         REPLICAS   IMAGE           PORTS
rn1yfs2802a7   elastic_nobel   replicated   3/3        alpine:latest
...
docker service ps elastic_nobel
ID             NAME              IMAGE           NODE      DESIRED STATE   CURRENT STATE                ERROR     PORTS
bunbot099yq2   elastic_nobel.1   alpine:latest   gsnsw     Running         Running 7 minutes ago                  
lkojtx8er6t7   elastic_nobel.2   alpine:latest   gsnsw     Running         Running about a minute ago             
ruib2ac070fg   elastic_nobel.3   alpine:latest   gsnsw     Running         Running about a minute ago 
```

Remove the service:

```bash
docker service rm elastic_nobel
```

## Swarm Cluster

Init swarm and advertise an IP accessible by other nodes

```bash
docker swarm init --advertise-addr [IP]
```

Then cut and paste the generated join command on the other nodes. To regenerate
the join command can run
`docker swarm join-token worker` or `docker swarm join-token manager`.

Swarm Visualiser, run this on one node

```bash
docker service create --name=viz --publish=8080:8080/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock bretfisher/visualizer 
```

Then access on  `:8080`

### Overlay Multi-Host Networking

- When creating a network choose `--driver overlay` so that containers can
  access each other as though on a VLAN.
- This is for container-to-container traffic inside a single swarm.
- Can enable full encryption between containers using IPSec (off by default).
- Each service can be added to zero or more overlay networks.

Example - Creating a drupal network:

```bash
docker network create --driver overlay mydrupal
docker network ls
docker service create --name psql --network mydrupal -e POSTGRES_PASSWORD=mypass postgres
docker service create --name drupal --network mydrupal -p 80:80 drupal
```

Once this is complete the drupal app will respond on any IP, even if it is only
running on one node. This is because of the _routing mesh_.

### Routing Mesh

- Routes ingress packets for a _Service_ to the appropriate _Task_ (container)
- Spans all nodes in swarm
- Uses IPVS from Linux Kernel
- Load balances Swarm Services across their Tasks
- Two ways this works:
    - Container-to-container in an Overlay network - puts a Virtual IP in front
      of a Service and acts like a load balancer
    - External traffic incoming to published ports is routed to the correct
      container and port

![routing mesh 1](./routing-mesh1.png)

![routing mesh 2](./routing-mesh2.png)

- This is _stateless_ load balancing, ie cannot handle sessions
- Is layer 3 (IP/Port), not layer 4 (DNS)
- Both these limitations can be overcome using a reverse proxy, or Docker
  Enterprise Edition has a builtin layer 4 web proxy.

### Swarm Stacks in Production

- Stacks accept compose files as a declarative definition for services, networks
  volumes, and secrets.
- Use `docker stack deploy` instead of `docker service create`
- Stacks can manage all of the objects, including an overlay network, prefixed
  with the stack name.
- Compose file with use a `deploy:` key and cannot do `build:`
- Both `build:` and `deploy:` can exist in the compose file as Compose will just
  ignore `deploy:` and Swarm will ignore `build:`
- `docker-compose` cli does _not_ need to be present on the Swarm server as it
  is not considered to be a production tool. So Swarm can process the compose
  file without the need for `docker-compose`.
- A stack definition can be used to manage many services, but only a single
  swarm.
- Note: need to use version "3" or more for the compose file

![stack](./stack.png)

- Example run:

```bash
docker stack deploy -c app-stack.yml appname
#... to deploy and then can look at things with:
docker stack ls
docker stack services appname
docker stack ps appname
```

- To update a stack, change the compose file and then redeploy it with the same
  command.

#### Secrets storage

- Easy, secure solution for storing secrets in Swarm.
- Encrypted on disk and in transit.
- Use for passwords, keys, TLS certs etc.
- Can store string or binary up to 500kb.
- Since Docker 1.13 Swarm Raft DB is encrypted on disk.
- Only stored on disk on _manager_ nodes
- Secrets are first stored in Swarm using `docker secrets` and then assigned to
  services.
- Only assigned services can see the assigned secrets.
- Secrets appear as files in the container but are actually stored in an
  _in-memory_ file system, eg:
    - `/run/secrets/<secret_name>`
    - `/run/secrets/<secret_alias>`
- Local `docker-compose` can use file-based secrets as well, but not secure.

Create a secret from a file:

```bash
docker secret create psqluser psql_user.txt
```

Echo a secret and read from stdin:

```bash
echo "superSecretPassword" | docker secret create psql_pass -
```

Note: need to ensure that secrets are not visible in bash history etc.

Can list and inspect secrets but will not be able to read the actual value -
this is only available to running containers.

```bash
docker secret ls
ID                          NAME        DRIVER    CREATED         UPDATED
k1heq0feos0tzkyexc2llb9y4   psql_pass             3 minutes ago   3 minutes ago
m0gjg9yz42d4fakvgiklxr5h0   psql_user             2 minutes ago   2 minutes ago
```

```bash
docker secret inspect psql_user
[
    {
        "ID": "m0gjg9yz42d4fakvgiklxr5h0",
        "Version": {
            "Index": 216
        },
        "CreatedAt": "2021-06-27T22:07:45.827612968Z",
        "UpdatedAt": "2021-06-27T22:07:45.827612968Z",
        "Spec": {
            "Name": "psql_user",
            "Labels": {}
        }
    }
]
```

#### Starting a container with secrets

Note that the postgres image has the convention of being able to set required
env vars from a file using `POSTGRES_USER_FILE` and `POSTGRES_PASSWORD_FILE`.

So to use swarm secrets as below the image would have to support a similar
convention so that it could read to secrets from the container
at `/run/secrets/`.

```bash
docker service create --name psql \
  --secret psql_user \
  --secret psql_pass \
  -e POSTGRES_USER_FILE=/run/secrets/psql_pass \
  -e POSTGRES_PASSWORD_FILE=/run/secrets/psql_pass \
  postgres
```

Can add or remove secrets for an existing service however this will recreate the
container:

```bash
docker service update --secret-rm [list]
docker service update secret-add [secret]
```

#### Using secrets with compose

- Needs `version: "3.1"` minimum

```yaml
# docker-compose.yml

version: "3.1"

services:
  psql:
    image: postgres
    secrets:
      - psql_user
      - psql_password
    environment:
      POSTGRES_USER_FILE: /run/secrets/psql_user
      POSTGRESS_PASSWORD_FILE: /run/secrets/psql_password

# secrets are defined at root level key and then assigned to specific images
secrets:
  psql_user:
    # can use file: or external:
    file: ./psql_user.txt
  psql_password:
    file: ./psql_password.txt
```

To start:

```bash
docker stack deploy -c docker-compose.yml mydb
```

Just to re-iterate, the above will _work_ with `docker-compose` locally as well.
It does this by copying the secrets to `/run/secrets` on the running container
as a normal file mount. This means it only works with `file:` secrets and not
`external:`.

This allows for local development environment to match the production
configuration but is _not_ secure. It is not _intended_ to be secure because
`docker-compose` is not intended for production use.

### Service Updates

- Rolling replacement of tasks/containers in a service
- Limits downtime depending on the type of service
- Will generally _replace_ containers
- `update` command has many options
- Many options are `create` options with a combination of `-add` and `-rm`
- Also includes `scale`, `rollback` and `healthcheck` options
- A `stack deploy` toa an existing stack will issue service updates

#### Examples

- Update image to a newer version

```bash
docker service update --image fooapp:1.2.1 <servicenmae>
```

- Add env var and remove a port:

```bash
docker service update --env-add MODE_ENV=production --publish-rm 8080
```

- Change replicas for two services:

```bash
docker service scale web=8 api=6 
```

- In a stack file just edit the `YAML` file and:

```bash
docker stack deploy -c file.yaml <stackname>
```

Example run:

```bash
# Create a service
docker service create -p 8088:80 --name web nginx:1.13.7

# Scale up
docker service scale web=3

# Change image version
docker service update --image nginx:1.13.6 web

# Switch ports (remove and add)
docker service update --publish-rm 8088 --publish-add 9090:80 web

# TIP: re-balance service (update with no changes)
docker service update --force web 

# Clean up
docker service rm web
```

### Docker Healthchecks

- Healthcheck feature added in 1.12, and can be used in many places
- Recommended for production
- Docker engine runs the command in the container, eg `curl localhost`
- Expects `exit 0` (ok) or `exit 1` (error)
- Healthcheck has 3 states:
    - Starting
    - Healthy
    - Unhealthy
- Better than _in binary still running?_ but not a replacement for monitoring
  tools
- Where healthcheck can be seen:
    - `docker container ls`
    - `docker contaienr inspect` - last 5 healthchecks
- Docker run does nothing with healthchecks
- Services will replace tasks if they fail healthcheck
- Service updates wait for healthchecks before continuing

**`docker run` example using Elastic Search image:**

```bash
docker run  \
  --health-cmd="curl -f localhost:9200/_cluster/health || false" \  
  --health-interval=5s \
  --health-retries=3 \
  --health-timeout=2s \ 
  --health-start-period=15s \
  elasticsearch:2
```

Note that ` || false` in the `curl` command ensures that `exit 1` is issued for
any error as `curl` can exit with other codes, but docker requires only `0`
or `1`. Could also use `exit 1` instead of `false`

**`Dockerfile` example using nginx:**

- Note this fails because `curl` cannot be found

```dockerfile
FROM nginx:1.13

HEALTHCHECK --interval=30s --timeout=3s CMD curl -f http://localhost/ || exit 1
```

- If above is built and run can inspect health with:

```bash
docker inspect <container_id> | jq .[0].State
```

**Compose/Stack file example:**

- Version "2.1" is minimum for healthcheck
- Version "3.4" is minimum for `start_period:`

```yaml
version: "3.4"

services:
  web:
    image: nginx
    healthcheck:
      test: [ "CMD", "curl", "-f", "http://localhost" ]
      interval: 1m30s
      timeout: 10s
      retries: 3
      start_period: 1m
```

**`docker service` example using postgres:**

```bash
docker service create --name pg --health-cmd="pg_isready -U postgres || exit 1" -e POSTGRES_PASSWORD="abc123" postgres
```

### Container Placement (Task Assignment)

- By default swarm spreads tasks across nodes with least-used nodes next
- Can control which node a container runs on using:
    - Service constraints
    - Service modes
    - Placement preferences
    - Node availability
    - Resource requirements

#### Service constraints

- Hard requirement, placement fails if not matched
- Can filter task placement based on built-in or custom labels
- Can be added at create or at update
- Supports multiple constraints
- Supports either a <key> or <key>=<value> pair matching `==` or `!=`
- Labels can be node.labels or engine.labels
    - node.labels are stored in the raft database and can only be added from a
      manager node with a command - more convenient and easier.
    - engine.labels are added in `daemon.json` and then restarting docker so
      less convenient and used when can't access manager or to autoscale
      hardware or operating system setup.
- Note that constraints starting with `node.labels.xxx` or `engine.labels.xxx`
  are always custom labels, but there are existing constraints that can be used
  as well.

**Examples**

- Place only on a manager (`node.role` is built-in):

```bash
docker service create --constraint=node.role==manager nginx
# OR
docker service create --constraint=node.role!=worker nginx
```

- Add label to node2 and a constraint:

```bash
docker node update --label-add=dmz=true node2
docker service create --constraint=node.labels.dmz==true nginx
```

- Can't change an existing constraint, only add/remove

- Constraint in a stack file, example where specific hardware label has been
  added

```yaml
version: "3.1"
services:
  db:
    image: mysql:5.7
    deploy:
      placement:
        constraints
        - node.labels.disk == ssd
```

##### Built-in Labels

- `node.id` - listed in `docker node ls`
- `node.hostname` - listed in `docker node ls`
- `node.ip`
- `node.role` - manager/worker
- `node.platform.os` - linux/windows/...
- `node.labels` - empty by default

#### Service modes - replicated / global

- Can view with `docker service ls`
- Default is `replicated` which spreads tasks across nodes
- Can set `global` which will put the service once only on _every_ node
- Must set on `service create` and remove service to change
- Can combine with constraints

**Examples**

Put nginx on all nodes:

```bash
docker service create --name web1 --mode global nginx
```

Put nginx on all workers:

```bash
docker service create --name web2 --constraint node.role==worker --mode global nginx
```

Service mode in a stack file:

```yaml
version: "3.1"
services:
  web:
    image: nginx
    deploy:
      mode: global
```

#### Placement preferences

- _Soft_ requirement, so tries but will carry on if it does not succeed
- Works on `service create` and `service update`
- Can add multiple preferences for multi-layer placement control
- Only one strategy at the moment - `spread`:
    - spreads tasks amongst all values of a label
    - Useful for ensuring distribution across availability zones, subnets etc
- Won't move service tasks if labels are changed
- Use with constraints if the targeted label is not specified on all nodes. This
  is because a missing label is treated as though the label exists, but has a
  null value.

**Example: Spreading tasks across AWS AZs**

```bash
# Label ALL nodes
docker node update --label-add azone=1 node1
docker node update --label-add azone=2 node2
docker node update --label-add azone=3 node3

# Deploy service across all availability zones
docker service create placement-pref spread=node.labels.azone --replicas 3 nginx
```

Could also spread across subnets in a similar way

Ref: <https://docs.docker.com/engine/swarm/services/#placement-constraints>

**Example: Placement preference in a stack file**

```yaml
version: "3.1"
services:
  web:
    image: nginx
    deploy:
      placement:
        preferences:
          spread: node.labels.azone
```

#### Node availability

- Each node can have one of three states:
    - `active` - runs existing tasks, available for new tasks
    - `pause` - runs existing tasks, NOT available for new tasks, good for
      troubleshooting
    - `drain` - reschedules existing tasks, not available for new tasks, good
      for maintenance / updates
- Note that paused or drained nodes won't get service updates

**Examples:**

```bash
# pause node 2
docker node update --availability pause node2

# drain node 3
docker node update --availability drain node3
```

#### Resource requirements

- `docker run` has many options, `service create` has fewer
- set at `service create/ update` but are controlled per container
- set for cpu and memory, reserving and limiting
- Limiting: eg maximum given to a container
    - `--limit.cpu .5`
    - `--limit-memory 256M`
- need to be very careful about limiting memory
- Reservations: ensure minimum free resources in order to schedule a container
    - Swarm keeps track in its database
    - Does NOT involve and actual process monitoring so has no relationship to
      the actual amount being used by a task, rather is a promise to ensure a
      certain amount of resource will be available, eg:
    - `--reserve-cpu .5`
    - `--reserve-memory 256M`

**Examples:**

```bash
# Full CPU for MySQL
docker service create --reserve-memory 800M --reserve-cpu 1 mysql

# Limit resource for nginx
docker service create --limit-memory 150M --limit-cpu .25 nginx

# To remove limits / reservations (not actually set to 0)
docker service update --limit-memory 0 --limit-cpu 0 mysql 
```

**Stack file example:**

```yaml
version: "3.1"
service:
  database:
    image: mysql
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 1G
        reservations:
          cpus: '0.5'
          memory: 500M
```

---

## Swarm Operation in Production

### Service Logs

- Same as `docker container logs` but aggregates all service tasks
- Can return all tasks at once, or just one task's logs
- Great for real time cli troubleshooting
- Can follow and tail logs, and other options
- Not for log storage or searching or feeding into other systems
- Does _not_ work if you use `--log-driver` for sending logs off server

**Examples:**

```bash
# Return all logs for a service
docker service logs <servicename/id>

# Return all logs for a single task
docker service logs <taskid>

# Return unformatted logs with no trunking
docker service logs --raw --no-trunc <servicename/id>

# Return last 50 and all future logs
docker service logs --tail 50 --follow <servicename/id>
```

Reminder: `docker service ps <service name / id>`shows the _processes_ (tasks),
belonging to that service, and the task id for each - eg:

```bash
docker service ps web01
```

Although logging is not very sophisticated can `grep` logs to find particular
lines. Note that need to ensure both stderr and stdout are grep'd:

```bash
docker service logs <service name> 2>&1 | grep <search term>
```

### Docker Events

- Shows _actions taken_ in docker engine and swarm, eg `network create`,
  `service update`
- Has search filtering and formatting
- Limited to last 1000 events, not stored on disk so no long term logs.
- Good for watching real-time events.
- Two scopes "swarm" - see things on swarm managers, and "local" - confined to
  events on each worker node.

**Examples:**

```bash
# watch events from now, swarm scope if on manager, else local 
docker events

# Filter by time - lots of formats, see docs
docker events --since YYYY-MM-DD
docker events --since YYYY-MM-DDTHH:MM:SS
docker events --since 30m
docker events --since 2h30m 

# Last hour of events, with filters- note same filter key acts like OR
# and diff filter keys is like AND 
docker events --since 1h --filter event=start
docker events --since 1h --filter scope=swarm --filter type=network
```

## Swarm Configs

- Available since 17.06+
- Provides a way to map strings or files to _any_ file path in a task
- Similar to secrets but files can be stored anywhere in container
- Useful for things like nginx, mysql configuration files
- Benefit is being able to tweak a configuration without having to create a
  custom image or use a bind mount
- Immutable, so need to use a rotation process to rollout
- Can be removed once service is removed
- Strings are saved to the Raft log, so are highly available
- Should _NOT_ be used for private keys or other sensitive data (use secrets)
- Should _NOT_ be used in place of env vars

**Examples:**

```bash
# create a new config from an existing nginx config
docker config create nginx01 ./nginx.conf
# Use the above config
docker service create --config source=nginx01,target=/etc/nginx/conf.d/default.conf

# To update the above config
docker config create nginx02 ./nginx.conf
docker service update --config-rm nginx01 \ 
  --config-add source=nginx02,target=/etc/nginx/conf.d/default.conf
```

**Stack file:**

```yaml
version: "3.3" # minimum
services:
  web:
    image: nginx
    configs:
      - source: nginx-proxy
        target: /etc/nginx/conf.d/default.conf

configs:
  nginx-proxy:
    file: ./nginx-app.conf
```

## Limiting downtime with rolling updates, healthchecks and rollbacks

### Rolling service updates

- Most 'Day 2' ops involve `docker service update` so understanding how these
  work is crucial
- By default, `service update` replaces each replica one at a time, but many
  things can be customised.
- Update processes should be very thoroughly tested before going into production
  because different applications handle sessions and reconnections differently.

Testing tools:

- `httping` - Like ping but uses TCP instead of ICMP. Can install with `apt` or
  `docker run bretfisher/httping localhost`
- `browncoat` - simple web app with settings for acting badly
  <https://hub.docker.com/r/bretfisher/browncoat/>

#### A sample run

- Create a network

```bash
docker network create --driver overlay --attachable verse
```

- Create browncoat service at v1 - returns a 201

```bash
docker service create --name firefly -p80:80 --network verse --replicas 5 bretfisher/browncoat:v1
```

- Run httping with a few options

```bash
docker run --rm --network verse bretfisher/httping -i .1 -GsY firefly/healthz
```

- Rolling update to the v2 image - returns a 202

```bash
docker service update --image bretfisher/browncoat:v2 firefly
```

- Slow the container start up

```bash
docker service update --env-add DELAY_STARTUP=5000
```

#### Timeline of a service update

This process happens for each task:

1. Swarm will update N instances at a time, defaults to 1 but can be updated
   with `update-parallelism`
1. New tasks are created, and their desired state set to **Ready**:
    1. Ensures resource availability - **Pending**
    1. Pulls the image, if necessary - **Pending**
    1. Creates the container without starting it - **Ready**
    1. If task fails to achieve **Ready** state it is deleted and a new one
       started
1. When the new task is **Ready**, desired state for old one is set to
   **Shutdown** - this may take some time, depending on the task and its current
   state.
1. When the old task achieves **Shutdown** the new task is started and set to **
   Running**
1. Waits for `update-delay` (default is 0) and continues with next task

#### Update options

- `--stop-grace-period` - time to wait before force killing container
  (ns|us|ms|m|h)
- `--stop-signal <string>` - signal to stop the container
- `--update-delay` - Delay between updates
- `--update-failure-action` - Action on update failure, `pause`|`continue`
  |`rollback` - in a test env probably `pause` or `continue`, in production
  probably `rollback`
- `--update-max-failure-ratio` - Tolerable failure rate during an update
- `--update-monitor` - Duration after each update to monitor for failure
- `--update-order` - Update order, `start-first`|`stop-first`
- `--update-parallelism` - Max tasks updated simultaneously

**Examples:**

```bash
# Monitor for 5 mins before next, rollback on failure
docker service update --update-failure-action rollback --update-monitor 5m node

# Update 5 at a time, up to 25% can fail until failure action
# If lots of containers and distributed failures are ok
docker service update --update-parallelism 5 --update-max-failure-ratio .25

# Start a new container BEFORE killing the old one
# Good for single replica services, but NOT for databases with vol storage
docker service update --update-order start-first wordpress
```

**Sample run:**

```bash
# Create a service, constrain to current node
docker network create --driver overlay --attachable verse
docker service create --name firefly -p 80:80 --network verse --replicas 5 --constraint node.hostname==node1 bretfisher/browncoat:v1

# Monitor for 15 seconds before next task (no op)
docker service update --update-monitor 15s firefly
# View the change in monitor time
docker service inspect --pretty firefly

# Scale up and then update 3 at a time, force update without any changes
docker service scale firefly=15
docker service update --update-parallelism=3 --force firefly 

# Start new container before killing old one (watch with docker events)
docker service scale firefly=1
docker service update --update-order start-first --force firefly
```

### Healthchecks and Updates

#### Healthcheck options

- `--health-cmd` - command to check health
- `--health-interval` - time between health checks, `ms`|`s`|`m`|`h`,
  default `30s`
- `--health-retries` - consecutive failures before considered unhealthy,
  default `3`
- `--health-start-period` - grace time for container to start before considered
  unstable, default `0s`
- `--health-timeout` - Max time for a check to run, default `30s`

- `--stop-grace-period` - Time to wait before force killing container,
  `ns`|`us`|`ms`|`s`|`m`|`h`, default `10s`
- `--no-healthcheck` - disable container-specific healthcheck


- Generally good to start simple and use defaults, then tune over time.
- Examples:
    - Ensure web app returns 200
    - Ensure nginx returns 200 for `/ping`
    - Ensure db accepts connects, return a tmp db
- Remember that container healthchecks are not useful for _integration_ checks -
  these should be left for external monitoring (Prometheus etc), eg:
    - Web API returns valid data
    - DB record has proper count
    - Web front end can query API, etc

Checkout <https://github.com/docker-library/healthcheck>

### Service Rollbacks - handling failures gracefully

Two main ways rollbacks are used:

**1. Manual rollback**

```bash
# New way
docker service rollback <service>

# Old way
docker service update --rollback
```

- No options, goes back to last service definition
- Only one previous _spec_ stored to rollback to, so successive rollbacks will
  _toggle_ between the current and previous state

**2. Automated rollback during update**

```bash
docker service update --on-failure-action ...
```

- Last resort if update doesn't go as planned
- Should set if possible as default may not be useful

**Rollback options:**

- `--rollback-delay` - delay between task rollbacks, `ns`|`us`|`ms`|`s`|`m`|`h`
- `--rollback-failure-action` - action on rollback failure, `pause`|`continue`
- `--rollback-max-failure-ratio` - tolerable rollback failure rate
- `--rollback-monitor` - duration after each task rollback to monitor for
  failure
- `--rollback-order` - Rollback order, `start-first`|`stop-first`
- `--rollback-parallelism` - max simultaneous rollbacks

Timeline of rollback is similar to service update.

### Stack compose file example

```yaml
version: '3.9'
services:
  redis:
    image: redis
    networks:
      - frontend
    healthcheck:
      test: [ "CMD", "redis-cli", "-h", "localhost", "ping" ]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 30s
    deploy:
      replicas: 1
      update_config:
        failure_action: rollback
  db:
    image: postgres:13
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - backend
    environment:
      - POSTGRES_HOST_AUTH_METHOD=trust
    healthcheck:
      test: [ "CMD-shell", "pg_isready"]
      interval: 10s
      timeout: 5s
      retries: 5
    deploy:
      replicas: 1
      update_config:
        failure_action: rollback
  vote:
    image: bretfisher/examplevotingapp_vote
    networks:
      - frontend
    deploy:
      replicas: 3
      update_config:
        failure_action: rollback
    healthcheck:
      test: [ "CMD", "curl", "-f", "http://localhost" ]
      interval: 30s
      timeout: 5s
      retries: 3
      start_period: 20s
    ports:
      - '5000:80'  
  result:
    image: bretfisher/examplevotingapp_result
    networks:
      - backend
    deploy:
      replicas: 3
      update_config:
        failure_action: rollback
    healthcheck:
      test: [ "CMD", "curl", "-f", "http://localhost" ]
      interval: 30s
      timeout: 5s
      retries: 3
      start_period: 10s
    ports:
      - '5001:80'  
  worker:
    image: bretfisher/examplevotingapp_worker
    networks:
      - frontend
      - backend
    deploy:
      replicas: 2
      update_config:
        failure_action: rollback
networks:
  frontend:
  backend:
volumes:
  db-data:
```
