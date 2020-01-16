# REST api server to communicate with Keyclaok

## Run Keycloak Server

1. Install Keycloak server
2. Create admin creds and a Realm
3. Start Keycloak server

> OR Use Docker

```sh
docker run -p 8080:8080 -e KEYCLOAK_USER="ankit" -e KEYCLOAK_PASSWORD="admin" jboss/keycloak
```

> [Checkout](https://hub.docker.com/r/jboss/keycloak/) for more configuration in docker
