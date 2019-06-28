package gonxn

type NxN struct {
    forward  map[interface{}]map[interface{}]bool
    backward map[interface{}]map[interface{}]bool
}

func (m *NxN) ensureA(a interface{}) {
    if m.forward == nil {
        m.forward = make(map[interface{}]map[interface{}]bool)
    }

    if m.forward[a] == nil {
        m.forward[a] = make(map[interface{}]bool)
    }
}

func (m *NxN) ensureB(b interface{}) {
    if m.backward == nil {
        m.backward = make(map[interface{}]map[interface{}]bool)
    }

    if m.backward[b] == nil {
        m.backward[b] = make(map[interface{}]bool)
    }
}

func (m *NxN) Link(a interface{}, b interface{}) {
    m.ensureA(a)
    m.ensureB(b)
    m.forward[a][b] = true
    m.backward[b][a] = true
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

func (m *NxN) IsLinked(a interface{}, b interface{}) bool {
    if m.isEmptyAB(a, b) {
        return false
    }

    return m.forward[a][b]
}

func getKeys(m map[interface{}]bool) []interface{} {
    keys := make([]interface{}, 0, len(m))
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
