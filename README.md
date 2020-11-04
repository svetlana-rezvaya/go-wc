# go-wc

[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-wc)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-wc)

The clone of the [Unix wc tool](<https://en.wikipedia.org/wiki/Wc_(Unix)>).

## Installation

```
$ go get github.com/svetlana-rezvaya/go-wc
```

## Usage

```
$ go-wc -h | -help | --help
$ go-wc
```

Stdin: text in the [UTF-8 encoding](https://en.wikipedia.org/wiki/UTF-8).

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit.

## Output Example

```
amount of lines = 10
amount of words = 10
amount of bytes = 21
amount of characters = 21
```

## License

The MIT License (MIT)

Copyright &copy; 2020 svetlana-rezvaya
