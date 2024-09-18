grpc_code:
	protoc \
	--go_out=sample_service \
	--go_opt=paths=source_relative \
	--go-grpc_out=sample_service \
	--go-grpc_opt=paths=source_relative \
	sample_service.proto