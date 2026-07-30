gen_user:
	@protoc \
	--proto_path=proto "proto/user.proto" \
	--go_out=internal/services/common/user \
	--go_opt=paths=source_relative \
	--go-grpc_out=internal/services/common/user \
	--go-grpc_opt=paths=source_relative

gen_auth:
	@protoc \
	--proto_path=proto "proto/auth.proto" \
	--go_out=internal/services/common/auth \
	--go_opt=paths=source_relative \
	--go-grpc_out=internal/services/common/auth \
	--go-grpc_opt=paths=source_relative

gen: gen_user gen_auth