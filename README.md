# pdf2jpg

Convert PDF to JPG using the go-fitz Go wrapper for MuPDF fitz library.

## Install

```sh
go get -u github.com/sipkg/pdf2jpg
```

## Usage

* `-in filename` path to pdf file to convert
* `-out file%0d.jpg` pattern file for jpg files
* `-qual 10`  quality of jpg, 1 to 100, the higher the better, default to 10

## Example

```sh
pdf2jpg -out test%03d.jpg -qual 1 -in test.pdf
```
