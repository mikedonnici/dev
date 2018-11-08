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



 
