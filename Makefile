# Parent makefiles at https://gitlab.intercloud.fr/intercloud/io/templates/make

-include .make/golang.mk

MAKE_BRANCH = main
MAKE_REPO   = ssh://git@gitlab.intercloud.fr:10022/intercloud/io/templates/make.git
GOARGS      = .env

make-import: # Install parent makefiles
	$(title)
	@rm -rf .make
	@git clone --branch $(MAKE_BRANCH) $(MAKE_REPO) .make

.DEFAULT_GOAL :=
default: lint

clean: go-clean # Clean generated and cached files
lint:  go-lint  # Check Go code
build: go-build # Build server
run:   go-run   # Run server
consul:         # Start consul server
	$(title)
	@docker-compose up
