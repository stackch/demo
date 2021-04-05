// Package collections definitions
package collections

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"

	"std.ch/errors"
)

// StdSyncLinkedMap struct
type StdSyncLinkedMap struct {
	*StdSyncMap // anonymous field StdLinkedMap
	linkedKeys  []string
	lock        sync.Mutex
}

// NewSyncLinkedMap func
func NewSyncLinkedMap() *StdSyncLinkedMap {
	m := new(StdSyncLinkedMap)
	m.StdSyncMap = NewSyncMap()
	m.linkedKeys = make([]string, 0)
	return m
}

// NewSyncLinkedMapByParam func
func NewSyncLinkedMapByParam(paramMap map[string]interface{}) *StdSyncLinkedMap {
	m := new(StdSyncLinkedMap)
	m.StdSyncMap = NewSyncMapByParam(paramMap)
	m.linkedKeys = make([]string, 0)
	return m
}

// Add func
func (m *StdSyncLinkedMap) SetValue(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if !m.HasKey(key) {
		m.StdSyncMap.Map[key] = value
		m.linkedKeys = append(m.linkedKeys, key)
	}
}

// LinkedKeys func
func (m *StdSyncLinkedMap) LinkedKeys() ([]string, int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	length := len(m.linkedKeys)
	return m.linkedKeys, length
}

// LinkedKey func
func (m *StdSyncLinkedMap) LinkedKey(pos int) (string, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	length := len(m.linkedKeys)
	if pos < 0 || pos >= length {
		return "", errors.NewErrorOutOfRange("pos " + strconv.Itoa(pos) + " is out of range")
	}
	return m.linkedKeys[pos], nil
}

// ToString func
func (m *StdSyncLinkedMap) ToString() string {
	m.lock.Lock()
	defer m.lock.Unlock()
	var buffer bytes.Buffer
	buffer.WriteRune('[')
	for i, key := range m.linkedKeys {
		if i > 0 {
			buffer.WriteRune(',')
		}
		buffer.WriteString(key)
		buffer.WriteRune(':')
		buffer.WriteString(fmt.Sprintf("%+v", m.Map[key]))
		i++
	}
	buffer.WriteRune(']')
	return buffer.String()
}
