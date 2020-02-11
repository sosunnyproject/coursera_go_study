package main
import "fmt"
import "strings"
import "os"
import "bufio"

func main() {
	var input string
	fmt.Printf("original input: ")
	// fmt.Scanf("%d", &input)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan();
	line := scanner.Text()
	
	fmt.Println("Input is", line)
	input = strings.ToLower(line)
	// fmt.Println(input)
	if (strings.Contains(input, "a") && strings.HasSuffix(input, "n") && strings.HasPrefix(input, "i")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
