package piscine 

func ThirdTimeIsACharm(str string) string {

    // if string is empty → just return newline
    if str == "" {
        return "\n"
    }  

    // if string has less than 3 chars → no 3rd char → newline
    if len(str) < 3 {
        return "\n"
    }

    result := ""
    // start at index 2 (the 3rd char), then jump by 3 each time
    // collect every 3rd character
    for i := 2; i < len(str); i += 3 {
        result = result + string(str[i])
    }

    // add newline at the end
    return result + "\n"
}
