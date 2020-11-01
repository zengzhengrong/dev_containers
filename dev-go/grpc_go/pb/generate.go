package sum

//go:generate protoc -I/usr/local/include -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --go_out=plugins=grpc:. --goclay_out=allow_delete_body=true,grpc_api_configuration=$GOPACKAGE.yaml:. ./$GOPACKAGE.proto
//go:generate bash remove_omitempty.sh
