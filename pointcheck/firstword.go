package piscine 

func FirstWord(s string) string {
// We start at the first character of the string.
start := 0 

//Skip spaces (the first loop) 
for start < len(s)  && s[start] == ' '{

    start++ 
}
// for example it's  end is 4 saved (it's ramdom number yeha)  
end := start 

//Find the end of the word (the second loop)
//This means: keep moving until you hit a space or the end of the string.
for end < len(s) && s[end] != ' '{
    end++ 
}
//Now end = 8.

return s[start:end] + "\n"
}

a
