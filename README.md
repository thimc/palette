# palette

Simple color extractor CLI tool that either reads from standard input, or from a
file passed as an argument and dumps the colors in hexadecimal values.

The colors are sorted by their occurrence and printed out to the console with a
`#RRGGBB` format.

## Example usage

	./palette /path/to/my/image.png

	./palette <image.png

	cat my_image.png | palette

	maim -s | palette

## TODOs

* add support for multiple files. There are two ways this could be handled.
First, the files are processed separately and `palette` will dump the `n * amount
of files` most dominant colors.  The second alternative to this is to process
the files in a similar manner but merging their results and ultimately only
printing `n` amount of colors.
* find a suck less way to exclude colors that look to similar, using an
algorithm such as Euclidean distance could perhaps be of use here.

## License

MIT
