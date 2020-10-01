# go-image-resizer [![Go Report Card](https://goreportcard.com/badge/github.com/SKF/go-image-resizer)](https://goreportcard.com/report/github.com/SKF/go-image-resizer) [![Build Status on master](https://travis-ci.org/SKF/go-image-resizer.svg?branch=master)](https://travis-ci.org/SKF/go-image-resizer)
go package for resizing images

This package is able to resize images of type gif, jpeg, and png.

## Example
```
package main

import "github.com/SKF/go-image-resizer/resizer"

func main() {
    image, _ := ioutil.ReadFile("image.png")

    newImage, err := ResizeImage(image, resizer.ImageEncodingPNG, 800, 800)

}

```