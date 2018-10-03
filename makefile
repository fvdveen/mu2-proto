all: gen-search gen-encode

gen-search:
	protoc -I. --go_out=plugins=micro:go \
		proto/search/search.proto

gen-encode:
	protoc -I. --go_out=plugins=micro:go \
		proto/encode/encode.proto