# AUTHOR: Don B. Stringham <$(REPO_OWNER)@icloud.com>
.DEFAULT_GOAL=clean

# VARIABLES
CC=go
ALL_PACKAGES ?= $(shell go list ./...)
APP_NAME ?= gsh
BUILD_TIME ?= $(shell date +%FT%T%z)
REPO_OWNER ?= $(shell cd .. && basename "$$(pwd)")
VERSION ?= $(shell git describe --tags 2>/dev/null)

# TARGETS
bld:
	go install $(ALL_PACKAGES)

bld.cli: clean
	@echo Building $(VERSION)...
	-mkdir -p ./bld
	# GOOS=linux GOARCH=amd64 $(CC) build -a -o ./bld/$(APP_NAME).linux-amd64 -ldflags='-s -w -X github.com/$(REPO_OWNER)/$(APP_NAME)/ver.Version=$(VERSION) -X github.com/$(REPO_OWNER)/$(APP_NAME)/ver.Buildtime=$(BUILD_TIME)' ./
	GOOS=darwin GOARCH=amd64 $(CC) build -a -o ./bld/$(APP_NAME).darwin-amd64 -ldflags='-s -w -X github.com/$(REPO_OWNER)/$(APP_NAME)/ver.Version=$(VERSION) -X github.com/$(REPO_OWNER)/$(APP_NAME)/ver.Buildtime=$(BUILD_TIME)' ./
	# GOOS=windows GOARCH=amd64 $(CC) build -a -o ./bld/$(APP_NAME).windows-amd64.exe -ldflags='-s -w -X github.com/$(REPO_OWNER)/$(APP_NAME)/ver.Version=$(VERSION) -X github.com/$(REPO_OWNER)/$(APP_NAME)/ver.Buildtime=$(BUILD_TIME)' ./

	cd ./bld && find . -name 'final*' | xargs -I{} tar czf {}.tar.gz {}
	cd ./bld && shasum -a 256 * > sha256sum.txt
	cat ./bld/sha256sum.txt

clean:
	-rm -r ./bld

info:
	@echo VERSION: $(VERSION)
	@echo REPO_OWNER: $(REPO_OWNER)
	@echo ALL_PACKAGES: $(ALL_PACKAGES)

test:
	-rm -r .coverage
	@mkdir -p .coverage
	make test.with.flags TEST_FLAGS='-v -race -covermode atomic -coverprofile .coverage/_$$(RAND).txt -bench=. -benchmem'
	@echo 'mode: atomic' > .coverage/combined.txt
	@cat .coverage/*.txt | grep -v 'mode: atomic' >> .coverage/combined.txt

test.with.flags:
	@go test $(TEST_FLAGS) .
	@go test $(TEST_FLAGS) ./cmd/...
	@go test $(TEST_FLAGS) ./pkg/...
	@go test $(TEST_FLAGS) ./ver/...

coverage.html:
	go tool cover -html=.coverage/combined.txt

docs:
	-make kill-docs
	nohup godoc -play -http=127.0.0.1:6064 </dev/null >/dev/null 2>&1 & echo $$! > .godoc.pid
	cat .godoc.pid

docs.kill:
	@cat .godoc.pid
	kill -9 $$(cat .godoc.pid)
	rm .godoc.pid

docs.open:
	open http://localhost:6064/pkg/github.com/$(REPO_OWNER)/final

push:
	git add .
	git status
	git commit -m "Fixed bugs and updated files"
	git push -u origin master

# example: make release V=0.0.0
release:
	git tag v$(V)
	@read -p "Press enter to confirm and push to origin ..." && git push origin v$(V)

.PHONY: bld.cli clean clean.vendor info test test.with.flags coverage.html \
        release docs kill-docs open-docs

SHELL = /bin/bash

