default: run

fmt:
    go fmt ./...

install_deps:
    go get .

build: install_deps fmt
    go build -o build/mdn

run: build
    build/mdn --help
    build/mdn --update
    build/mdn html
