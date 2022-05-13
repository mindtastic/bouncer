# bouncer
The authentication service for mindtastic

## Requirements

Docker 20.10.14  
Docker Compose v2.5.0

## Setup instructions

Building the React auth-app image and run containers simultaneously.

    docker-compose up -d --build

The react app can be reached using.

    localhost:3000
This service uses Keycloak with a PostgreSQL database for user management and react app for testing. Spin up the containers using

    docker compose up

To run the production configuration, run

    docker compose -f docker-compose.prod.yml up

The Keycloak admin console should then be reachable at

    localhost:8080


## Files

**Base configuration:**  
docker-compose.yml

**Overrides for dev and testing configuration** (applied by default):  
docker-compose.override.yml

**Overrides for production configuration:**  
docker-compose.prod.yml

## REST Endpoints

Documentation of built-in endpoints to communicate with Keycloak via REST API can be found in the api folder.
