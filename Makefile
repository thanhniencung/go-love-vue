web:
	cd ./web-ui/ && vite build
	cd ./cmd/web && go run main.go