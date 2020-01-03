package main

import (
	"fmt"
	"strings"
)

func main(){
	sl:=[]string{"ab","bc","a"}
	s:="abc"
	mp:=map[string]bool{}
	Stringinperm(len(sl),sl,s,mp)
	fmt.Println(mp)
}

// check if a string is included permuting the strings of a slice
// return a map with the string perm and a bool value indicating the presence of the string
// using the recursive version of heap's permutation algo
// the complexity is N!

func Stringinperm(k int,sl []string, s string, mp map[string]bool){
	if k==1{
		mp[strings.Join(sl,"")]=strings.Contains(strings.Join(sl,""),s)
		}else{
		Stringinperm(k-1,sl,s,mp)
		for i:=0;i < k-1; i ++{
			if k % 2 == 0 {
				sl[i],sl[k-1]=sl[k-1],sl[i]
			}else{
				sl[0],sl[k-1]=sl[k-1],sl[0]
			}
			Stringinperm(k-1,sl,s,mp)
		}
	}

}
// suggested optimization
// we can optimize using moaization tecnique:
// we generates a map of true permutations (string found included on a set of strings),
//we encapsulate the search on a hashe-map (closure) of pre allocated strings
//using this we can save the N! permutations using a costant time access scan of an hash map.

