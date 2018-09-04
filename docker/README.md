# Docker

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


        



 
