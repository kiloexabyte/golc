package main

// Leetcode #20
func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
            last := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]
            combination := string(last) + string(char)
            if(Valid(combination)) {
                continue
            } else {
                return false
            }

		}
	}

    return len(stack) == 0
}

func Valid(s string) bool {
    if s == "()" || s=="[]" || s=="{}"{
        return true
    }
    return false
}
