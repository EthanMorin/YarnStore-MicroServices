codegen:
	oapi-codegen \
	-generate gin,spec \
	-package api config/catalog_api.yml > api/api.gen.go
	oapi-codegen \
	-generate types \
	-package models config/catalog_api.yml > models/models.gen.go
	go mod tidy