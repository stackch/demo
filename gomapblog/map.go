// Package collections definitions
package collections

import "fmt"

// StdMap struct
type StdMap struct {
	Map map[string]interface{}
}

// NewMap func
func NewMap() *StdMap {
	return NewMapByParam(nil)
}

// NewMapByParam func
func NewMapByParam(paramMap map[string]interface{}) *StdMap {
	m := new(StdMap)
	if paramMap == nil {
		m.Map = make(map[string]interface{})
	} else {
		m.Map = paramMap
	}
	return m
}

// IsNull func
func (m *StdMap) IsNull() bool {
	return m.Map == nil
}

// Init func
func (m *StdMap) Init() {
	m.Map = make(map[string]interface{})
}

// HasKey func
func (m *StdMap) HasKey(key string) bool {
	_, ok := m.Map[key]
	return ok
}

// Value func
func (m *StdMap) Value(key string) interface{} {
	return m.Map[key]
}

// ValueOrDefault func
func (m *StdMap) ValueOrDefault(key string, defaultValue interface{}) interface{} {
	val, ok := m.Map[key]
	if ok {
		return val
	}
	return defaultValue
}

// String func
func (m *StdMap) String() string {
	return fmt.Sprintf("%+v", m.Map)
}
