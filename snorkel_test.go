package snorkel_test

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/maragudk/is"

	"github.com/maragudk/snorkel"
)

func TestLogger_Log(t *testing.T) {
	t.Run("logs a message without level", func(t *testing.T) {
		l, b := newL()
		l.Log("Yo", 1, "foo", "bar")
		equal(t, `{"name":"Yo","foo":"bar"}`, b)
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
				for range 10000 {
					l.Log("Yo", test.Rate)
				}

				messageLength := len(`{"name":"Yo"}` + "\n")
				messageCount := len(b.String()) / messageLength
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

func equal(t *testing.T, expected string, actual fmt.Stringer) {
	t.Helper()

	if expected != "" {
		expected += "\n"
	}

	is.Equal(t, expected, actual.String())
}
