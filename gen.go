//go:build generate
// +build generate

package main

//go:generate npx tailwindcss -i ./src/input.css -o ./static/output.css
//go:generate npx esbuild ./src/input.js --bundle --outfile=./static/output.js
