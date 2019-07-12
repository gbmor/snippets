package dedupe

func dedupe(list []string) []string {
	var out []string
	var seen map[string]bool

	for _, e := range list {
		if !seen[e] {
			out = append(out, e)
			seen[e] = true
		}
	}

	return out
}
