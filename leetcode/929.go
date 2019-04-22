package leetcode

func emailHash(email string) int {
	hash := 0
	metAt := false
	for i := 0; i < len(email); i++ {
		if !metAt {
			if email[i] == '.' {
				continue
			} else if email[i] == '+' {
				i++
				for email[i] != '@' {
					i++
				}
				metAt = true
			} else if email[i] == '@' {
				metAt = true
			}
		}
		hash = hash*31 + int(email[i])
	}

	return hash
}

func numUniqueEmails(emails []string) int {
	hashes := map[int]bool{}

	for _, email := range emails {
		hash := emailHash(email)
		hashes[hash] = true
	}

	return len(hashes)
}
