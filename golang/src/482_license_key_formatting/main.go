package _482_license_key_formatting

func licenseKeyFormatting(S string, K int) string {
	res := ""
	sectionCount := 0
	for i := len(S); i >= 0; i-- {
		if S[i] != '-' {
			res += string(S[i])
			sectionCount++
			if sectionCount == K {
				res += "-"
				sectionCount = 0
			}
		}
	}
	i := 0
	j := len(res) - 1
	runes := []rune(S)
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
