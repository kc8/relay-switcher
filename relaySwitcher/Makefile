default: build

test: 
	go test ./...

build:
	mkdir -p dist
	go build -v -o dist/relaySwitcher

lint:
	golangci-lint run ./...
	cd terraform/ && golangci-lint run ./...

clean:
	rm -rf dist/

deploy: default build
	ansible-playbook -i ../scripts/ansible_settings.ini ../scripts/deploy_relaybackend.yaml

.PHONY: prepare test build install e2e lint clean generate
