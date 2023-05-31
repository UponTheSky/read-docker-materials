# Interacting with the container via volumes and ports
- **volume**: to store data outside the container
- but with **bind mount**, we can mount a file or directory from our own machine(host machine) into the container
- `-v` option: 
  - bind mount: `docker run -v <host_directory>:<container_dir> <image>`(you can also specify a file instead of a directory - but if the file doesn't exist, docker creates a directory by default)
  - for volume mount, replace `<host_directory>:<container_dir>` with `<volume_name>:<container_dir>`

- a volume is simply a directory/file that is shared between the host machine and the container
  - for saving data persistently
  - for sharing files bewteen containers
  - for running programs that are able to load changed files

## ex 1.9

```
touch ./text.log
docker run -v "$(pwd)/text.log:/usr/src/app/text.log" --name ex1-9-c2 devopsdockeruh/simple-web-service
```

# Allowing external connections into containers
- mapping the host machine port(publishing port) to a container port(the port the container listens to)
  - exposing a container port: `EXPOSE <port>` in dockerfile
  - publishing a port: `-p <host_port>:<container_port>`
  - (remark): `<host_port>` is actually `<ip>:<port>`, where `<ip>` is `0.0.0.0` by default

## ex 1.10

```
docker run --name ex1-10-v2 -p 3000:8080 devopsdockeruh/simple-web-service server

curl http://127.0.0.1/I-will-be-a-billinaire
```


