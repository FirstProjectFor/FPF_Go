package options

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func TestNewWithDefault(t *testing.T) {
	background := context.Background()

	person, err := NewWithDefault(background, FallBackDefault)

	assert.Nil(t, err)
	assert.Equal(t, "default", person.Name())
	assert.Equal(t, 20, person.Age())

	person, err = NewWithDefault(background, WithName("xiaotian"), FallBackDefault)
	assert.Nil(t, err)
	assert.Equal(t, "xiaotian", person.Name())
	assert.Equal(t, 20, person.Age())

	person, err = NewWithDefault(background, WithAge(27), FallBackDefault)
	assert.Nil(t, err)
	assert.Equal(t, "default", person.Name())
	assert.Equal(t, 27, person.Age())

	person, err = NewWithDefault(background, WithAge(27), WithName("xiaotian"), FallBackDefault)
	assert.Nil(t, err)
	assert.Equal(t, "xiaotian", person.Name())
	assert.Equal(t, 27, person.Age())
}
