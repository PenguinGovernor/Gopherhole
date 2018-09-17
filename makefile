# Tools
all: client-encode server-compress client-decode
clean:
	rm -f GopherHoleProtocol/gopherhole.pb.go client-encode server-compress client-decode
# Nicknames
proto: GopherHoleProtocol/gopherhole.pb.go

# Protocol
GopherHoleProtocol/gopherhole.pb.go: GopherHoleProtocol/gopherhole.proto
	cd GopherHoleProtocol && \
	protoc gopherhole.proto --go_out .

client-encode:	compression/client/main.go proto
	cd compression/client && \
	go build -o ../../client-encode -i main.go 

server-compress: compression/server/main.go proto
	cd compression/server && \
	go build -o ../../server-compress -i main.go

client-decode: decompression/client/main.go proto
	cd decompression/client && \
	go build -o ../../client-decode -i main.go