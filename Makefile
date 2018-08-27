
DEP_VERSION=0.5.0

dependency:
	curl -fsSL -o ${GOPATH}/bin/dep https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64
	chmod +x ${GOPATH}/bin/dep
	dep ensure

unittest:
	go test ./... -cover

clean:
	echo Removing all vendoring files...
	rm -rf vendor/


deps:
	dep ensure -update
