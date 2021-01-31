web:
	cd ./web-ui/ && vite build
	cd ./cmd/web && go run main.go

web-dev:
	cd ./web-ui/ && vite