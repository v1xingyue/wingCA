all:linux mac

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux

linux:.pre-build
	GOOS=linux go build -ldflags "-w -s" -o build/linux/wingCA

mac:.pre-build
	GOOS=darwin go build -ldflags "-w -s" -o build/darwin/wingCA

clean:
	rm build/linux/wingCA
	rm build/darwin/wingCA
