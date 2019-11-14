PLATFORMS := linux/amd64 windows/amd64 linux/386 windows/386 darwin/amd64 freebsd/amd64 darwin/386 freebsd/386

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

install: 
	go get github.com/medisafe/jenkins-api/jenkins
	go get github.com/cheggaaa/pb/v3
	go get github.com/BurntSushi/toml

clean: 
	rm -rf bin

build:	clean $(PLATFORMS)

release: 
	install
	clean
	test
	build

test: 
	go test ./src/test/...

run: 
	go run jcli.go

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o 'bin/jcli-$(os)-$(arch)' src/main/jcli.go

.PHONY: build run clean install release test $(PLATFORMS)
