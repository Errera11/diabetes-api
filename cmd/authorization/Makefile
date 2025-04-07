gen:
	@protoc \
			--proto_path=../../protobuf \
			--go_opt=paths=source_relative --go_out=internal/protogen/user \
			../../protobuf/user.proto

	@protoc \
    		--proto_path=../../protobuf \
    		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/protogen/authorization \
    		--go_opt=paths=source_relative --go_out=internal/protogen/authorization \
    		../../protobuf/authorization.proto



#@protoc  \
#  --proto_path=${GOPATH}/src \
#  --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
#  --proto_path=. \
#  --go_out=. \
#  --govalidators_out=. \
#  *.proto