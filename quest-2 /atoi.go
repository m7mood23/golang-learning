package piscine 

func Atoi(s string) int {

  if s == ""{
    return 0 
  }
  result := 0 
  sign := 1
  start := 0
  //position check 0 if there is minus - 
  if s[0] == '-'{
    sign = -1
    start = 1 
  }else if s[0] == '+'{ // also plus check position if there s[0] there + 
    sign = 1 
    start = 1 
  }

  for i := start ; i < len(s); i++{

    char := s[i]  
    if char >= '0' && char <='9'{
      result = result * 10 + int(char-'0')
    }else {
      return 0 
    }
  }
  return result * sign 
}
