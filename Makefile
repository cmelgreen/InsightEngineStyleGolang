build:
	go build -o functions/dynamic-css src/dynamicCSS/main.go 
	go build -o functions/cors-proxy src/corsProxy/main.go
	ls functions