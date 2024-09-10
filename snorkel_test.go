package snorkel_test

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"testing"

	"maragu.dev/is"

	"maragu.dev/snorkel"
)

func ExampleLogger() {
	log := snorkel.New(snorkel.Options{})
	log.Event("Yo", 1, "env", "sparkly")
}

func TestLogger_Event(t *testing.T) {
	t.Run("logs a message without level", func(t *testing.T) {
		l, b := newL()
		l.Event("Yo", 1, "foo", "bar")
		is.True(t, strings.Contains(b.String(), `"name":"Yo","rate":1,"foo":"bar"`))
	})

	t.Run("logs message depending on sample rate", func(t *testing.T) {
		tests := []struct {
			Rate  float32
			Count int
		}{
			{0, 0},
			{0.25, 2530},
			{0.5, 5028},
			{0.75, 7540},
			{1, 10000},
		}

		for _, test := range tests {
			t.Run(fmt.Sprint(test.Rate), func(t *testing.T) {
				l, b := newL()
				var messageCount, messageLength int

				// If the rate is non-zero, try until there's a message in the log
				for test.Rate > 0 && messageLength == 0 {
					l.Event("Yo", test.Rate)
					messageLength = len(b.String())
				}

				l, b = newL()
				for range 10000 {
					l.Event("Yo", test.Rate)
				}

				if messageLength > 0 {
					messageCount = len(b.String()) / messageLength
				}
				is.Equal(t, test.Count, messageCount)
			})
		}
	})
}

func newL() (*snorkel.Logger, *strings.Builder) {
	var b strings.Builder
	return snorkel.New(snorkel.Options{
		NoTime:       true,
		RandomSource: rand.NewPCG(1, 2),
		W:            &b,
	}), &b
}
