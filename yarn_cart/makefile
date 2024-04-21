codegen:
	oapi-codegen \
	-generate gin,spec \
	-package api config/cart_api.yml > api/api.gen.go
	oapi-codegen \
	-generate types \
	-package models config/cart_api.yml > models/models.gen.go
	go mod tidy 