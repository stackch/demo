/*
 * see https://www.simtech-ag.ch/blog/golang/gomap
 * author: Daniel Schmutz
 * copyright: Simtech AG (https://www.simtech-ag.ch)
 */

// Package collections definitions
package collections

import (
	"sync"
)

// StdSyncMap struct
type StdSyncMap struct {
	*StdMap // anonymous field StdMap
	lock    sync.Mutex
}

// NewSyncMap func
func NewSyncMap() *StdSyncMap {
	m := new(StdSyncMap)
	m.StdMap = NewMap()
	return m
}

// NewSyncMapByParam func
func NewSyncMapByParam(paramMap map[string]interface{}) *StdSyncMap {
	m := new(StdSyncMap)
	m.StdMap = NewMapByParam(paramMap)
	return m
}

// IsNull func
func (m *StdSyncMap) IsNull() bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.StdMap.IsNull()
}

// Init func
func (m *StdSyncMap) Init() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StdMap.Init()
}

// HasKey func
func (m *StdSyncMap) HasKey(key string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.StdMap.HasKey(key)
}

// Value func
func (m *StdSyncMap) Value(key string) interface{} {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.StdMap.Value(key)
}

// String func
func (m *StdSyncMap) String() string {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.StdMap.String()
}
