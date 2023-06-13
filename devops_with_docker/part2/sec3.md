# Volumes in action
- how to mount?: bind mount vs easy-to-locate
- docker compose uses the current directory as a prefix for container and volume names by default, if not specified(`COMPOSE_PROJECT_NAME`)
- `restart` option(`unless-stopped`)
- checking docker volumes: `docker volumn ls`

- explicitly specify the volume:

```yml
services:
  db:
    volumes:
      - database:/var/lib/postgresql/data

volumes:
  database:
```

- `depends_on` option: the container is relying on this service -> so start later
