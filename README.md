# bouncer
The authentication service for mindtastic

## Requirements

Docker 20.10.14  
Docker Compose v2.5.0

## Setup instructions

This service uses Keycloak with a PostgreSQL database for user management. Spin up the containers using

    docker compose up

To run the production configuration, run

    docker compose -f docker-compose.prod.yml up

## Files

**Base configuration:**  
docker-compose.yml

**Overrides for dev and testing configuration** (applied by default):  
docker-compose.override.yml

**Overrides for production configuration:**  
docker-compose.prod.yml
