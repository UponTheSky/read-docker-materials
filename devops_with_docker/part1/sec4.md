# Defining start conditions for the container
- we should keep the most prone to change rows at the bottom, so that we can preservce our cached layers => faster subsequent builds

- `docker run [options] <image> <command>`
  - when the container starts to run, the `<command>` is executed
  - in the `Dockerfile`, we have `CMD`, which acts as a default value for this `<command>`; That is, `CMD` is overwritten by `<command>` when you run the docker container

- `ENTRYPOINT`: we need to have a certain command before the `<command>`  
  - if not given, by default it is set as `/bin/sh`
  - if it is set, `CMD` is used to give default arguments to the entrypoint(of course, it could be overriden when we pass `<command>` argument in `docker run`)


- **exec** form and **shell** form
  - the two ways to set `ENTRYPOINT` and `CMD`
  - shell form
    - the command that is executed is wrapped with `/bin/sh -c`
    - the comamnd is provided as a string without brackets
  - exec form
    - the command and the arguments are provided as a list

- but most of time, we ignore `ENTRYPOINT`, since by default it is `/bin/sh`

# Improved curler

