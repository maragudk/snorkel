# snorkel

<img src="logo.png" alt="Logo" width="300" align="right"/>

[![GoDoc](https://pkg.go.dev/badge/maragu.dev/snorkel)](https://pkg.go.dev/maragu.dev/snorkel)
[![Go](https://github.com/maragudk/snorkel/actions/workflows/ci.yml/badge.svg)](https://github.com/maragudk/snorkel/actions/workflows/ci.yml)

[Scuba](https://research.facebook.com/publications/scuba-diving-into-data-at-facebook/) for the rest of us.

```shell
go get maragu.dev/snorkel
```

Made in 🇩🇰 by [maragu](https://www.maragu.dk/), maker of [online Go courses](https://www.golang.dk/).

## Features

- Simple logger with just a single method, `Event`, which logs named events with a given sample rate.
- Sane defaults log to STD_ERR and include event name, sample rate, time, and build info. More to come.
- No external dependencies.

## Usage

```go
package main

import (
	"maragu.dev/snorkel"
)

func main() {
	log := snorkel.New(snorkel.Options{})

	// Log an event with the name "Yo", sample rate 1, and env=sparkly
	log.Event("Yo", 1, "env", "sparkly")
}
```
