all: gen-search gen-encode

gen-search:
	protoc --proto_path=. --go_out=go --micro_out=go proto/search/search.proto

gen-encode:
	protoc --proto_path=. --go_out=go --micro_out=go proto/encode/encode.proto