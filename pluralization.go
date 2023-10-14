package sentence

type pluralFunc func(int, []string) string

func pluralEnglish(n int, forms []string) string {
	if n < 0 {
		n *= -1
	}

	if n == 1 {
		return forms[0]
	}

	return forms[1]
}

func pluralRussian(n int, forms []string) string {
	if n < 0 {
		n *= -1
	}

	switch n % 100 {
	case 11, 12, 13, 14:
		return forms[2]
	}

	switch n % 10 {
	case 0, 5, 6, 7, 8, 9:
		return forms[2]
	case 1:
		return forms[0]
	default:
		return forms[1]
	}
}
