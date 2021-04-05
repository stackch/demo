/*
 * see https://www.simtech-ag.ch/blog/golang/gomap
 * author: Daniel Schmutz
 * copyright: Simtech AG (https://www.simtech-ag.ch)
 */

// Package collections definitions
package collections

import (
	"bytes"
	"fmt"
	"strconv"

	"std.ch/errors"
)

// StdLinkedMap struct
type StdLinkedMap struct {
	StdMap     // anonymous field StdMap
	linkedKeys []string
}

// NewLinkedMap func
func NewLinkedMap() *StdLinkedMap {
	m := new(StdLinkedMap)
	m.StdMap = *NewMap()
	m.linkedKeys = make([]string, 0)
	return m
}

// NewLinkedMapByParam func
func NewLinkedMapByParam(paramMap map[string]interface{}) *StdLinkedMap {
	m := new(StdLinkedMap)
	m.StdMap = *NewMapByParam(paramMap)
	m.linkedKeys = make([]string, 0)
	return m
}

// Add func
func (m *StdLinkedMap) SetValue(key string, value interface{}) {
	if !m.HasKey(key) {
		m.StdMap.Map[key] = value
		m.linkedKeys = append(m.linkedKeys, key)
	}
}

// LinkedKeys func
func (m *StdLinkedMap) LinkedKeys() ([]string, int) {
	length := len(m.linkedKeys)
	return m.linkedKeys, length
}

// LinkedKey func
func (m *StdLinkedMap) LinkedKey(pos int) (string, error) {
	length := len(m.linkedKeys)
	if pos < 0 || pos >= length {
		return "", errors.NewErrorOutOfRange("pos " + strconv.Itoa(pos) + " is out of range")
	}
	return m.linkedKeys[pos], nil
}

// String func
func (m *StdLinkedMap) String() string {
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
