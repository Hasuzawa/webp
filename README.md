## Introduction

A Go module for decoding the WebP image format, heavily inspired by existing implementations and the [experimental package for image/webp](https://pkg.go.dev/golang.org/x/image/webp).

## Chunk representation of WebP

(to be added)

## Inspecting Webp file

[The official CLIs](https://developers.google.com/speed/webp/download)

`file {filename}` // e.g. RIFF (little-endian) data, Web/P image

`hexdump {filename}`

`wc -c {filename}`

For advanced viewing and editing, consider using a graphics software or hex editor.

## Documentation & Specification

[WebP image format | Google](https://developers.google.com/speed/webp)

[Source code for libwebp | Google](https://chromium.googlesource.com/webm/libwebp/)

[Resource Interchange File Format (RIFF)](https://www.loc.gov/preservation/digital/formats/fdd/fdd000025.shtml)

[VP8 Data Format and Decoding Guide | RFC 6386](https://datatracker.ietf.org/doc/html/rfc6386)

## Reference

[Experimental module Go webp | Google](https://pkg.go.dev/golang.org/x/image/webp)

[Go image/png](https://pkg.go.dev/image/png)

[Go image/jpeg](https://pkg.go.dev/image/jpeg)

[Go image/gif](https://pkg.go.dev/image/gif)

[go-webp | kolesa-team](https://github.com/kolesa-team/go-webp)
