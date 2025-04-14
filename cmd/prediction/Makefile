gen:
	@protoc \
		--proto_path=../../protobuf \
		--go_opt=paths=source_relative --go_out=internal/protogen/ \
		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/protogen/ \
		../../protobuf/prediction.proto

	@protoc \
		--proto_path=../../protobuf \
		--go_opt=paths=source_relative --go_out=internal/protogen \
		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/protogen \
		../../protobuf/authorization.proto

migrate-db:
	@migrate \
 			-database ${POSTGRESQL_URL} -path internal/prediction/infrastructure/db/migrations up

#	@protoc \
#			--proto_path=../../protobuf \
#			--go-grpc_opt=paths=source_relative --go-grpc_out=internal/protogen/user \
#			--go_opt=paths=source_relative --go_out=internal/protogen/user \
#			../../protobuf/user.proto

#	@protoc \
#    		--proto_path=../../protobuf \
#    		--go-grpc_opt=paths=source_relative --go-grpc_out=internal/protogen/authorization \
#    		--go_opt=paths=source_relative --go_out=internal/protogen/authorization \
#    		../../protobuf/authorization.proto
