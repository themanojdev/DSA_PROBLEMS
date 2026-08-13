func isPalindrome(s string) bool {
    
    result := []rune{}

    for _ , value := range s {
        if unicode.IsLetter(value) || unicode.IsDigit(value) {
            result = append(result,unicode.ToLower(value))
        }
    }

    start,end := 0,len(result)-1

    for start < end {

        if result[start] != result[end] {
            return false
        }

        start++
        end--
    }

    return true
}