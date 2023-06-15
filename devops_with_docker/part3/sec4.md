# Optimizing the image size
- why small image size?
    - takes less time to pull the image
    - security: bigger images have larger surfaces to be attacked

- **layer**:
    - each instruction in a Dockerfile creates a layer in the image
    - each layer potentially adds something extra to the resulting image
    - when you change the Dockerfile and rebuild the image, only those layers which have changed are rebuilt
    - it is good to maintain the number of layers small

- checking the image layers: `docker image history <your_image>`

## ex3.6: code omitted
- frontend: 532.57MB => 532.73(?!)
- backend: 644.53MB => 644.53MB

## Alpine Linux variant
- alpine: based on musl(alternative glibc implementation) and busybox binaries
    - `apk` instaed of `apt-get`
    - most of the package names are the same as ubuntu

## ex3.7
- frontend: 532.73MB => 470.15MB
- backend: 644.53MB => no change

## Multi-stage builds
- compiled languages require building tools, but we don't need them when actually running the production code
- we don't really need them so let's separate them 
- each `FROM` instruction uses a different base

```Dockerfile
FROM python:3.9 AS build-stage
...

FROM python:3.9-alpine
COPY --from=build-stage /usr/src/app/ /usr/src/app
```

## ex3.8
- frontend: 470.15MB => 128.98MB

## ex.3.9
- backend: => 18.04MB

## ex.3.10: skipped
