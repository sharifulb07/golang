package main
import "fmt"


// time complexity O(n) and space complexity O(1)
func sortColors(colors[]int){
	low:=0;
	mid:=0;
	high:=len(colors)-1

	for mid<=high{
		switch colors[mid]{

		case 0:
			colors[low], colors[mid]=colors[mid], colors[low]

			low++
			mid++
		case 1:
			mid++
		case 2:
			colors[mid], colors[high]=colors[high], colors[mid]
			high--
		}
	}
}

func main(){


	colors:=[]int{0,1,0,1,2,0,1,2,2,0,1,0,1,0,2,1,1}
sortColors(colors)
fmt.Println(colors)
}