# Protobufs

Create dnd.proto

	protoc --go_out=. dnd.proto
	
Generates Go code dnd.pb.go

gRPC Server
	
	protoc -I . list.proto --go_out=plugins=grpc:.
