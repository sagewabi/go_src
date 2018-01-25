package a1

import(
	"io/ioutil"
	"fmt"
	"strings"
)

func countStrings(filename string)(map[string]int){
	
	convert, err:= ioutil.ReadFile(filename)
    str := string(convert)
	if err != nil {
		fmt.Print(err)
	}
	 
	countWord := map[string]int{}
	for _,word := range strings.Fields(str){ // looping over the string looking for the initial word
		countWord[word]++                     //increments for eachone of the same word

	}

return countWord

}