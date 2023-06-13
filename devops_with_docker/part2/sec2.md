# Docker networking
- **Docker network**: connecting several services using docker compose
- when starting the services, docker compose automatically creates and joins both containers into a network with a DNS
    - other than the Frontend, ports don't have to be exposed outside of the network
- each container can be referenced with its service name

# Manual network definition
- why manually?
    - easy to set up a configuration where containers defined in two different compose files share a network, interacting with each other

```yml
version: "3.8"

services:
    db:
        image: postgres:13.2-alpine
        networks:
            - database-network

networks:
    database-network:
        name: database-network
```

- accessing from another compose file:
```yml
version: "3.8"

services:
  db:
    image: backend-image
    networks:
      - database-network

networks:
  database-network:
    external:
        name: database-network
```

- by default, all services are added to a network called `default` - but you can configure it as well

```yml
version: "3.8"

services:
  db:
    image: backend-image

networks:
    default:
        external:
            name: database-network
```

# Scaling
- compose can scale up the service to run multiple instances:
    - `docker compose up --scale <service>=3`

- how about port conflict?: when leaving the host port unspecified, Docker will automatically choose a free port
    - so only specify the container port
    - find the ports using `docker compose port --index <index> <service> <container_port>`

- using nginx as a loadbalancer

## ex 2-5
- `docker compose up --scale compute=5`
