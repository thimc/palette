package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"sort"
)

var (
	amountFlag = flag.Int("d", 15, "number of colors to be printed")
)

func usage(err *error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", *err)
	}
	flag.Usage()
	os.Exit(1)
}

func main() {
	var err error
	file := os.Stdin
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		file, err = os.Open(args[0])
		if err != nil {
			usage(&err)
		}
		defer file.Close()

	} else {
		fi, err := file.Stat()
		if err != nil {
			usage(&err)
		}

		if fi.Mode()&os.ModeNamedPipe == 0 {
			usage(nil)
		}
	}

	colors, err := processImage(file)
	if err != nil {
		usage(&err)
	}

	for i := 0; i < *amountFlag; i++ {
		if i < len(colors) {
			fmt.Printf("%v\n", colors[i])
		}
	}
}

func processImage(r io.Reader) ([]string, error) {
	out := []string{}

	img, _, err := image.Decode(r)
	if err != nil {
		return out, fmt.Errorf("cannot decode file: %s", err)
	}

	bounds := img.Bounds().Size()
	colors := make(map[color.Color]int, 0)

	for y := 0; y < bounds.X; y++ {
		for x := 0; x < bounds.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			col := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
			}
			colors[col]++
		}
	}

	for col, count := range colors {
		if count > 1 {
			r, g, b, _ := col.RGBA()
			hex := fmt.Sprintf("#%02x%02x%02x", uint8(r), uint8(g), uint8(b))
			out = append(out, hex)
		}
	}
	sort.Strings(out)

	return out, nil
}
