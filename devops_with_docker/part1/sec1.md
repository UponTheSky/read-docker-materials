# Definitions and basic concepts
- https://devopswithdocker.com/part-1/section-1

# What is DevOps?
- Dev + Ops: developers manage the operations(release, configuring, monitoring, etc.) as well

- By Jabbari et al.:

> "DevOps is a development methodology aimed at bridging the gap between Development and Operations, emphasizing communication and collaboration, continuous integration, quality assurance and delivery with automated deployment utilizing a set of development practices".

- sometimes it means a SWE role

# What is Docker?
- definition
  1. a set of tools to deliver SW in containers
  2. Containers are packages of SW

- one container is isolated from the others

# Benefits from containers

## Scenario 1: Works on my machine
- you won't realize what dependencies an application requires if you only develop it on your machine
- solution: container includes all the dependencies required

## Scenario 2: Isolated environments
- some parts of the system may change over time => making application not working
- isolated environment is not affected from such changes

## Scenario 3: Development
- installing third party applications could be burdensome
- docker makes it easy to do

## Scenario 4: Scaling
- provide more instances(virtual scaling) and place the load balancer

# Virtual machines
- containers have a direct access to the OS kernel and resources
- docker uses Linux kernels

# Running containers
- read

# Image and containers
- container(food) = an instance of an image(recipe)

## Image
- a docker image is a file, consisting of a base image and additional layers on top of it
- it is immutable; you can't edit it once it is created
- an image file is built from an instructional file `Dockerfile`, which is parsed when youu run `docker image build`

## Container
- basically containers are isolated, but they can interact with each other, and the host machine

# Docker CLI basics
- the docker engine is made up of 3 parts: CLI / REST API / Docker daemon
- when we run a command, the client sends a request through the REST API to the Docker daemon
  - Docker daemon takes care of images, containers, and others

## Most used commands
- use this part as a quick reference
- https://devopswithdocker.com/part-1/section-1#most-used-commands

## ex1.1, 2
- this part is fairly easy

```sh
docker pull redis
docker run -d --name first redis
docker run -d --name second redis
docker run -d --name third redis

docker stop first second
docker ps -a >> ex1-1.txt

docker stop third
docker rm first second third
docker ps -a >> ex1-2.txt 
docker rmi redis
docker images >> ex1-3.txt
``` 
