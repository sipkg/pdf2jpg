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

### How to cross-compile from linux to windows

You must enable [CGO](https://golang.org/cmd/cgo/) as go-fitz uses it.
You also need an appropriate compiler and toolchain for the targeted arch. 
Here, we will use [MinGW-w64](https://www.mingw-w64.org/). 
See their website for instructions to install it for your distribution.

```sh
# install MinGW the Debian/Ubuntu way
sudo apt-get install mingw-w64

# then you can cross-compile
GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" go build
```
