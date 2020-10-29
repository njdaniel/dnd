.PHONEY: build clean

build:
	packr2
	go build

clean:
	go clean
	packr2 clean