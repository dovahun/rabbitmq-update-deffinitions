# RABBITMQ-UPDATE-DEFINITIONS

## Made on Golang v1.24.7

### Example structure
    nexus-repost-state/
    ├── .bin/
    │   └── rabbitmq-update-definitions
    ├── src/
    │   ├── config.go
    │   ├── validate.go
    │   └── workWithApi.go
    ├── Dockerfile
    ├── schema.json
    ├── .dockerignore
    ├── .gitignore
    ├── main.go
    ├── go.mod
    ├── create_definitions.sh
    └── README.md

###  *Flags*
| Flags                   | Descripition                                             |
|:------------------------|:---------------------------------------------------------|
| -file                   | Path to deffinition file                                 |
| -host                   | Host api url for RabbitMQ with endpoint /api/definitions | 
| -user                   | User for basic auth to RabbitMQ                          |
| -password               | Password for user                                        |
| -update                 | Run update definitions                                   |
| -validate               | Run validate definitions file                            |


### How to build

go build .

### Example run
go run . -file ./definitions -host http://127.0.0.1:15672/api/definitions -user admin -password admin -update -validate 



