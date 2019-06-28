package gonxn

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNxN(t *testing.T) {
    var x NxN

    x.Link("qwer", "ty")
    x.Link("ururu", "azaza")
    x.Link("ururu", "ololo")
    x.Link("kek", "ololo")

    assert.True(t, x.IsLinked("ururu", "azaza"))
    assert.True(t, !x.IsLinked("ururu", "epepe"))
    assert.Equal(t, 2, len(x.ForA("ururu")))
    assert.Equal(t, 0, len(x.ForA("ololo")))
    assert.Equal(t, 2, len(x.ForB("ololo")))
    x.Unlink("kek", "ololo")
    assert.Equal(t, 1, len(x.ForB("ololo")))
}
