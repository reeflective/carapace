package state

func copyStringSlice(s []string) (c []string) {
	if s != nil {
		c = make([]string, len(s))
		copy(c, s)
	}
	return
}
