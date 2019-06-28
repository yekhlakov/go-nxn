package gonxn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNxN(t *testing.T) {
	var x NxN

	x.Link("qwer", "ty", time.Now())
	x.Link("ururu", "azaza", time.Now())
	x.Link("ururu", "ololo", time.Now())
	x.Link("kek", "ololo", time.Now())
	x.Link("trololo", "azaza", time.Now())

	assert.True(t, x.IsLinked("ururu", "azaza"))
	assert.True(t, !x.IsLinked("ururu", "epepe"))
	assert.Equal(t, 2, len(x.ForA("ururu")))
	assert.Equal(t, 0, len(x.ForA("ololo")))
	assert.Equal(t, 2, len(x.ForB("ololo")))
	x.Unlink("kek", "ololo")
	assert.Equal(t, 1, len(x.ForB("ololo")))
	assert.Equal(t, 2, len(x.ForB("azaza")))
	x.RemoveB("azaza")
	assert.Equal(t, 0, len(x.ForB("azaza")))

	assertPanic(t, func() {
		// wrong key a type
		x.Link(t, "lol", time.Now())
	})

	assertPanic(t, func() {
		// wrong key b type
		x.Link("aza", time.Now(), time.Now())
	})

	assertPanic(t, func() {
		// wrong value type
		x.Link("qwer", "lol", "pepepe")
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	f()
}
