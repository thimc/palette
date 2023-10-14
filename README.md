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

## License

MIT
