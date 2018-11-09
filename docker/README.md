# Docker

## Editions

There are various editions of Docker for different OS and for cloud.

On mac it is not supported natively so runs in a VM.

## Linux Installation

*Don't use package manager*

Installation script here: https://get.docker.com/

Install instruction at the top of the comments:above script:

```bash
$ curl -fsSL get.docker.com -o get-docker.sh
$ sh get-docker.sh
```

Or, shorthand: `curl -sSL https://get.docker.com/ | sh`

Docker required root access so will require `sudo` to run things. This is painful in local dev
so to run docker as non-root user need to add user account to 'docker' group, eg:

```bash
$ sudo usermod -aG docker mike
```

*This won't work on Red Hat flavours of Linux.*

[`docker-machine`](https://docs.docker.com/machine) and [`docker-compose`](https://docs.docker.com/compose/) are not installed automatically on Linux so need to 
install those manually. 

Check GitHub releases for latest version and there is a curl script to install.

* [`docker-machine` releases](https://github.com/docker/machine/releases) 
* [`docker-compose` releases](https://github.com/docker/compose/releases)

**These will need to be manually updated.**

## docker cli

As the number of docker cli commands grew a *management command* format was introduced, ie `docker [command] [sub-command]`.

For example:

```bash
$ docker container run
```

The old way still works for many commands, ie:

```bash
$ docker run
```

Newer commands will generally follow the *management command* format.

## Image vs Container

* An *image* is the blueprint for the application that will be run
* A *container* is an instance of an *image*, running as a process

Images are store in a *registry* such as [hub.docker.com](https://hub.docker.com)

**A container is not a VM - it is a process running on the OS with limted access to resources**

For example, if an nginx container is started in the background, thus:

```bash
$ docker container run --publish 80:80 --detach --name webserver nginx
```

The process running in the container can be listed with:

```bash
$ docker top [id...]
```

If `ps aux` is run on the host machine the same processes are visible - that is, they are not hidden inside a virtual container.

## Running a container

Use `docker container run ...` with various options

For example:

```bash
$ docker container run --publish 80:80 --name webserver --detach nginx
```

This will start an nginx server in the background, mapping port 80 on the host to port 80 in the container. The `--name` flag gives the container a more convenient name that the generated, random id.

There are many options - see `docker container run --help`.

## Monitoring container processes

Some CLI commands for looking at container processes:

* `docker container top` - process list in one container
* `docker container inspect` - config details for one container
* `docker container stats` - performance stats for one or all containers

## Running a shell in a container

Can use the `-it` flag (**i**nteractive, **t**erminal) when a container is run.

For example, run a container using the nginx image and start `bash` instead of the default command.

```bash
$ docker container run -it --name webserver nginx bash
```

If we `exit` bash the container stops. To restart the same container (bash is the start command), and connect interactively again, use the `-ai` flag (**a**ttach, **i**nteractive):

```bash
$ docker container start -ai webserver
```

To connect to an existing, running container use the `exec` command:

For example, run bash on a container named *db1*:

```bash
$ docker container exec -it db1 bash
```

## Docker networking

* Each container connects to a private virtual netwoek *bridge*
* Each *virtual network* routes through NAT firewall on host IP
* Containers on a virtual network can communicate without `-p`
* Best practice is to create a *virtual network* for each app

`--publish` or `-p host:container` exposes a port on the host to a port on the container. To see which ports are open for a particular container:

```bash
$ docker container port [name/id]
```

`docker container inspect` can be used to look at other networking details. It also has a `--format` flag which uses Go templates to format the output. For example:

```bash
docker container inspect --format '{{ .NetworkSettings.IPAddress }}'
```

### Network CLI commands

* `docker network ls` - show networks
* `docker network inspect` - inspect a network
* `docker network create --driver` - create a network
* `docker network connect` - attach a network to a container
* `docker network disconnect` - detach a network from a container

Docker handles network set up with sensible defaults, especially for local dev and testing. However, the networking can be customised.

Running `docker network ls` will generally show at least the following default networks:

```bash
NETWORK ID          NAME                DRIVER              SCOPE
fa13e96005b6        bridge              bridge              local
7ce977bddfb9        host                host                local
2eecffb3120b        none                null                local
```

The network *DRIVER* refers to the type of network, the default is *bridge*. 

The *host* network is used to bypass the virtual network and connect the container directly to the host's network. This is faster but bypasses the security features of containers.

The *none* network is exactly that - no network.

To inspect the *bridge* network:

```bash
$ docker network inspect bridge
```


When new containers are created they are attached to the default *bridge* network.

New networks can also be created, eg:

```bash
docker network create my_app_net
```

This creates a new virtual network, assigned the next available IP range in the `172.x.x.x` series. This can be customises, of course. See `--help` for options.

`docker network ls` now shows:

```bash
NETWORK ID          NAME                DRIVER              SCOPE
fa13e96005b6        bridge              bridge              local
7ce977bddfb9        host                host                local
0f7f556d5f95        my_app_net          bridge              local
2eecffb3120b        none                null                local
```

To create a container and attach it to this new network:

```
$ docker container run -d --name web1 -p 80:80 --network my_app_net nginx:alpine
```

Inspecting *my_app_net* will show that the new container is attached:

```bash
$ docker network inspect my_app_net
```
...

To connect an existing container to a network:

```bash
$ docker network connect [network id] [container id]
```

`docker container inspect [name/id]` will show that the container has been connected to the new network.

Similarly, to detach a container:

```bash
$ docker network disconnect [network id] [container id]
```

## Docker DNS

Containers can't rely on IPs for communication as they are dynamic. So docker uses container names as host names for DNS resolution.

The default *bridge* network does **not** link containers together using DNS. This *can* be done using the `--link` flag.

However, if a *new* network is created containers that are added will be automatically linked via DNS.

## DNS round robin

Multiple containers can respond to the same DNS name by using the `--net-alias` flag.

This can be used to create a multi-container application that can respond to one end point.

This exercise sets up three elasticsearch containers and then, from a centos container on the same network, checxks the roundrobin responses.

Create the *roundrobin* network:

```bash
$ docker network create roundrobin
```

Start three containers on the *roundrobin* network:

```bash
$ docker run -d --net roundrobin --net-alias search elasticsearch:2
$ docker run -d --net roundrobin --net-alias search elasticsearch:2
$ docker run -d --net roundrobin --net-alias search elasticsearch:2
```

Run `nslookup search` on alpine to check roundrobin set up:

```bash
$ docker run --rm --name t1 --net roundrobin alpine nslookup search localhost
Server:    127.0.0.1
Address 1: 127.0.0.1 localhost

Name:      search
Address 1: 172.21.0.4 e3.roundrobin
Address 2: 172.21.0.2 e1.roundrobin
Address 3: 172.21.0.3 e2.roundrobin
```

Run `curl -s search:9200` multiple times and see different names in response:

```bash
$ docker run --rm --name t1 --net roundrobin centos:7 curl -s search:9200
{
  "name" : "Hairbag" <- should get 3 variations
}
```

## Docker images

Docker registry as https://hub.docker.com

To fetch an image:

```bash
$ docker pull image:tag
```

If no tag is specified will default to `latest`.


The details of an image can be viewed with:

```bash
$ docker image inspect [image:tag]
```

This shows available ports, default command and lots of other stuff.


### Image layers and history

```bash
$ docker image history [image:tag]
```

Show the image layer history. 

Images start as a blank layer known as *scratch* and every set of changes to the 
file system on the image, is another layer.

Each layer has a unique *sha256* which identifies it. So image layers can be cached and referenced by this identifier. Thus, if two images are fetched that have, for example, the same base layer operating system, it is only downloaded once. This saves time and storage space.

This applies to container instances as well. For example, two containers running the same base image may have differences in their file systems (eg logs), however the original files from the image are referenced from the *same* image *layer*. Only the file system *differences* are stored in the container as a new layer. This uses a thing called a **file union system**.

**A container only stores the running process, and any file differences from the image.**

When pushing image to a repository layers are also not uploaded more than once. So both sides of the image process save space and time through the use of layers.


### Image tagging

Images are referenced with a specific format: `<user>/<repo>:<tag>`

***Official*** images don't have the `<user>/` part in the repository name.

If the tag is not specified then it will default to `latest`. That does not mean it is the *latest*, but generally it will be.

Often there are multiple tags that point to the same image. If an image is pulled with different tags that reference the *same* image there will be multiple entries shown with `docker image ls`, however the IMAGE ID will be the same and it is only stored once.

Any image can be re-tagged locally:

```bash
$ docker image tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
```
Again, if unspecified, the default tag will be `latest`.

## Building images: Dockerfile basics

Dockerfiles are comprised of different *stanzas* that perform specific tasks. The order matters as they run top-down.

`FROM`  -base image, required in all Dockerfiles

Base image depends on needs - if need to install other commands in the image then generally will use a Linux image like ubuntu or centos. A very basic, and small starting image is *alpine*.

`ENV` - set environment variables, important for setting key values in containers.

`RUN` - for executing shell commands such as installation of packages etc.

To ensure a set of commands are run within a single layer use `... \ ... &&` to join them together.

`EXPOSE` - declares the ports that are exposed on the docker virtual network

`CMD` - require parameter that declares the final command that is run when the container starts.

To build an image from `Dockerfile ` in the current dir:

```bash
$ docker build <imagename>:<tag> .
```

Each step in the build process, as defined in the `Dockerfile` is assigned a cache id. Subsequent builds will use previously cached layers **unless the line in the Dockerfile has changed**. 

If a line has changed the step is re-run, and all subsequent steps are also re-run. Hence, order is important - keep things at the top that change *least*, and things at the bottom that change more often.

`WORKDIR` - defines the working directory, better that `RUN cd /some/dir`

`COPY` - copy files from host to container

Dockerfile example:

```dockerfile
FROM nginx:latest
WORKDIR /usr/share/nginx/html
COPY index.html index.html
```

Here an official image is used. As the `nginx` images includes `EXPOSE` and `CMD` there is no need to include them in this Dockerfile.

## Persistent data

Docker has two methods of persisting data:

1. Volumes
1. Bind mounting

### Volumes 

This methods creates a new volume on the host and maps it to a location in the container.

Volumes can be created in a Dockerfile.

To view volumes:

```bash
$ docker volume ls
```

To inspect a volume: 

```bash
$ docker volume inspect [name/id]
```

**Volumes are not deleted when a container is deleted** - hence, they are persistent! They must be manually deleted.

By default, volumes are named with a long hash. However, they can be given custom names that are easier to work with. Named volumes are specified with `-v`.

```bash
$ docker container run ... -v nicename:/path/in/container image
```

The `nicename` can be used to refence the volume, and also is used in the host storage path in place of the longer hash.


`-v` can also be used to create new volumes. If the name is ommitted a new volume is created and named with a hash.

In some instances volumes need to be created ahread time time - that is, before the `run`. For example, as if a non-default volume driver is required.

### Bind mounting

A bind mount is a direct mapping from a host file/directory to a container file/directory. As such, they *cannot* be specified in a Dockerfile - only at `container run`.

A bind mount is also specifed with `-v`, however docker distinguishes a *bind mount* from a *named volume* when the argument start with a `/`.

For example:

```bash
$ docker container run ... -v /path/on/host:/path/in/container image
```

`$(pwd)` can also be used to specify the current host directory:

```bash
$ docker container run ... -v $(pwd):/path/in/container image
```




Bind mounts are very useful for local development when source files are changing frequently.









































## Port Mapping

## Volumes

Used at run to map files between the local file system and the container filesystem.

This can be used to avoid completely rebuilding the container when application source code is changed.

Assuming this `Dockerfile.dev`:

```docker
FROM node:alpine
WORKDIR '/app'
COPY package.json .
RUN npm install
COPY . .
CMD ["npm", "run", "start"]
```  

In the example below the *present working directory* `$(pwd)` is mapped to `/app` in the container. 
        
```bash
$ docker build -t mikedonnici/myapp -f Dockerfile.dev .
$ docker run -p 3000:3000 -v /app/node_modules -v $(pwd):/app  mikedonnici/myapp
```

Now, when the source code is changed the changes are propagated into the container. For React/Vue apps the hot reload 
feature will work when code is updated.

The presence of `node_modules` in the working directory will also slow the container build down. To avoid this we can 
 delete `node_modules` in the local working directory and add another `-v` flag to `docker run`. 
        
```bash
$ rm -fr node_modules
$ docker build -t mikedonnici/myapp -f Dockerfile.dev .
$ docker run -p 3000:3000 -v /app/node_modules -v $(pwd):/app  mikedonnici/someimage
```

As `npm install` is run in the container (when it is built) we don't need a copy in our local working directory. 
However, if we only map the volume with `-v $(pwd):/app` we get an error because there is no `node_modules` in the 
working directory.

The `-v /app/node_modules` directive tells docker to leave that folder in place, in the container, and map the rest. 
Thus all of the files are in place to start up the app.   

Note: `docker-compose.yml` for the above looks like this:

```yaml
version: '3'
services:
  web-app:
    build: 
      context: .
      dockerfile: Dockerfile.dev
    ports:
      - "3000:3000"
    volumes:
      - /app/node_modules
      - .:/app
```

Note: Strictly speaking the `COPY . .` is not required in `Dockerfile.dev` with the volume mapping. 
          
## docker-compose

Looks for `docker-compose.yml` in current dir.

### Restart Policies

* "no" - never (requires quotes as yml no = false)
* always - regardless of stop reason
* on-failure - if stopped with error code
* unless-stopped - always restart unless stopped by dev

### Example

```yaml
version: '3'
services:
    redis-server:
      image: 'redis'
    go-app:
      restart: on-failure
      build: .
      ports:
        - "8088:8088"
```

Start services:
```sh
$ docker-compose up
```

Start services (rebuild)
```sh
$ docker-compose up --build
```

Start service in background:
```sh
$ docker-compose up -d
```

Shutdown services (background)
```sh
$ docker-compose down
``` 

Running containers
```sh
docker-compose ps
```


## Referencea and Resources

* [Docker Cheatsheet](/dev/dcoker/Docker-CheatSheet-08.09.2016-0.pdf)
* https://www.bretfisher.com/docker/



 
