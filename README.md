# Stego CLI

![Github Action Build](https://github.com/gzcharleszhang/stego/workflows/Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/gzcharleszhang/stego)](https://goreportcard.com/report/github.com/gzcharleszhang/stego)

Stego is a command-line interface for encoding and decoding secret data in images,
using the Least Significant Bit (LSB) steganography technique. Currently, Stego only supports PNG images,
but is able to hide any type of data and files.

Stego achieves this by first prepending the size of the secret data to the data itself.
It then writes each bit of the data
to the least significant bit of each pixel's RGB values
in the image. When Stego decodes an image, it collects each
least significant bit of the image,
then recompose the bits back into data. The first 8 LSB will tell
Stego the total size of the hidden data.

Stego is also available as a Go [package](#Using-Stego-as-a-package).

## Getting Started
```shell
$ go get -u github.com/gzcharleszhang/stego
```

## Usage

### Encoding a message in an image
```shell
$ stego encode ./stego/example.png --data "Stego is a steganography CLI tool."
```
By default, Stego will add a `-out` suffix to the output image. For example, the above encoded image
can be found at `./stego/example-out.png`

To specify an output path, use the `--out` or `-o` flag
```shell
$ stego encode ./stego/example.png -d "Stego is a steganography CLI tool." -o ./out/example.png
```

### Decoding a message from an image
```shell
$ stego decode ./example.png
```

### Checking the maximum encoding size of an image
Before encoding data in an image, you may want to know
how much hidden data an image can hold.

Stego has a `size` command that calculates the
maximum encoding size of an image in bytes.
```shell
$ stego size ./example.png
1179644
```

Pretty print using the `--pretty` or `-p` flag.
```shell
$ stego size --pretty ./example.png
1.12 MB
```

## Using Stego as a package
### Package-only installation
```shell
$ go get -u github.com/gzcharleszhang/stego/pkg/stegolsb
```

### Encoding
Stego can encode a message into an image from the [Go image package](https://golang.org/pkg/image/)

```go
import (
    "fmt"
    "github.com/gzcharleszhang/stego/pkg/stegolsb"
    "image"
)

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
Stego attempts to decode an [image](https://golang.org/pkg/image/)
and prints the hidden data.
```go
message, err := stego_lsb.Decode(img)
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
