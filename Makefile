mock-expected-keepers:
	mockgen -source=x/dataocean/types/expected_keepers.go \
		-package testutil \
		-destination=x/dataocean/testutil/expected_keepers_mocks.go 

install-protoc-gen-ts:
	mkdir -p scripts
	cd scripts && npm install ts-proto --save-dev --save-exact
	mkdir -p scripts/protoc
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.11/protoc-21.11-osx-aarch_64.zip -o scripts/protoc/protoc.zip
	cd scripts/protoc && unzip -o protoc.zip
	rm scripts/protoc/protoc.zip

cosmos-version = v0.45.4

download-cosmos-proto:
	mkdir -p proto/cosmos/base/query/v1beta1
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/proto/cosmos/base/query/v1beta1/pagination.proto -o proto/cosmos/base/query/v1beta1/pagination.proto
	mkdir -p proto/google/api
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/annotations.proto -o proto/google/api/annotations.proto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/http.proto -o proto/google/api/http.proto
	mkdir -p proto/gogoproto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/gogoproto/gogo.proto -o proto/gogoproto/gogo.proto

gen-protoc-ts: 
	mkdir -p ./client/src/types/generated/
	ls proto/dataocean/dataocean | xargs -I {} ./scripts/protoc/bin/protoc \
		--plugin="./scripts/node_modules/.bin/protoc-gen-ts_proto" \
		--ts_proto_out="./client/src/types/generated" \
		--proto_path="./proto" \
		--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
		dataocean/dataocean/{}
	
install-and-gen-protoc-ts: download-cosmos-proto install-protoc-gen-ts gen-protoc-ts