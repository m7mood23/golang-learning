package piscine 

func CheckNumber(arg string) bool {

    numberexit := false 
    for _, char := range arg {

       if char >= '0' && char <= '9'{
        numberexit = true 
       }
    } 
      return numberexit  
}
