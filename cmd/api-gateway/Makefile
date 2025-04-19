gen:
	@protoc -I ../../protobuf \
	  --go_out ./internal/protogen \
	  --go_opt paths=source_relative \
	  --go-grpc_out ./internal/protogen \
	  --go-grpc_opt paths=source_relative \
	  ../../protobuf/user.proto
	@protoc -I ../../protobuf \
      --go_out ./internal/protogen \
      --go_opt paths=source_relative \
      --go-grpc_out ./internal/protogen \
      --go-grpc_opt paths=source_relative \
      ../../protobuf/authorization.proto
	@protoc -I ../../protobuf \
      --go_out ./internal/protogen \
      --go_opt paths=source_relative \
      --go-grpc_out ./internal/protogen \
      --go-grpc_opt paths=source_relative \
      ../../protobuf/prediction.proto

gen-gateway:
	@protoc -I ../../protobuf \
               --grpc-gateway_out ./internal/protogen \
               --grpc-gateway_opt paths=source_relative \
               --grpc-gateway_opt grpc_api_configuration=./gateway-config.yaml \
               --experimental_allow_proto3_optional \
               ../../protobuf/*.proto
#	@protoc -I ../../protobuf \
#       --grpc-gateway_out ./internal/protogen/user \
#       --grpc-gateway_opt paths=source_relative \
#       --grpc-gateway_opt grpc_api_configuration=./gateway-config.yaml \
#       ../../protobuf/user.proto


#gen-proxy:
#	@protoc -I . \
#	  --grpc-gateway_out ./gen/go \
#	  --grpc-gateway_opt paths=source_relative \
#	  --grpc-gateway_opt grpc_api_configuration=/your_service.yaml \
#	  your/service/v1/your_service.proto