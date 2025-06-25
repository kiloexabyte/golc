package greedy

// Leetcode #678
func checkValidString(s string) bool {
    leftMin, leftMax := 0, 0

    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' {
            leftMin++
            leftMax++
        } else if c == ')' {
            leftMin--
            leftMax--
        } else { // c == '*'
            leftMin--
            leftMax++
        }

        if leftMax < 0 {
            return false
        }
        if leftMin < 0 {
            leftMin = 0
        }
    }

    return leftMin == 0
}