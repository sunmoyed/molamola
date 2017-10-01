PATH  := node_modules/.bin:$(PATH)
SHELL := /bin/bash

.PHONY : default clean

default: dev

dev: build-web serve & watch-js & watch-css

# TODO minify for web builds
build-web: css index js

css: dist/main.css
index: dist/index.html
js: dist/js/

dist/main.css : src/css/*
	node-sass src/css/main.sass -o dist

dist/index.html: src/index.html
	cp src/index.html dist/index.html

dist/js/ : src/js/*
	browserify -t -d src/js/index.js -o dist/bundle.js

clean:
	rm -rf dist
	mkdir dist

watch-js:
	watchify src/js/index.js -o dist/bundle.js -v

watch-css:
	node-sass -w src/css/main.sass -o dist

serve:
	pushd dist; python -m SimpleHTTPServer; pop
