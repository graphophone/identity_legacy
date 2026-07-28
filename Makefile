gen:
	@protoc \
	--proto_path=proto "proto/user.proto" \
	--go_out=internal/services/common/user \
	--go_opt=paths=source_relative \
	--go-grpc_out=internal/services/common/user \
	--go-grpc_opt=paths=source_relative