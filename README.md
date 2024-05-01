# snorkel

<img src="logo.png" alt="Logo" width="300" align="right"/>

[![GoDoc](https://pkg.go.dev/badge/github.com/maragudk/snorkel)](https://pkg.go.dev/github.com/maragudk/snorkel)
[![Go](https://github.com/maragudk/snorkel/actions/workflows/ci.yml/badge.svg)](https://github.com/maragudk/snorkel/actions/workflows/ci.yml)

[Scuba](https://research.facebook.com/publications/scuba-diving-into-data-at-facebook/) for the rest of us.

Made in 🇩🇰 by [maragu](https://www.maragu.dk/), maker of [online Go courses](https://www.golang.dk/).

## Usage

```go
package main

import (
	"github.com/maragudk/snorkel"
)

func main() {
	log := snorkel.New(snorkel.Options{})

	// Log an event with the name "Yo", sample rate 1, and env=sparkly
	log.Event("Yo", 1, "env", "sparkly")
}
```
