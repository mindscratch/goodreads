build:
	(cd cmd/markdown && GO111MODULE=on go build -o markdown .)

code:
	GO111MODULE=on code .
