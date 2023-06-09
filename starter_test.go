package starter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("Meee")
	assert.Equal(t, "Hello Meee. Welcome!", greeting)

	another_greeting := SayHello("Fern")
	assert.Equal(t, "Hello Fern. Welcome!", another_greeting)
}
