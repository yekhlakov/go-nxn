package gonxn

import (
	"errors"
	"reflect"
)

type NxN struct {
	forward  map[interface{}]map[interface{}]interface{}
	backward map[interface{}]map[interface{}]interface{}
	lastLink struct {
		A interface{}
		B interface{}
		V interface{}
	}
}

func (m *NxN) ensureA(a interface{}) {
	if m.forward == nil {
		m.forward = make(map[interface{}]map[interface{}]interface{})
	}

	if m.forward[a] == nil {
		m.forward[a] = make(map[interface{}]interface{})
	}
}

func (m *NxN) ensureB(b interface{}) {
	if m.backward == nil {
		m.backward = make(map[interface{}]map[interface{}]interface{})
	}

	if m.backward[b] == nil {
		m.backward[b] = make(map[interface{}]interface{})
	}
}

func (m *NxN) typeCheckA(a interface{}) {
	if m.lastLink.A == nil {
		m.lastLink.A = a
	}

	if reflect.TypeOf(a) != reflect.TypeOf(m.lastLink.A) {
		panic(errors.New("key a type mismatch"))
	}
}

func (m *NxN) typeCheckB(b interface{}) {
	if m.lastLink.B == nil {
		m.lastLink.B = b
	}

	if reflect.TypeOf(b) != reflect.TypeOf(m.lastLink.B) {
		panic(errors.New("key b type mismatch"))
	}
}

func (m *NxN) typeCheckV(v interface{}) {
	if m.lastLink.V == nil {
		m.lastLink.V = v
	}

	if reflect.TypeOf(v) != reflect.TypeOf(m.lastLink.V) {
		panic(errors.New("value type mismatch"))
	}
}

func (m *NxN) Link(a interface{}, b interface{}, val interface{}) {
	m.typeCheckA(a)
	m.typeCheckB(b)
	m.typeCheckV(val)
	m.ensureA(a)
	m.ensureB(b)
	m.forward[a][b] = val
	m.backward[b][a] = "1" // use string
}

func (m *NxN) isEmptyAB(a interface{}, b interface{}) bool {
	return m.forward == nil ||
		m.backward == nil ||
		m.forward[a] == nil ||
		m.backward[b] == nil
}

func (m *NxN) isEmptyA(a interface{}) bool {
	return m.forward == nil ||
		m.backward == nil ||
		m.forward[a] == nil
}

func (m *NxN) isEmptyB(b interface{}) bool {
	return m.forward == nil ||
		m.backward == nil ||
		m.backward[b] == nil
}

func (m *NxN) Unlink(a interface{}, b interface{}) {
	if m.isEmptyAB(a, b) {
		return
	}
	delete(m.forward[a], b)
	delete(m.backward[b], a)
}

func (m *NxN) RemoveA(a interface{}) {
	if m.isEmptyA(a) {
		return
	}

	for b := range m.forward[a] {
		delete(m.backward[b], a)
	}

	delete(m.forward, a)
}

func (m *NxN) RemoveB(b interface{}) {
	if m.isEmptyB(b) {
		return
	}

	for a := range m.backward[b] {
		delete(m.forward[a], b)
	}

	delete(m.backward, b)
}

func (m *NxN) ForAB(a interface{}, b interface{}) interface{} {
	if m.isEmptyAB(a, b) {
		return nil
	}

	return m.forward[a][b]
}

func (m *NxN) IsLinked(a interface{}, b interface{}) bool {
	if m.isEmptyAB(a, b) {
		return false
	}

	return m.forward[a][b] != nil
}

func getKeys(m map[interface{}]interface{}) []interface{} {
	keys := make([]interface{}, len(m))
	c := 0
	for k := range m {
		keys[c] = k
		c++
	}
	return keys
}

func (m *NxN) ForA(a interface{}) []interface{} {
	if m.isEmptyA(a) {
		return []interface{}{}
	}

	return getKeys(m.forward[a])
}

func (m *NxN) ForB(b interface{}) []interface{} {
	if m.isEmptyB(b) {
		return []interface{}{}
	}

	return getKeys(m.backward[b])
}
