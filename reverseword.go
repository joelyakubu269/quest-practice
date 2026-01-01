package main
import "fmt"

func main() {
	for r:=1; r<=2; r++ {
		if r==1 || r==2 {
			for c:=0;c<4;c++ {
				fmt.Print("*")
			}
		}else {
			for c:=0; c<4; c++ {
				if c == 0 || c == 3 {
					fmt.Print("*")
				}else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
	
