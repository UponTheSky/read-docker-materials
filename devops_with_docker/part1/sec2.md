# Running and stopping containers
- example: ubuntu container
  - need to add `-t` flag to interact with the container, by creating a tty
  - also, need to add `-i` flag to send our STDIN to the container
  - (covered in sec1.md) `-d` is for running the container in detached state

- other commands
  - `docker logs -f <container>`: follows log output
  - `docker attach <container>`: attach local stdio streams to `<container>`
    - `--no-stdin`: not attaching the stdin
  - `docker exec [options] container <command> [args]`
    - `-it`: allocate a pseudo-tty, and keep stdin open even if not attached
    - if only `-i`: since we don't have tty, we can't interact with the container(paused)
    - if only `-t`: no stdin => the terminal shows up but no command works

- docker run vs start / docker stop vs pause, kill?
  - run vs start vs create: https://linuxhandbook.com/docker-run-vs-start-vs-create/
  - stop vs pause: https://signoz.io/blog/docker-container-lifecycle/#stopped-state
  - stop vs kill: https://signoz.io/blog/docker-container-lifecycle/#killed--deleted-state

## ex1.3

```sh
docker pull devopsdockeruh/simple-web-service:ubuntu
docker run -d --name ex1-3 devopsdockeruh/simple-web-service:ubuntu
docker exec -it ex1-3 bash
(in bash)tail -f ./text.log
```

The secret message is: 'You can find the source code here: https://github.com/docker-hy'

# Ubuntu in a container is just... Ubuntu
- the container behaves just like the normal Ubuntu
- but the storage is not permanent(it is removed once we stop the container)

## ex1.4

```sh
docker pull ubuntu
docker run -d -it --name ex1-4 ubuntu
docker exec -it ex1-4 bash
bash > apt-get update; apt-get -y install curl
bash > exit
docker exec -d -it ex1-4 sh -c 'while true; do echo "Input website:"; read website; echo "Searching.."; sleep 1; curl http://$website; done'
```

- the more neat solution is: https://github.com/oneiromancy/devops-with-docker#14-missing-dependencies

