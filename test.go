import (
	"os"
	"fmt"
)
func main() {
	f, err := os.Open("/tmp")
	if err != nil {
		panic("Error open file")
	}
	defer f.Close()
	v := "a"
	switch V {
	case "a":
		fmt.Println("handling task a")
		// many lines of code
		return
	case "b":
		fmt.Println("handling task a")
		// many lines of code
		return
	default:
	}
	fmt.Println("handling default")
	//many lines of code
	type read func(string) 
	func getF(config string) read {
		readS
	}
}