package ghevent

import (
	"encoding/json"
	"testing"
)

func TestJSONDecodeInteger(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"int":123}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.IntegerField == nil {
		t.Error("ts.IntegerField was <nil>")
	}
	if *ts.IntegerField != 123 {
		t.Errorf("*ts.IntegerField was %v (should have been 123)", *ts.IntegerField)
	}
	// Test that decoded field is nil when it is missing from JSON input
	jsonStr = `{"string":"hello"}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.IntegerField != nil {
		t.Error("ts.IntegerField was not <nil> when field was missing from input")
	}
}

func TestJSONDecodeFloat(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"float":1.23}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.FloatField == nil {
		t.Error("ts.FloatField was <nil>")
	}
	if *ts.FloatField != 1.23 {
		t.Errorf("*ts.FloatField was %v (should have been 1.23)", *ts.FloatField)
	}
	// Test that decoded field is nil when it is missing from JSON input
	jsonStr = `{"string":"hello"}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.FloatField != nil {
		t.Error("ts.FloatField was not <nil> when field was missing from input")
	}
}

func TestJSONDecodeString(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"string":"hello"}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.StringField == nil {
		t.Error("ts.StringField was <nil>")
	}
	if *ts.StringField != "hello" {
		t.Errorf("*ts.StringField was %v (should have been \"hello\")", *ts.StringField)
	}
	// Test that decoded field is nil when it is missing from JSON input
	jsonStr = `{"int":123}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.StringField != nil {
		t.Error("ts.StringField was not <nil> when field was missing from input")
	}
}

func TestJSONDecodeBool(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"bool":true}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.BoolField == nil {
		t.Error("ts.BoolField was <nil>")
	}
	if *ts.BoolField != true {
		t.Errorf("*ts.BoolField was %v (should have been True)", *ts.BoolField)
	}
	jsonStr = `{"bool":false}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.BoolField == nil {
		t.Error("ts.BoolField was <nil>")
	}
	if *ts.BoolField != false {
		t.Errorf("*ts.BoolField was %v (should have been False)", *ts.BoolField)
	}
	// Test that decoded field is nil when it is missing from JSON input
	jsonStr = `{"int":123}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.BoolField != nil {
		t.Error("ts.BoolField was not <nil> when field was missing from input")
	}
}

func TestJSONDecodeStringSlice(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"stringslice":["hello","world"]}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.StringSliceField == nil {
		t.Error("ts.StringSliceField was <nil>")
	}
	if len(ts.StringSliceField) != 2 {
		t.Errorf("len(ts.StringSliceField) was %v (should have been 2)", len(ts.StringSliceField))
	}
	if ts.StringSliceField[0] != "hello" {
		t.Errorf("ts.StringSliceField[0] was \"%v\" (should have been \"hello\")", ts.StringSliceField[0])
	}
	if ts.StringSliceField[1] != "world" {
		t.Errorf("ts.StringSliceField[1] was \"%v\" (should have been \"world\")", ts.StringSliceField[1])
	}
	// Test that an empty array in the JSON data translates to an empty slice but not a nil slice
	jsonStr = `{"stringslice":[]}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	// is it a nil slice?  (shouldn't be)
	if ts.StringSliceField == nil {
		t.Error("ts.StringSliceField was <nil> when input was an empty array")
	}
	// is it an empty slice?  (should be)
	if len(ts.StringSliceField) != 0 {
		t.Error("ts.StringSliceField was not an empty slice when input was an empty array")
	}
	// Test that decoded field is a nil slice when it is missing from JSON input
	jsonStr = `{"int":123}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.StringSliceField != nil {
		t.Error("ts.StringSliceField was not <nil> when field was missing from input")
	}
}

func TestJSONDecodeIntSlice(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"intslice":[123,456]}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.IntSliceField == nil {
		t.Error("ts.IntSliceField was <nil>")
	}
	if len(ts.IntSliceField) != 2 {
		t.Errorf("len(ts.IntSliceField) was %v (should have been 2)", len(ts.IntSliceField))
	}
	if ts.IntSliceField[0] != 123 {
		t.Errorf("ts.IntSliceField[0] was \"%v\" (should have been 123)", ts.IntSliceField[0])
	}
	if ts.IntSliceField[1] != 456 {
		t.Errorf("ts.IntSliceField[1] was \"%v\" (should have been 456)", ts.IntSliceField[1])
	}
	// Test that an empty array in the JSON data translates to an empty slice but not a nil slice
	jsonStr = `{"intslice":[]}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	// is it a nil slice?  (shouldn't be)
	if ts.IntSliceField == nil {
		t.Error("ts.IntSliceField was <nil> when input was an empty array")
	}
	// is it an empty slice?  (should be)
	if len(ts.IntSliceField) != 0 {
		t.Error("ts.IntSliceField was not an empty slice when input was an empty array")
	}
	// Test that decoded field is a nil slice when it is missing from JSON input
	jsonStr = `{"int":123}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.IntSliceField != nil {
		t.Error("ts.IntSliceField was not <nil> when field was missing from input")
	}
}

func TestJSONDecodeFloatSlice(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"floatslice":[1.23,4.56]}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.FloatSliceField == nil {
		t.Error("ts.FloatSliceField was <nil>")
	}
	if len(ts.FloatSliceField) != 2 {
		t.Errorf("len(ts.FloatSliceField) was %v (should have been 2)", len(ts.FloatSliceField))
	}
	if ts.FloatSliceField[0] != 1.23 {
		t.Errorf("ts.FloatSliceField[0] was \"%v\" (should have been 1.23)", ts.FloatSliceField[0])
	}
	if ts.FloatSliceField[1] != 4.56 {
		t.Errorf("ts.FloatSliceField[1] was \"%v\" (should have been 4.56)", ts.FloatSliceField[1])
	}
	// Test that an empty array in the JSON data translates to an empty slice but not a nil slice
	jsonStr = `{"floatslice":[]}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	// is it a nil slice?  (shouldn't be)
	if ts.FloatSliceField == nil {
		t.Error("ts.FloatSliceField was <nil> when input was an empty array")
	}
	// is it an empty slice?  (should be)
	if len(ts.FloatSliceField) != 0 {
		t.Error("ts.FloatSliceField was not an empty slice when input was an empty array")
	}
	// Test that decoded field is a nil slice when it is missing from JSON input
	jsonStr = `{"int":123}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.FloatSliceField != nil {
		t.Error("ts.FloatSliceField was not <nil> when field was missing from input")
	}
}

func TestJSONDecodeBoolSlice(t *testing.T) {
	ts := TestStruct{}
	jsonStr := `{"boolslice":[true,false]}`
	err := json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.BoolSliceField == nil {
		t.Error("ts.BoolSliceField was <nil>")
	}
	if len(ts.BoolSliceField) != 2 {
		t.Errorf("len(ts.BoolSliceField) was %v (should have been 2)", len(ts.BoolSliceField))
	}
	if ts.BoolSliceField[0] != true {
		t.Errorf("ts.BoolSliceField[0] was \"%v\" (should have been true)", ts.BoolSliceField[0])
	}
	if ts.BoolSliceField[1] != false {
		t.Errorf("ts.BoolSliceField[1] was \"%v\" (should have been false)", ts.BoolSliceField[1])
	}
	// Test that an empty array in the JSON data translates to an empty slice but not a nil slice
	jsonStr = `{"boolslice":[]}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	// is it a nil slice?  (shouldn't be)
	if ts.BoolSliceField == nil {
		t.Error("ts.BoolSliceField was <nil> when input was an empty array")
	}
	// is it an empty slice?  (should be)
	if len(ts.BoolSliceField) != 0 {
		t.Error("*ts.BoolSliceField was not an empty slice when input was an empty array")
	}
	// Test that decoded field is a nil slice when it is missing from JSON input
	jsonStr = `{"int":123}`
	ts = TestStruct{}
	err = json.Unmarshal([]byte(jsonStr), &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.BoolSliceField != nil {
		t.Error("ts.BoolSliceField was not <nil> when field was missing from input")
	}
}
