.PHONY: api
api:
	# go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g api.go -d internal
	go build -o api api-layout/cmd