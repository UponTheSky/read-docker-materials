# Docker networking
- **Docker network**: connecting several services using docker compose
- when starting the services, docker compose automatically creates and joins both containers into a network with a DNS
    - other than the Frontend, ports don’t have to be exposed outside of the network
- each container can be referenced with its service name
