9gc: *.go clean ./parser/*.go ./vm/*.go
	go build -o 9gc .

test: 9gc
	go test -v .
	./test.sh

clean:
	-rm -f tmp* 2>/dev/null
