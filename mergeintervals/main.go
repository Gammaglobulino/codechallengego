package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type interval struct {
	start int
	end   int
}
type ByStart []interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start }

func main() {
	is := []interval{
		{1, 3},
		{5, 6},
		{7, 8},
	}

	newintval := interval{4, 6}
	newis, err := insertInterval(is, newintval)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%v interval is merged in %v.\n New slice of intervals is %v", newintval, is, newis)

}

//insert a new interval to a slice re-arranging the start-end based on the max of the end of intervals
// the complexity is N as we should travers all the N intervals of the source slice

func insertInterval(sofin []interval, newint interval) ([]interval, error) {

	if sofin == nil || len(sofin) == 0 {
		err := fmt.Errorf("Source silce is empty")
		return nil, err
	}
	res := []interval{}

	sofin = append(sofin, newint)
	sort.Sort(ByStart(sofin))
	n := len(sofin)
	res = append(res, sofin[0])
	for i := 1; i < n; i++ {
		if sofin[i].start <= res[len(res)-1].end {
			res[len(res)-1].end = int(math.Max(float64(res[len(res)-1].end), float64(sofin[i].end)))
		} else {
			res = append(res, sofin[i])
		}
	}
	return res, nil
}
