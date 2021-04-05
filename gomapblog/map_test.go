/*
 * see https://www.simtech-ag.ch/blog/golang/gomap
 * author: Daniel Schmutz
 * copyright: Simtech AG (https://www.simtech-ag.ch)
 */

// Package collections
package collections

import (
	"fmt"
	"testing"
)

func Test_StdMapNewMap(t *testing.T) {
	m := NewMap()	
	m.Map["one"] = 1
	m.Map["two"] = 2
	m.Map["three"] = 3
	m.Map["four"] = 4

	fmt.Printf("m = %s\n", m.String())
}

func Test_StdMapNewMapByParam(t *testing.T) {
	params := make(map[string]interface{})
	params["one"] = "one"
	m := NewMapByParam(params)	
	fmt.Printf("m = %s\n", m.String())
}

func TestStdMap_HasKey(t *testing.T) {
	m := NewMap()	
	m.Map["one"] = 1
	m.Map["two"] = 2
	m.Map["three"] = 3
	m.Map["four"] = 4

	if !m.HasKey("one") {
		t.Error("m.HasKey(\"one\") is false");
	}

	if m.HasKey("One") {
		t.Error("m.HasKey(\"One\") is true");
	}
}

func Test_StdMapValueOrDefault(t *testing.T) {
	m := NewMap()	
	m.Map["one"] = 1
	if m.ValueOrDefault("one", 0) != 1 {
		t.Error("m.ValueOrDefault(\"one\") != 1");		
	}
	if m.ValueOrDefault("One", 0) != 0 {
		t.Error("m.ValueOrDefault(\"One\") != 0");		
	}
	if m.ValueOrDefault("one", 0) == 0 {
		t.Error("m.ValueOrDefault(\"one\") == 0");		
	}
}
