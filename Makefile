KEYCLOAK_OPENAPI_SPEC=https://raw.githubusercontent.com/ccouzens/keycloak-openapi/main/keycloak/20.0.3.yml

.PHONY: build
build: generate
	go build ./v20

.PHONY: deps
deps:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0

.PHONY: generate
generate: deps
	oapi-codegen -package keycloak -generate client ${KEYCLOAK_OPENAPI_SPEC} > ./v20/client.go
	oapi-codegen -package keycloak -generate types ${KEYCLOAK_OPENAPI_SPEC} > ./v20/types.go
