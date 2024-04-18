codegen:
	oapi-codegen \
	-generate gin,spec \
	-package main config/cart_api.yml > api.gen.go
	oapi-codegen \
	-generate types \
	-package main config/cart_api.yml > models.gen.go
	go mod tidy