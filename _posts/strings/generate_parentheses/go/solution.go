package main

func generate(n int, row string, opened int, closed int, answer *[]string) {
	if len(row) == n*2 {
		*answer = append(*answer, row)
		return
	}

	if opened >= closed && opened < n {
		generate(n, row+"(", opened+1, closed, answer)
	}

	if opened > closed {
		generate(n, row+")", opened, closed+1, answer)
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	generate(n, "(", 1, 0, &res)

	return res
}
