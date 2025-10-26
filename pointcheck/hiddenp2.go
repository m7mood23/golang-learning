package main 

import ("os"
"github.com/01-edu/z01")

func main(){

args:= os.Args[1:]


if len(args) != 2{
    return 
}

s1:= args[0]
s2:= args[1]

if s1 == ""{
    z01.PrintRune('1')
     z01.PrintRune('\n')
    return 
}

i := 0
for _, char := range s2 {
    if i == len(s1) {
        continue 
    } else {
        if char == rune(s1[i]) {
            i++
        }
    }
}
//-------------------------------
 if i == len(s1) {
        z01.PrintRune('1')
    } else {
        z01.PrintRune('0')
    }
    z01.PrintRune('\n')

}
