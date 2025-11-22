# Docker

<br>
<br>
<br>
<br>
<br>

## Index

- [Foundations](#foundations)
  - [What is Docker?](#what-is-docker)
  - [Containers vs VMs](#containers-vs-vms)
  - [How Docker Works Internally](#how-docker-works-internally)
  - [Why Docker?](#why-docker)
  - [Advantages](#advantages)
  - [Disadvantages / Limitations](#disadvantages--limitations)
  - [Core Docker Concepts](#core-docker-concepts)
  - [Putting It All Together (Mental Model)](#putting-it-all-together-mental-model)
- [Dockerfile](#dockerfile)
  - [What is a Dockerfile?](#what-is-a-dockerfile)
  - [Core Dockerfile Instructions](#core-dockerfile-instructions)
  - [Multi-Stage Builds](#multi-stage-builds)
  - [Dockerfile Best Practices](#dockerfile-best-practices)
  - [Building & Managing Images](#building--managing-images)
- [Working with containers](#working-with-containers)
  - [Core container commands](#core-container-commands)
  - [Container states & Restart policies](#container-states--restart-policies)
  - [Networking Concepts](#networking-concepts)
  - [Useful Commands Summary](#useful-commands-summary)

<br>
<br>
<br>
<br>
<br>

## Foundations

### What is Docker?

- Docker is a tool to build, package, ship, and run applications inside lightweight, isolated environments called **Containers**.
- A container bundles our application + runtime + dependencies.
- It shares the host kernel, making it extremely fast and lightweight compared to VMs.

<br>
<br>

### Containers vs VMs

| Aspect    | Containers                                | VMs                                        |
| --------- | ----------------------------------------- | ------------------------------------------ |
| Isolation | Process-level                             | Full OS                                    |
| Kernel    | Shared with host                          | Own guest kernel                           |
| Size      | MBs                                       | GBs                                        |
| Startup   | Seconds / milliseconds                    | Minutes                                    |
| Use case  | Microservices, CI/CD, ephemeral workloads | Heavy workloads, full OS, strong isolation |

<br>
<br>

### How Docker Works Internally

#### Images

- Read-only templates containing application code and dependencies.
- Built using a **Dockerfile**.

<br>

#### Layers

- Each Dockerfile instruction produces a layer.
- Layers are cached and reused across builds, resulting in faster builds.

<br>

#### Containers

- A Docker container is a running instance of a Docker image.
- It is the actual execution environment where the application runs.

<br>
<br>

### Why Docker?

- **Consistency**: Same environment across dev, staging, prod.
- **Portability**: Runs anywhere: your laptop, server, cloud.
- **Reproducibility**: Deterministic builds with Dockerfiles.
- **Isolation**: Dependencies don’t conflict with host or other apps.
- **Speed**: Faster than VMs; ideal for microservices, local setups.
- **Scalability**: Works well with orchestration tools (Kubernetes, ECS).

<br>
<br>

### Advantages

- Lightweight & fast startup.
- Easy to version, distribute, and replicate environments.
- Immutable, declarative infrastructure (Dockerfile).
- Eliminates 'works on my machine' issues.
- Great for microservices, CI pipelines, reproducible builds.

<br>
<br>

### Disadvantages / Limitations

- Weaker isolation than VMs (shared kernel).
- Not ideal for GUI apps or full OS-level virtualization needs.
- Can become messy without cleanup (dangling images, volumes).
- Learning curve (images, networks, volumes, compose).
- Security vulnerabilities if images are not updated/scanned.

<br>
<br>

### Core Docker Concepts

#### Image

- Blueprint / template for containers.
- Immutable.
- Created from a `Dockerfile`.

<br>

#### Container

- Running instance of an image.
- Has its own filesystem, process tree, network namespace.

<br>

#### Registry

- Storage for docker images (E.g. Docker Hub).

<br>

#### Entrypoint vs CMD

- **CMD**: Sets default command which can be **overridden** at runtime.
- **ENTRYPOINT**: Defines the main executable which is harder to override.
- Often used together:
  ```dockerfile
  ENTRYPOINT ["python"]
  CMD ["app.py"]
  ```

<br>

#### Volumes & Bind Mounts

- **Volumes**: Managed by Docker, for persistent data.
- **Bind mounts**: Link host directory to container directory.

<br>

#### Networks

- Containers communicate over virtual networks.
- Docker provides internal DNS which allows containers to use service names.

<br>
<br>

### Putting It All Together (Mental Model)

- A **Dockerfile** builds a **layered image**.
- Running the image starts a **container** with isolation.
- Containers communicate using **Docker networks** and store data using **volumes**.
- Images live in **registries** and are portable.

<br>
<br>
<br>
<br>
<br>

## Dockerfile

### What is a Dockerfile?

- A Dockerfile is a declarative recipe for building Docker images.
- It describes:
  - Base image
  - Dependencies
  - Environment setup
  - Application code
  - Startup command
- Docker builds the image top-to-bottom, caching each instruction as a layer.

<br>
<br>

### Core Dockerfile Instructions

#### FROM

- Sets the base image.
- Always the first instruction.
- Should use minimal base images for smaller, secure images.
- Can have multiple FROM statements (**multi-stage builds**).

```dockerfile
FROM python:3.10-slim
```

<br>

#### WORKDIR

- Sets the working directory inside the image.
- All subsequent commands run relative to this directory.
- Creates the directory if it doesn’t exist.

```dockerfile
WORKDIR /app
```

<br>

#### USER

- Runs the container as a non-root user.
- Best practice for security.

```dockerfile
RUN useradd -m appuser
USER appuser
```

<br>

#### COPY

- Copy files/folders from host to image.

```dockerfile
COPY requirements.txt .
```

<br>

#### ADD

- Does everything `COPY` does plus:
  - Can download URLs
  - Can auto-extract local archives
- **Best practice**: Use `COPY` unless `ADD`'s special features are needed (rare).

```dockerfile
# extracts app.tar.gz into /usr/src/app/
ADD app.tar.gz /usr/src/app/
```

<br>

#### RUN

- Executes commands (during build, not runtime).
- Produces a new layer.
- Should combine related commands to reduce layers.

```dockerfile
RUN apt-get update && apt-get install -y gcc
```

<br>

#### ENV

- Set environment variables in the image.

```dockerfile
ENV APP_ENV=production
```

<br>

#### EXPOSE

- Document the port the container listens on.
- **Does not** publish ports outside.
- Only informational / for tools like docker-compose.

```dockerfile
EXPOSE 8000
```

<br>

#### HEALTHCHECK

- Define a container health probe.
- Used by orchestration (compose/k8s) to restart unhealthy containers.

```dockerfile
HEALTHCHECK CMD curl --fail http://localhost:8000/health || exit 1
```

<br>

#### CMD

- Default command when a container runs.
- **Overridden** if you pass a command in `docker run`.

```dockerfile
CMD ["python", "app.py"]
```

<br>

#### ENTRYPOINT

- Defines the main executable that **should not** be overridden.
- ENTRYPOINT = fixed part
- CMD = arguments
- Common pattern: ENTRYPOINT + CMD combination.

```dockerfile
ENTRYPOINT ["uvicorn"]
CMD ["app:app"]
```

<br>
<br>

### Multi-Stage Builds

- Used to:
  - Reduce image size
  - Separate build + runtime environments
- Benefits:
  - Final image contains only the binary → tiny, fast, secure.
  - No compilers or build dependencies in production.

Example: Multi-stage build for a Go application

```dockerfile
# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o server .

# Runtime stage
FROM debian:bookworm-slim
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

<br>
<br>

### Dockerfile Best Practices

#### Use small base images

- python:3.10-slim
- debian:slim
- alpine (only when compatible)

<br>

#### Order instructions for caching

- Place least frequently changed steps at top:

```dockerfile
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
```

<br>

#### Minimize number of layers

- Combine related commands:

```dockerfile
RUN apt-get update \
 && apt-get install -y curl \
 && rm -rf /var/lib/apt/lists/*
```

<br>

#### Avoid unnecessary files inside the image

- Use `.dockerignore`.

<br>

#### Do not run as root

- Add a user with limited privileges.

<br>

#### Keep ENTRYPOINT predictable

- Good pattern:

```dockerfile
ENTRYPOINT ["python"]
CMD ["app.py"]
```

<br>

#### Use multi-stage builds for production

- Especially for Go, Rust, Node, Python.

<br>
<br>

### Building & Managing Images

#### Build an image

```bash
docker build -t myimage:latest .
```

- `-t` sets name:tag.
- Uses the Dockerfile in the current directory by default.

<br>

#### Build with no cache

```bash
docker build --no-cache -t myimage:latest .
```

- Rebuilds every layer from scratch.
- Useful when debugging or forcing fresh dependencies.

<br>

#### Use BuildKit (modern, faster builds)

```bash
DOCKER_BUILDKIT=1 docker build -t myimage:latest .
```

- Faster, parallel builds + better caching.
- Supports advanced features (secrets, inline cache).

<br>

#### Inspect an image

```bash
docker inspect myimage:latest
```

- Shows full metadata in JSON.
- Use `--format` to extract specific fields.

<br>

#### List images

```bash
docker images
```

- Shows repository, tag, size, image ID.
- Add `-a` to show intermediate images.

<br>

#### Remove an image

```bash
docker rmi myimage:latest
```

- Use `-f` to force removal.
- Image cannot be removed if containers depend on it.

<br>

#### Push to registry

```bash
docker push myrepo/myimage:latest
```

- Requires docker login first.
- Image must be tagged with the registry name.

<br>
<br>
<br>
<br>
<br>

## Working with containers

### Core container commands

#### Run a container

```bash
docker run <image>

docker run -d --name myapp -p 8000:8000 myimage
```

- `-d`: Run in detached mode.
- `--name`: Container name.
- `-p host:container`: Port mapping.
- `-e KEY=VALUE`: Environment variables.
- `-v host:container`: Volume mounting.
- `--network`: Attach to network.
- `--restart`: Restart policy.

<br>

#### Other important commands:

View running containers:
```bash
docker ps
```

<br>

View all containers (including stopped)
```bash
docker ps -a
```

<br>

Stop a running container
```bash
docker stop <containerName>
```

<br>

Start a stopped container
```bash
docker start <containerName>
```

<br>

Remove a container
```bash
docker rm <containerName>
```

<br>

View container logs
```bash
docker logs <containerName>
```

<br>

View container logs (with follow)
```bash
docker logs -f <containerName>
```

<br>

Execute inside a running container (SSH)
```bash
docker exec -it <containerName> bash
```

<br>

Copy files between host and container
```bash
# host to container
docker cp ./local.txt <containerName>:/app/text.txt

# container to host
docker cp <containerName>:/app/text.txt ./local.txt
```

<br>

Inspect container metadata
```bash
docker inspect <containerName>
```

- Useful for inspecting:
  - IP address
  - Mounted volumes
  - Environment variables
  - Network information
  - Entrypoint / Cmd

<br>
<br>

### Container states & Restart policies

#### Container states

- **Created**: Container is created but not yet started.
- **Running**: Container is actively executing its process.
- **Paused**: All processes in the container are temporarily suspended. `docker pause <containerName>`.
- **Restarting**: Container is restarting due to a configured restart policy.
- **Exited**: Container has stopped after completing its main process.
- **Dead**: Container failed to stop properly and is in an unusable state.

<br>

#### Restart policies

- Restart policy is set for a container using the `--restart` flag.
- Most commonly used restart policies are:
  - `no`: Never restart the container (default).
  - `on-failure`: Restart only if the container exits with a non-zero (error) status.
  - `on-failure:N`: Restart on failure but limit retries to N times.
  - `always`: Always restart the container regardless of exit status.
  - `unless-stopped`: Restart the container unless it was manually stopped by the user.

<br>
<br>

### Networking Concepts

#### Default bridge network

- The default bridge network is a built-in network that Docker creates automatically (named `bridge`) when you install Docker.
- Containers attached to this network can communicate with each other using their IP addresses.
- These containers are isolated from the host unless ports are explicitly published.

Key points:

- **Pre-created network**: Docker automatically provides a network named `bridge`.
- **Container-to-container communication**: Containers can talk to each other only via IP addresses, not by container name.
- **Basic isolation**: Containers are isolated from the host and external networks unless you publish ports using `-p`.
- **Advantage**: Suitable for testing or small setups, but not recommended for production.
- **User-defined bridge is better**: Creating your own bridge network allows **name-based DNS**, better isolation, and more control.

<br>

Command to list all networks:
```bash
docker network ls
```

<br>

#### User-defined bridge network

- A user-defined bridge network is a custom bridge network that you create in Docker.
- It provides better isolation, improved container discovery, and more control compared to the default bridge network.

Key points:

- **Custom-created**: You create it using `docker network create <name>`.
- **Built-in DNS**: Containers can communicate using container names instead of IPs.
- **Better isolation**: Containers in a user-defined network are isolated from containers in other networks unless connected explicitly.
- **Automatic service discovery**: Docker’s embedded DNS automatically resolves container names.
- **Fine-grained control**: You can configure subnet, gateway, IP ranges, and network options.
- **Enhanced security**: Traffic stays within that network; easier to apply network-level rules and policies.

<br>

#### Host network

- The host network makes a container share the host machine’s network stack directly.
- This means the container does not get its own IP, instead it uses the host’s network interface as if it were running directly on the host.
- `docker run --network=host myimage`

Key points:

- **No network isolation**: The container uses the host’s network namespace directly.
- **No port mapping needed**: Services inside the container are accessible on the host’s IP and ports without -p publishing.
- **Higher network performance**: Eliminates virtual network overhead; useful for high-performance or low-latency workloads.
- **Potential port conflicts**: Since container ports == host ports, conflicts can occur.
- **Reduced security isolation**: Less isolation compared to bridge networks; increases exposure if misconfigured.

<br>

#### Port Publishing

- Use to expose container's port to the host.
- `docker run -p <hostPort>:<containerPort> myimage`.
- Allows host / external client to access services running inside the container.

<br>

#### Expose vs Publish

- `EXPOSE 8000` (Dockerfile): Expose command is for documentation only.
- `-p 8000:8000`: Actual port forwarding while running the container.

<br>

#### Inspecting networks

- `docker network inspect <networkName>`.
- Shows:
  - Attached containers
  - Their IPs
  - Subnet & gateway

<br>
<br>

### Useful Commands Summary

| Action               | Command                          |
| -------------------- | -------------------------------- |
| Run container        | `docker run -d image`            |
| Shell into container | `docker exec -it container bash` |
| List running         | `docker ps`                      |
| List all             | `docker ps -a`                   |
| Stop                 | `docker stop container`          |
| Start                | `docker start container`         |
| Remove container     | `docker rm container`            |
| Logs                 | `docker logs -f container`       |
| Copy                 | `docker cp`                      |
| Inspect              | `docker inspect`                 |
| List images          | `docker images`                  |
| Remove image         | `docker rmi image`               |
| List networks        | `docker network ls`              |
| Inspect network      | `docker network inspect`         |
| Create network       | `docker network create`          |

<br>
<br>
<br>
<br>
<br>

## 

<br>
<br>
<br>
<br>
<br>
