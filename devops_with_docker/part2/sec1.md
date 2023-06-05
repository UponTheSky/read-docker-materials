# Migrating to Docker Compose
- **Docker Compose**?: designed to simplify running multi-container applications to using a single command

- syntax:

```yaml
version: '3.8'

services:
  service-name:
    image: <user-name>/<repo-name>
    build: . # see: https://docs.docker.com/compose/compose-file/build/
    volumes:
      - .:/mydir
    container_name: <service-name>
```

# Volumes in Docker Compose
- syntax: `location-in-host:location-in-container
- container name: either `container_name` or the service name

# Web services in Docker Compose

