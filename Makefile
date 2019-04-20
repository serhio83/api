.PHONY: all dep build config clean
all: dep build config
dep:
	dep ensure -v
build:
	go build -o api
config:
	cp cfg/config.toml .
clean:
	rm -rf api config.toml vendor
