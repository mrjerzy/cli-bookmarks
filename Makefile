build:
	go build

install: build
	go install

clean:
	rm bookmarks
