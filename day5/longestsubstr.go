package main
import "fmt"


func lengthOfSubStr(s string)int{
	charIndex:=make(map[byte]int)

	maxlength:=0
	left:=0

	for right:=0;right<len(s);right++{
		char:=s[right]

		if idx, found:=charIndex[char]; found && idx>=left{
			left=idx+1
		}

		charIndex[char]=right

		if right-left+1>maxlength{
			maxlength=right-left+1
			
		}
	}

	return maxlength

}

func main(){

	s:="abcabdc"

	fmt.Println(lengthOfSubStr(s))

}