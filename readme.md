Configuration
-

Before you run the apps, make sure you already install the tech stack.

*requirement*

- mysql
- redis
- mongodb
- golang v1.2x
- Kafka 2.11

Copy the configuration file example.yaml to env.yaml

- run `cp example.yaml env.yaml`
- make an adjustment to your local configurations

HTTP Server
-
After the environment file already being setup, you can run the http server by running this command ``go run src/cmd/http/main.go``

Kafka Consumer Server
-
Run the kafka consumer by running this command ``go run src/cmd/kafka/consumer/main.go``


Migration
-

These apps already build with a migration command. to set up the migration you can follow this step:

- build the migration by running this command `go build -o migration src/cmd/migrate/main.go`
- for migration up use `./migration up`
- for migration down use `./migration down`
- for create a new migration `./migration create [migration name]`
- for check migration version `./migration version`
- for help `./migration help`