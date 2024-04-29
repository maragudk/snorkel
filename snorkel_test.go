package snorkel_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maragudk/is"

	"github.com/maragudk/snorkel"
)

func TestLogger_Log(t *testing.T) {
	t.Run("logs a message without level", func(t *testing.T) {
		l, b := newL()
		l.Log("Yo", "foo", "bar")
		equal(t, `{"msg":"Yo","foo":"bar"}`, b)
	})
}

func newL() (*snorkel.Logger, *strings.Builder) {
	var b strings.Builder
	return snorkel.New(snorkel.Options{NoTime: true, W: &b}), &b
}

func equal(t *testing.T, expected string, actual fmt.Stringer) {
	t.Helper()

	is.Equal(t, expected+"\n", actual.String())
}
