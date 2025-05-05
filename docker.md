## Definitions
- [Docker](#docker)
- [Docker Benefits](#docker-benefits)
- [Docker Image](#docker-image)
- [Docker Container](#docker-container)
- [Docker Engine](#docker-engine)
- [Docker Hub](#docker-hub)
- [Virtual Machine vs Container](#virtual-machine-vs-container)
- [Dockerfile](#dockerfile)
- [CMD vs ENTRYPOINT](#cmd-vs-entrypoint)

<br>
<br>
<br>

### Docker

- Docker is a tool to containerize our application.
- We collect our code and all of its dependencies and create a docker image.
- We can then run this image on any platform that supports docker and use our application.
- A running image (called container) is self sufficient to run our application.
- The application will perform similarly on all of these platforms.

<br>
<br>
<br>

### Docker Benefits

- No Setup: A docker image will contain all the code and required dependencies. To run the application, we just have to run the image inside a container.
- Consistency: As we all are using the same docker image, it avoid inconsistencies in the setup and make sure all of our applications perform identically.
- Scalability: If we want multiple instances of our application running, we just fire up multiple containers.
- Lightweight & Quick: Docker images are lightweight and easy to share. When you run this image, the container starts almost immediately.

<br>
<br>
<br>

### Docker Image

- A docker image is a read-only file that contains everything needed to run our application.
- It contains:
    - Source code
    - Dependencies
    - Environment Variables
    - Configurations
- A docker image is a blueprint to create a docker container.
- We use a `Dockerfile` to create a docker image.
- A docker image is composed of layers. Each instruction in a Dockerfile (like `RUN`, `COPY`, etc.) creats a new layer. This make images efficient and reusable.

<br>
<br>
<br>

### Docker Container

- A docker container is a running instance of a docker image.
- It is a live, executable environment where our application actually runs.
- Each docker container runs in isolation.
- We can go inside the container and execute commands as well.

<br>
<br>
<br>

### Docker Engine

- Docker engine is responsible for building & running containers on our system.
- It consists of:
  - Docker daemon: A background process that manages docker containers, images, network and volumes.
  - Docker CLI: A command line interface to inetract with docker daemon (`dockerd`)
- It is also responsible to connect containers to specific networks and volumes.

<br>
<br>
<br>

### Docker Hub

- Docker hub is a cloud-based repository where you can store & share docker images.
- It is like Github for docker images.

<br>
<br>
<br>

### Virtual Machine vs Container

**Virtual Machine**
- A virtual machine has it's own OS and uses the hardware of the host.
- It runs on a Hypervisor like `VMware`, `Hyper-V`.
- They require more CPU and memory, because of separate OS and thus have slower boot time.
- Has strong isolation (harware level).
- Suitable for running different OS like Linux VM on Windows.

<br>

**Container**
- A container shares the host's Operating System.
- It runs on a container engine like `Docker Engine`.
- Because it uses host's OS, it requires fewer resources and has quick startup.
- Has moderate isolation (process level).
- Suitable for web applications, cloud-native apps, etc.

<br>
<br>
<br>

### Dockerfile

- A dockerfile contains step by step instructions to build a docker image.
- It tell docker:
  - Which base image to pick?
  - Which code to copy?
  - How to install dependencies?
  - Which environment variables to set?
  - Which port to expose?
  - Entrypoint command, etc.

<br>
<br>
<br>

### CMD vs ENTRYPOINT

**ENTRYPOINT**

- It specifies the main command like `uvicorn`, `celery`, `nginx`, etc.
- It cannot be overridden when running the container (not recommended).
- ENTRYPOINT -> what to run.

<br>

**CMD**

- It specifies the default arguments that will be passed to the main command. Like `["main:app", "--port", "8080"]`.
- It can be overridden when running the container.
- CMD -> default arguments to pass to ENTRYPOINT.

<br>

```dockerfile
FROM ubuntu
ENTRYPOINT ["echo"]
CMD ["Hello from CMD"]
```

<br>

```bash
docker run myimage
# Output: Hello from CMD

docker run myimage "Hi there"
# Output: Hi there
```

- ENTRYPOINT stays in place; the arguments can come from CMD or docker run.

<br>
<br>
<br>

### COPY vs ADD

<br>
<br>
<br>

### Docker Volumes

<br>
<br>
<br>

### Dockerignore

<br>
<br>
<br>

### Docker Compose

<br>
<br>
<br>

### How to persist data in docker container?

<br>
<br>
<br>

### How can docker containers communicate with each other?

<br>
<br>
<br>

### How does networking work in docker?

<br>
<br>
<br>

### How do you optimize Dockerfile for smaller image size?

<br>
<br>
<br>

### What is a multi-stage build in Docker?

<br>
<br>
<br>
<br>
<br>

### Dockerfile

```dockerfile
# Use official Python image as a base
FROM python:3.11-slim

# Set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Set working directory
WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    libpq-dev \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# Install Python dependencies
COPY requirements.txt .
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# Copy project files
COPY . .

# Run the Django development server
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]
```

<br>
<br>
<br>
<br>
<br>

### Docker Compose file

```yaml
version: '3.9'  # Specify the Docker Compose file format version

services:
  web:  # This is the Django application container
    build: .  # Build the image from the Dockerfile in the current directory
    command: python manage.py runserver 0.0.0.0:8000  # Command to run Django development server
    volumes:
      - .:/app  # Mount current directory to /app in the container (live code reloading)
    ports:
      - "8000:8000"  # Map port 8000 on host to port 8000 on container
    depends_on:
      - db  # Wait for PostgreSQL to start before launching Django
      - redis  # Wait for Redis to start before launching Django
    environment:
      - DEBUG=1  # Enable Django debug mode
      - DATABASE_NAME=postgres  # Name of the PostgreSQL database
      - DATABASE_USER=postgres  # PostgreSQL username
      - DATABASE_PASSWORD=postgres  # PostgreSQL password
      - DATABASE_HOST=db  # Hostname of the PostgreSQL service (matches the db service name)
      - DATABASE_PORT=5432  # PostgreSQL default port
      - REDIS_HOST=redis  # Redis service hostname (matches redis service name)
      - REDIS_PORT=6379  # Redis default port

  db:  # This is the PostgreSQL database container
    image: postgres:15  # Use the official PostgreSQL 15 image
    environment:
      POSTGRES_DB: postgres  # Default database name
      POSTGRES_USER: postgres  # Default user
      POSTGRES_PASSWORD: postgres  # Default password
    volumes:
      - postgres_data:/var/lib/postgresql/data  # Persist database data between container restarts

  redis:  # This is the Redis service container
    image: redis:7  # Use the official Redis 7 image
    ports:
      - "6379:6379"  # Map port 6379 on host to container

volumes:
  postgres_data:  # Named volume to persist PostgreSQL data
```

<br>
<br>
<br>
<br>
<br>

### Commands

<br>

**Start, stop & remove container**

<br>

**Start, stop & remove container**

