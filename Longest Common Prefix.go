package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
  var y = []string{} 
  if strs[0][0] == strs[1][0]{
    fmt.Println("Hello, World!")
  }else if strs[1][1] == strs[2][1] == strs[0][1]{
    
  } 
    return y; 
}

func main() {
  var x = []string{"Volvo", "Volvo", "Volvo", "Volvo"}
	longestCommonPrefix(x) 
}
/*
[] create new arr x 
[] if the first char in each str in the arr are not equal return empty ""  
[] else :
[] push the first char in the arr x 
[] if the second chars equal push it to the 
*/
