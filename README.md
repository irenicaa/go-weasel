# go-weasel

[![GoDoc](https://godoc.org/github.com/irenicaa/go-weasel?status.svg)](https://godoc.org/github.com/irenicaa/go-weasel)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-weasel)](https://goreportcard.com/report/github.com/irenicaa/go-weasel)
[![Build Status](https://travis-ci.org/irenicaa/go-weasel.svg?branch=master)](https://travis-ci.org/irenicaa/go-weasel)
[![codecov](https://codecov.io/gh/irenicaa/go-weasel/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-weasel)

[Weasel program](https://en.wikipedia.org/wiki/Weasel_program).

## Installation

```
$ go get github.com/irenicaa/go-weasel/...
```

## Usage

```
$ go-weasel -h | -help | --help
$ go-weasel [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-sample STRING` &mdash; target string (default: `METHINKS IT IS LIKE A WEASEL`);
- `-rate FLOAT` &mdash; mutation rate (default: `0.05`);
- `-population INTEGER` &mdash; population size (default: `100`).

## Output Example

```
0 WIAKYMUAFZISJOGLNHBBPJFUNYLL
10 WETHYVKSFITSO GLNLT PLLANSEL
20 WETHIVKS IT OSGLMKT ALDAVSEL
30 WETHINKS IT OS LMKE APWAVSEL
40 WETHINKS IT IS LIKE APWWVSEL
50 WETHINKS IT IS LIKE A WWVSEL
60 WETHINKS IT IS LIKE A WWVSEL
70 WETHINKS IT IS LIKE A WEASEL
80 YETHINKS IT IS LIKE A WEASEL
90 PETHINKS IT IS LIKE A WEASEL
100 SETHINKS IT IS LIKE A WEASEL
110 SETHINKS IT IS LIKE A WEASEL
114 METHINKS IT IS LIKE A WEASEL
```

## License

The MIT License (MIT)

Copyright &copy; 2020 irenica
