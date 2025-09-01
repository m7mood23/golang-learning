func lengthOfLastWord(s string) int {
    
    lastletter:= len(s)-1 
    count:=0 
    
   
        
        for lastletter >= 0 && s[lastletter] == ' '{
            lastletter--
        }

        start:= lastletter 

        for start >= 0 && s[start] != ' '{
            start-- 
            count++ 
        }
    
    return count
}
