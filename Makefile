build:
	go build -o bin/pricefetch

run: build
	./bin/pricefetch