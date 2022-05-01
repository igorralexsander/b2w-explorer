.PHONY: build-api
build-api: tests
	go build -o build/bin/b2w-explorer cmd/b2w-explorer/main.go

.PHONY: clean
clean:
	rm -rf build/bin
	go clean

.PHONY: run-api
run-api: build-api
	./build/bin/b2w-explorer

.PHONY: tests
tests:
	go test -v ./...

.PHONY: build-image
build-image:
	docker build -t b2w-explorer  -f  ./build/package/docker/Dockerfile .
   # docker tag b2w-explorer igorralexsander/b2w-explorer
    #docker push igorralexsander/b2w-explorer

.PHONY: tag-image
tag-image:
	docker tag b2w-explorer igorralexsander/b2w-explorer

.PHONY: push-image
push-image:
	docker push igorralexsander/b2w-explorer

.PHONY: build-deploy-image
build-deploy-image: build-image tag-image push-image