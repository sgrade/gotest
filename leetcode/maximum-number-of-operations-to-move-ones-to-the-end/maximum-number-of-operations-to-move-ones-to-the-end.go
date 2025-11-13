package maximumnumberofoperationstomoveonestotheend

func maxOperations(s string) int {
	ops, zero_intervals, i := 0, 0, len(s)-1
	for i >= 0 {
		if s[i] == '0' {
			zero_intervals++
			for i >= 0 && s[i] == '0' {
				i--
			}
		} else {
			ones := 0
			for i >= 0 && s[i] == '1' {
				ones++
				i--
			}
			ops += ones * zero_intervals
		}
	}
	return ops
}
