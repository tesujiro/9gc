9gc: *.go clean
	go build -o 9gc .

test: 9gc
	./test.sh

clean:
	-rm tmp* 2>/dev/null
