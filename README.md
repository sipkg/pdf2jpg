# pdf2jpg

Convert PDF to JPG using the go-fitz Go wrapper for MuPDF fitz library.

## Install

```sh
go install github.com/sipkg/pdf2jpg
```

## Usage

* `-in filename` path to pdf file to convert
* `-out file%0d.jpg` pattern file for jpg files
* `-qual 10`  quality of jpg, 1 to 100, the higher the better, default to 10

## Example

```sh
pdf2jpg -out test%03d.jpg -qual 1 -in test.pdf
```

## FAQ

### How to across-compile from linux to windows

Because go-fitz library uses https://golang.org/cmd/cgo/, so you need to have compiler and toolchain for arch you want to build for, and also, enable `CGO_ENABLED=1`. 

1. `sudo apt-get install mingw-w64`
1. `GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" go build`