package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"

	diffimage "github.com/murooka/go-diff-image"
)

const (
	exitOk = iota
	exitInvalidArg
	exitInvalidImage
	exitFailedOperation
)

var (
	out  string
	help bool
)

func init() {
	flag.StringVar(&out, "o", "diff.png", "output filename")
	flag.BoolVar(&help, "h", false, "display help")
	flag.Parse()
}

func main() {
	os.Exit(run())
}

func run() int {
	if help {
		fmt.Fprintf(os.Stdout, "Usage: imgdiff [<option>...] <old image> <new image>\n")
		flag.PrintDefaults()
		return exitOk
	}

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Please set args.\n Usage: imgdiff [<option>...] <old image> <new image>\n")
		flag.PrintDefaults()
		return exitInvalidArg
	}

	fOld, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "first arg is empty. please set old image path: %v\n", err)
		return exitInvalidArg
	}
	defer fOld.Close()

	fNew, err := os.Open(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "second arg is empty. please set new image path: %v\n", err)
		return exitInvalidArg
	}
	defer fNew.Close()

	// TODO: Decoding process based on extensions
	imgOld, formatOld, err := image.Decode(fOld)
	if err != nil {
		fmt.Fprintf(os.Stderr, "first arg's image is invalid: %v\n", err)
		return exitInvalidImage
	}
	imgNew, formatNew, err := image.Decode(fNew)
	if err != nil {
		fmt.Fprintf(os.Stderr, "second arg's image is invalid: %v\n", err)
		return exitInvalidImage
	}
	if formatOld != formatNew {
		fmt.Fprintf(os.Stderr, "image format does not match (first:%s, second:%s)\n", formatOld, formatNew)
		return exitInvalidImage
	}

	dst := diffimage.DiffImage(imgOld, imgNew)

	dir := filepath.Dir(out)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create directory to save diff file: %v\n", err)
			return exitInvalidArg
		}
	}
	fDiff, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "output image's path is invalid:%v\n", err)
		return exitInvalidArg
	}
	defer fDiff.Close()

	var e error
	switch formatNew {
	case "png":
		e = png.Encode(fDiff, dst)
	case "jpeg":
		e = jpeg.Encode(fDiff, dst, nil)
	case "gif":
		e = gif.Encode(fDiff, dst, nil)
	case "bmp":
		e = bmp.Encode(fDiff, dst)
	default:
		fmt.Fprintf(os.Stderr, "you passed unsupprted image format: %v\n", formatNew)
		e = png.Encode(fDiff, dst)
	}
	if e != nil {
		fmt.Fprintf(os.Stderr, "failed to save the output image: %v\n", err)
		return exitFailedOperation
	}

	return exitOk
}
