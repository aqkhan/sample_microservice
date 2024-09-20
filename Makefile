grpc_code:
	protoc \
	--go_out=sample_service \
	--go_opt=paths=source_relative \
	--go-grpc_out=require_unimplemented_servers=false:sample_service \
	--go-grpc_opt=paths=source_relative \
	sample_service.proto

streaming_code:
	protoc \
	--go_out=streaming_service \
	--go_opt=paths=source_relative \
	--go-grpc_out=require_unimplemented_servers=false:streaming_service \
	--go-grpc_opt=paths=source_relative \
	streaming_service.proto

code_update:
	git pull origin main
	go build -o sample_service_grpc main.go
	lsof -ti:6969 | xargs -r kill -9
	nohup ./sample_service_grpc &