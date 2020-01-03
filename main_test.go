package main

import (
	"reflect"
	"testing"
)


func TestStringinperm(t *testing.T) {
	sl:=[]string{"ab","bc","a"}
	s:="abc"
	mp_test:=map[string]bool{"aabbc":false,
		"ababc":true,
		"abbca":false,
		"abcab":true,
		"bcaab":false,
		"bcaba":false,
	}
	mp:=map[string]bool{}
	Stringinperm(len(sl),sl,s,mp)
	if reflect.DeepEqual(mp,mp_test) != true {
		t.Errorf("s = %s is not present in %v",s,mp  )
	}
}



