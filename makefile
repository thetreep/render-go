patch:
	git apply ./patches/*.patch

generate:
	oapi-codegen -package render openapi.yaml > render.go

build: patch generate