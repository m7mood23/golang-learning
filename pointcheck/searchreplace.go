package main 

import ("github.com/01-edu/z01"
"os")


func main(){

    args := os.Args[1:]  // get rid of first program we dont need it 


    if len(args) != 3 {  //check if there is 3 arguments, if not return nothing 
		return
	}

     input := args[0]       // "hella there"
    oldLetter := args[1]   // "a"
     newLetter := args[2]   // "o"

    //convert letters to runes
    oldrune:=  []rune(oldLetter)[0] // 'a'
    newrune:= []rune(newLetter)[0] // 'o'

    //now work with your logic  (please learn how to use )
    for _, char := range input{

        if char == oldrune{     //if it matches oldRune, print newRune
            z01.PrintRune(newrune) //print newRune
        }else{
            z01.PrintRune(char) //else, print the original char
        }
    }


z01.PrintRune('\n') //End with newline

}


