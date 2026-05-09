all: format analyze test

format:
	go fmt ./...
	goimports -w .

analyze:
	go vet ./...
	staticcheck ./...

test:
	go test ./...

test-ci:
	go test -v ./...

images:
	go run demo/demo.go
	# Optimise PNGs
	zopflipng -y logo.png images/logo.png
	zopflipng -y hilbert.png images/hilbert.png
	zopflipng -y peano.png images/peano.png
	zopflipng -y morton.png images/morton.png
	zopflipng -y moore.png images/moore.png
	# Optimise GIFs
	gifsicle -O3 --colors 256 -o images/hilbert_animation.gif hilbert_animation.gif
	gifsicle -O3 --colors 256 -o images/peano_animation.gif peano_animation.gif
	gifsicle -O3 --colors 256 -o images/morton_animation.gif morton_animation.gif
	gifsicle -O3 --colors 256 -o images/moore_animation.gif moore_animation.gif
	# Cleanup
	rm logo.png hilbert.png peano.png morton.png moore.png
	rm hilbert_animation.gif peano_animation.gif morton_animation.gif moore_animation.gif

demo: images

fix:
	go fmt ./...
	go fix ./...

upgrade:
	go mod tidy
	go get -u ./...
	go mod tidy
