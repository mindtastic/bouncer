# bouncer
The authentication service for mindtastic

## Registration Proxy

This repository contains a tiny reverse proxy we call "bouncer". It intercepts incoming registration requests and 
generates a (cryptographically strong) random UUID-like string to replace the incoming username and password.
The incoming user credentials are discarded and never stored in any database.

## Local Setup

### Requirements

Docker 20.10.14  
Docker Compose v2.5.0

### Setup instructions

In production, this service uses Ory Kratos with a PostgreSQL database for user management.  
  
For local testing, you can use the standard package of Kratos with SQLite:

    docker compose -f quickstart.yml up --build -d

The Kratos public API will be reachable at

    localhost:4433

With the admin endpoint at

    localhost:4434

### Configuration

The Kratos configuration file can be found in /config/kratos.yml.

*Note:* The config file **must** be named kratos.yml.
