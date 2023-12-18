.PHONY: full build build-npm test test-npm lint lint-npm fix fix-npm watch clean

SHELL=/bin/bash -o pipefail
$(shell git config core.hooksPath ops/git-hooks)

full: clean lint test build

## Build the project
build: build-npm

build-npm:
	[ -d node_modules ] || npm install
	npm run build

## Test the project
test: test-npm

test-npm:
	[ -d node_modules ] || npm install
	npm run test

## Lint the project
lint: lint-npm

lint-npm:
	[ -d node_modules ] || npm install
	npm run lint

## Fix the project
fix: fix-npm

fix-npm:
	[ -d node_modules ] || npm install
	npm run fix

## Watch the project
watch:

## Clean the project
clean:
	git clean -Xdff --exclude="!.env*local"
