package _4_json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}	`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(reflect.TypeOf(v), string(v))
	} else {
		t.Error(err)
	}
}
