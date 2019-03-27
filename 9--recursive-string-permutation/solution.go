package recursivestringpermutation

func Perm(in string) []string {
	n := len(in)
	if n <= 1 {
		return []string{in}
	}

	perms := []string{}

	firstCharsSeen := map[byte]bool{}
	for i := 0; i < n; i++ {
		firstChar := in[i]
		if firstCharsSeen[firstChar] {
			continue
		}
		firstCharsSeen[firstChar] = true

		remainingChars := in[:i] + in[i+1:]
		for _, tail := range Perm(remainingChars) {
			arrangement := string(firstChar) + tail
			perms = append(perms, arrangement)
		}
	}

	return perms
}
