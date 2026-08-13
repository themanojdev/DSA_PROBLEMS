func isAnagram(s string, t string) bool {
    arr := [26]int{}

    s = strings.ToLower(s)
    t = strings.ToLower(t)

    s = strings.ReplaceAll(s," ","")
    t = strings.ReplaceAll(t," ","")

    for _,value := range s {
        arr[value-'a']++
    }

    for _,value := range t {
        arr[value-'a']--
    }

    for i,_ := range arr {
        if arr[i] != 0 {
            return false
        }
    }

    return true

}