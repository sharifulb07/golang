package main 
import "fmt"

func firstUnique(s string)int{

	freq:=make(map[rune]int)

	// hashing map initializing
	for _, char:=range s{
		freq[char]++
	}

	// checking first character
	for i, char:=range s{
		if freq[char]==1{
			return  i 
		}
	}
	return  -1
}

func main(){
	examples:=[]string{"leetcode", "aabb", "love", "puremeasementp"}

	for _, s:=range examples{
		result:=firstUnique(s)

		if result !=-1{
			fmt.Println("First Unique Character :", result, string(s[result]))
		} else{
			fmt.Println("Result with no unique character is : ", s )
		}
	}
}