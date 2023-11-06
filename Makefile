build:
	go build -o bin/pricefetch

run: build
	./bin/pricefetch

gen-proto:
	cd ./proto && sh gen.sh

.PHONY: proto

