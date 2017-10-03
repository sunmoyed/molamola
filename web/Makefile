PATH := node_modules/.bin:$(PATH)

.PHONY : default clean

default: dev

dev:
	bash -c "${MAKE} build-web && \
		( \
			${MAKE} watch-js & \
			${MAKE} watch-css \
		)"

# TODO minify for web builds
build-web: css index js

css: dist/assets/main.css
index: dist/index.html
js: dist/assets/bundle.js

dist/assets/main.css : src/css/*
	node-sass src/css/main.sass -o dist/assets

dist/index.html: src/index.html
	cp src/index.html dist/index.html

dist/assets/bundle.js : src/js/*
	browserify -t -d src/js/index.js -o dist/assets/bundle.js

clean:
	rm -rf dist
	mkdir dist

watch-js:
	watchify src/js/index.js -o dist/assets/bundle.js -v

watch-css:
	node-sass -w src/css/main.sass -o dist/assets
