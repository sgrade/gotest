// 1541. Minimum Insertions to Balance a Parentheses String
// https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/

package minimuminsertionstobalanceaparenthesesstring

func minInsertions(s string) int {
	insertions := 0
	pendingClosings := 0

	for _, c := range s {
		if c == '(' {
			if pendingClosings%2 == 1 {
				insertions++
				pendingClosings--
			}
			pendingClosings += 2
		} else {
			pendingClosings--
			if pendingClosings < 0 {
				insertions++
				pendingClosings += 2
			}
		}
	}

	return insertions + pendingClosings
}
