package main


import (
"reflect"
"testing"
)


func TestInsertinterval(t *testing.T) {

	is:=[]interval{
		{1,3},
		{5,6},
		{7,8},
	}
	testing_res:=[]interval{
		{1,3},
		{4,6},
		{7,8},
	}

	newintval:=interval{4,6}
	newis:=insertInterval(is,newintval)

	if reflect.DeepEqual(testing_res,newis) != true {
		t.Errorf("cannot insert = %v in %v",newintval,is)
	}
}
