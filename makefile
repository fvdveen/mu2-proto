all: gen-search

gen-search:
	protoc -I. --go_out=plugins=micro:go \
		proto/search/search.proto