codegen:
	oapi-codegen \
	-generate gin,spec \
	-package main config/orders_api.yml > api.gen.go
	oapi-codegen \
	-generate types \
	-package main config/orders_api.yml > models.gen.go
	go mod tidy