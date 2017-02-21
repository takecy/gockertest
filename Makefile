.PHONY: update test

prepare:
	go get -u github.com/golang/dep

update:
	dep ensure -v -update

test:
	./test.sh