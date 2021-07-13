package strain

type Ints []int

type Strings []string

type Lists []Ints

func (x Lists) Keep(test func([]int) bool) Lists {
	keep := make(Lists, 0, len(x))
	for _, n := range x {
		if test(n) {
			keep = append(keep, n)
		}
	}
	return keep
}

func (x Strings) Keep(test func(string) bool) Strings {
	keep := make(Strings, 0, len(x))
	for _, n := range x {
		if test(n) {
			keep = append(keep, n)
		}
	}
	return keep
}

func (x Strings) Discard(test func(string) bool) Strings {
	discard := make(Strings, 0, len(x))
	for _, n := range x {
		if !test(n) {
			discard = append(discard, n)
		}
	}
	return discard
}

func (x Ints) Keep(test func(int) bool) Ints {
	if x == nil {
		return nil
	}
	keep := make(Ints, 0, len(x))
	for _, n := range x {
		if test(n) {
			keep = append(keep, n)
		}
	}
	return keep
}

func (x Ints) Discard(test func(int) bool) Ints {
	if x == nil {
		return nil
	}
	discard := make(Ints, 0, len(x))
	for _, n := range x {
		if !test(n) {
			discard = append(discard, n)
		}
	}
	return discard
}
