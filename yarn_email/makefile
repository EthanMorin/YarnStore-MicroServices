codegen:
	oapi-codegen \
	-generate gin,spec \
	-package api config/email_api.yml > api/api.gen.go
	oapi-codegen \
	-generate types \
	-package models config/email_api.yml > models/models.gen.go
	go mod tidy