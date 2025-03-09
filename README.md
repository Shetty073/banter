# banter
A social media app


## stack

- Phoenix (elixir / erlang)
- minio (S3 like storage)
- postgreSQL

> Using an API gateway such as AWS or Kong API gateway is recommended.

## For development

- install homebrew
- install Erlang and Elixir using brew
- install minio using brew
- install mongoDB
- install postgreSQL
- install direnv using brew (for configs)

### minio guide

- default creds upon installation minioadmin:minioadmin
- command to start service and provide a storage path: minio server /data
