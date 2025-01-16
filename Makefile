.PHONY: products_proto proto_clean products_run

products_proto:
	protoc --experimental_allow_proto3_optional \
		-I protos \
		--go_out=pkg --go_opt=paths=source_relative \
		--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pkg --grpc-gateway_opt=generate_unbound_methods=true --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=swagger \
		protos/crystal-services/products/v1/products.proto
proto_clean:
	rm -rf pkg/crystal-services
products_run:
	go run cmd/crystal-services/products/main.go -migration=false