# Stego CLI

![Github Action Build](https://github.com/gzcharleszhang/stego/workflows/Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/gzcharleszhang/stego)](https://goreportcard.com/report/github.com/gzcharleszhang/stego)

Stego is a command-line interface for encoding and decoding secret data in images,
using the Least Significant Bit (LSB) steganography technique. Currently, Stego only supports PNG images,
but is able to hide any type of data and files.

Stego is also available as a Go [package](#Using-Stego-as-a-package).

## Getting Started
```shell
go get -u github.com/gzcharleszhang/stego
```

## Usage

### Encoding a message in an image
```shell
stego encode ./stego/example.png --data "Stego is a steganography CLI tool."
```
By default, Stego will add a `-out` suffix to the output image. For example, the above encoded image
can be found at `./stego/example-out.png`

To specify an output path, use the `--out` or `-o` flag
```shell
stego encode ./stego/example.png -d "Stego is a steganography CLI tool." -o ./out/example.png
```

### Decoding a message from an image
```shell
stego decode ./example.png
```

## Using Stego as a package
### Installation
```shell
go get -u github.com/gzcharleszhang/stego/pkg/stegolsb
```

### Encoding
Stego can encode a message into an image from the [Go image package](https://golang.org/pkg/image/)

```go
outImg, err := stego_lsb.LSBEncode(img, "Hello, world!")
if err != nil {
    fmt.Printf("Error encoding message: %v\n", err)
}
```

#### Encoding with more than 1 bit
The package also supports encoding with multiple bits per byte.
This allows the image to encode more data, however it will decrease
the encoded image quality compared to the original image.

Stego will first use the least significant bit, then the second
least significant bit, and so on.

Encoding an image using up to 2 bit per byte
```go
outImg, err := stego_lsb.Encode(img, "Hello, world!", 2)
if err != nil {
    fmt.Printf("Error encoding message: %v\n", err)
}
```

### Decoding
Stego can attempt to decode an image from the [Go image package](https://golang.org/pkg/image/)

```go
message, err := stego_lsb.Decode(outImg)
if err != nil {
    fmt.Printf("Error decoding image: %v\n", err)
}
```

### Maximum message size for encoding
Stego can calculate the maximum number of bytes available for
encoding in an image.

```go
maxSize, err := stego_lsb.MaxLSBEncodeSize(img)
if err != nil {
    fmt.Printf("Error getting max encode size: %v", err)
}
```

Use `stego_lsb.MaxEncodeSize` to get the max size if using
more than 1 bit per byte for encoding.

The maximum number of bytes available in the image
when using up to 2 bits per byte for encoding.
```go
maxSize, err := stego_lsb.MaxEncodeSize(img, 2)
if err != nil {
    fmt.Printf("Error getting max encode size: %v", err)
}
```
