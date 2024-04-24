codegen:
	oapi-codegen \
	-generate gin,spec \
	-package main config/user_api.yml > api.gen.go
	oapi-codegen \
	-generate types \
	-package main config/user_api.yml > models.gen.go
	go mod tidy 