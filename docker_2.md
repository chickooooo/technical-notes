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

## 

<br>
<br>
<br>
<br>
<br>
