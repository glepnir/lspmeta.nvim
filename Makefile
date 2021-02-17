all:
	go build -o lspmeta .

manifest:
	./lspmeta -manifest lspmeta
