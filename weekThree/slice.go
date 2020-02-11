package main
import "fmt"
import "sort"
import "strconv"

func main() {
  //empty integer slice of size/length 3
  sli := make([]int, 3)
  var input string
  // for loops
  	// user enters integer
	// add the input integer to the slice
  var ind int = 0;
  for {
	fmt.Printf("number input: ")
	fmt.Scan(&input)
	inputNum, _ := strconv.Atoi(input)
	// quit only when user enters 'X' instead of integer
	if (input == "X" || input =="x") { break }
	
	if (ind < 3) {
		// for first three elements
		sli[ind] = inputNum
		ind += 1
	} else {
		// from 4th input: append
		sli = append(sli, inputNum)
	}

	// original slice
	// fmt.Println("original slice: ", sli)
	// sort the slice
	sortedSli := make([]int, len(sli))
	copy(sortedSli, sli)
	sort.Ints(sortedSli)
	// print the slice in sorted order
	fmt.Println("sorted slice: ", sortedSli)
  }
}