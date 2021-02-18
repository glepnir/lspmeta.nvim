build:
	go build -o ~/workstation/go/bin/lspmeta ./cmd/lspmeta/main.go

manifest:
	./bin/lspmeta -manifest lspmeta

clean:
	rm -rf ./bin/lspmeta
	rm -rf ~/workstation/go/bin/lspmeta
