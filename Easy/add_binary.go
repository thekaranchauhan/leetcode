package main

func addBinary(a string, b string) string {
	n, m := len(a), len(b)
	carry := 0
	result := []byte{}

	// Iterate from the end of both strings
	i, j := n-1, m-1
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// Append the current bit to the result
		result = append(result, byte(sum%2)+'0')
		// Update the carry
		carry = sum / 2
	}

	// Reverse the result to get the correct binary number
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return string(result)
}
