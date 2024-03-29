# imgdiff

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![test-and-build](https://github.com/kemokemo/imgdiff/actions/workflows/test-and-build.yml/badge.svg)](https://github.com/kemokemo/imgdiff/actions/workflows/test-and-build.yml)

This tool compares old and new versions of image files and generates color-coded image in a diff format.

![sample diff image](images/sample-diff-image.png)

This is based on the [diff-image](https://github.com/murooka/go-diff-image/blob/master/cmd/diff-image/main.go).
Thanks to the wonderful [@murooka](https://github.com/murooka)'s [go-diff-image](https://github.com/murooka/go-diff-image).

## Install

### Homebrew

```sh
brew install kemokemo/tap/imgdiff
```

### Scoop

First, add my scoop-bucket.

```sh
scoop bucket add kemokemo-bucket https://github.com/kemokemo/scoop-bucket.git
```

Next, install this app by running the following.

```sh
scoop install imgdiff
```

### Binary

Get the latest version from [the release page](https://github.com/kemokemo/imgdiff/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your `$PATH`.

## Usage

```sh
$ imgdiff -h
Usage: imgdiff [<option>...] <old image> <new image>
  -h	display help
  -o string
    	output filename (default "diff.png")
  -v	display version
```

This tool supports the following image formats.

- png
- jpeg
- gif
- bmp

### Example

```sh
$ imgdiff -o=diff/screen.png v1/screen.png v2/screen.png
```

## License

[MIT](https://github.com/kemokemo/imgdiff/blob/main/LICENSE)

## Author

[kemokemo](https://github.com/kemokemo)
