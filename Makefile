.PHONY: all format analyze test test-ci demo fix upgrade images

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

# Shared dependencies for all images
COMMON_SRCS := common.go demo/demo.go

# Pattern rule for static PNGs
images/%.png: %.go $(COMMON_SRCS)
	go run demo/demo.go -algo $* -output $@
	zopflipng -y $@ $@

# Pattern rule for animation GIFs
images/%_animation.gif: %.go $(COMMON_SRCS)
	go run demo/demo.go -algo $* -output $@
	gifsicle -O3 --colors 256 -o $@ $@

# Special cases for Moore (which depends on Hilbert)
images/moore.png: hilbert.go
images/moore_animation.gif: hilbert.go

# Special case for the logo
images/logo.png: hilbert.go $(COMMON_SRCS)
	go run demo/demo.go -logo -output $@
	zopflipng -y $@ $@

# Aggregate target
IMAGES := images/logo.png \
	images/hilbert.png images/hilbert_animation.gif \
	images/peano.png images/peano_animation.gif \
	images/morton.png images/morton_animation.gif \
	images/moore.png images/moore_animation.gif \
	images/sierpinski.png images/sierpinski_animation.gif

images: $(IMAGES)

demo: images

fix:
	go fmt ./...
	go fix ./...

upgrade:
	go mod tidy
	go get -u ./...
	go mod tidy
