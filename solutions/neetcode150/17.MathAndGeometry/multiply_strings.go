package mathandgeometry

// Leetcode #43
func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    n1 := reverseString(num1)
    n2 := reverseString(num2)
    res := make([]int, len(n1)+len(n2))

    for i := 0; i < len(n1); i++ {
        for j := 0; j < len(n2); j++ {
            d1 := int(n1[i] - '0')
            d2 := int(n2[j] - '0')
            sum := d1*d2 + res[i+j]
            res[i+j] = sum % 10
            res[i+j+1] += sum / 10
        }
    }

    // Reverse result
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }

    // Skip leading zeros
    i := 0
    for i < len(res) && res[i] == 0 {
        i++
    }

    var result string
    for ; i < len(res); i++ {
        result += string(res[i] + '0')
    }

    return result
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
