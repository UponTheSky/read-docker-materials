# In-depth dive to images
- **image**: a basic building block for containers and other images
- "containerize" = creating an image

# Where do the images come from?
- pulling from the Docker Hub the public images: `docker search <image>`
  - official: curated and reviewed by Docker, Inc.
  - automated: the image is automatically built from the source repo
  - other than the Docker Hub is also possible to use(like quay)

# A detailed look into an image
- `docker pull <registry>/<organization>/<image>:<tag>`
  - `<registry>`: default is Docker Hub
  - `<organization>`: default is library
  - unless `<tag>` is specified, the default tag is `latest`
  - layers in an image: downloaded in parallel

- tag: a way to rename an image: `docker tag ubuntu:18.04 foo`

## ex 1.5
- 83MB(ubuntu) vs 15.7MB(alpine)
- (review)

```
docker run -d --name ex1-5 devopsdockeruh/simple-web-service:<tag>
docker exec -it ex1-5 sh
ls
cat text.log
exit
```

## ex 1.6
- "This is the secret message"
- "basics"

# Building images
- **Dockerfile**: a file of the build instructions for an image
- dependencies: small images are important(like alpine)
- fixed version, tags: no breaking changes + what exact problems does the current image have

- docker build
```
docker build <docker_dir_path> -t <tag_name> -f <file_path>
```

- **layer**: a separate **step** for a single image to be built, acting like a cache for fast builds

- **intermediate container**: containers created from the image, in which the command is executed
  - the changed state is stored into an image

- commands:
  - **RUN**: during the build time
  - **CMD**: executed when the container starts to run(unless overwrited)
