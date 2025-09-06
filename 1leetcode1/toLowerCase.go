func toLowerCase(s string) string {
    result := ""

    for _, char := range s{  //  'H' 'e' 'l' 'l' 'o' 

        if char >= 'A' && char <= 'Z'{
            newchar:= char + 32 
            result+= string(newchar)
        }else{
            result+= string(char)
        }

    }
    return result 
}
