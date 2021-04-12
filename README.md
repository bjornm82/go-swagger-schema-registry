# Go Swagger Schema Registry

## Getting started

### Prerequisites

go >1.15 (https://golang.org/dl/)
mvn >3.6 (https://maven.apache.org/install.html)
go-swagger (https://github.com/go-swagger/go-swagger)

schema-registry running on localhost

### Workflow

* `git clone https://github.com/confluentinc/schema-registry`

* set env `REGISTRY_VERSION=v6.1.0`

* `git checkout tags/{REGISTRY_VERSION}`

* `mvn compile`

Should result in:

```
[INFO] --- spotbugs-maven-plugin:4.0.0:check (analyze-compile) @ kafka-schema-registry-benchmark ---
[INFO] BugInstance size is 0
[INFO] Error size is 0
[INFO] No errors/warnings found
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary for kafka-schema-registry-parent 6.1.0:
[INFO]
[INFO] kafka-schema-registry-parent ....................... SUCCESS [ 44.546 s]
[INFO] kafka-schema-registry-client ....................... SUCCESS [ 33.405 s]
[INFO] kafka-json-schema-provider ......................... SUCCESS [ 34.659 s]
[INFO] kafka-protobuf-provider ............................ SUCCESS [ 16.528 s]
[INFO] kafka-schema-serializer ............................ SUCCESS [ 24.441 s]
[INFO] kafka-avro-serializer .............................. SUCCESS [ 11.316 s]
[INFO] kafka-schema-registry .............................. SUCCESS [02:14 min]
[INFO] kafka-json-serializer .............................. SUCCESS [  8.427 s]
[INFO] kafka-connect-avro-data ............................ SUCCESS [ 36.130 s]
[INFO] kafka-connect-avro-converter ....................... SUCCESS [  5.863 s]
[INFO] kafka-schema-registry-package ...................... SUCCESS [  0.215 s]
[INFO] kafka-streams-avro-serde ........................... SUCCESS [ 10.212 s]
[INFO] kafka-json-schema-serializer ....................... SUCCESS [  4.878 s]
[INFO] kafka-connect-json-schema-converter ................ SUCCESS [  5.520 s]
[INFO] kafka-streams-json-schema-serde .................... SUCCESS [  2.942 s]
[INFO] kafka-protobuf-serializer .......................... SUCCESS [  5.647 s]
[INFO] kafka-connect-protobuf-converter ................... SUCCESS [  8.357 s]
[INFO] kafka-streams-protobuf-serde ....................... SUCCESS [  3.235 s]
[INFO] kafka-serde-tools-package .......................... SUCCESS [  5.333 s]
[INFO] maven-plugin ....................................... SUCCESS [  5.394 s]
[INFO] client-console-scripts ............................. SUCCESS [  0.145 s]
[INFO] schema-registry-console-scripts .................... SUCCESS [  0.152 s]
[INFO] kafka-schema-registry-benchmark .................... SUCCESS [  7.772 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  07:18 min
[INFO] Finished at: 2021-04-12T09:26:13+02:00
[INFO] ------------------------------------------------------------------------
```

At {PROJECT_DIR}/schema-registry/core/generated/swagger-ui/schema-registry-api-spec.yaml you should find the spec

* cp `schema-registry-api-spec.yaml` to `swagger.yaml` on current root path

* `swagger generate client`

* `go mod init`

* `go get ./...`

