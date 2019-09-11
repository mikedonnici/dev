package strain

type Ints []int
type Lists [][]int
type Strings []string

func (xi Ints) Keep(filter func(int) bool) Ints {
	var res Ints
	for _, i := range xi {
		if filter(i) {
			res = append(res, i)
		}
	}
	return res
}

func (xi Ints) Discard(filter func(int) bool) Ints {
	var res Ints
	for _, i := range xi {
		if !filter(i) {
			res = append(res, i)
		}
	}
	return res
}

func (xl Lists) Keep(filter func([]int) bool) Lists {
	var res Lists
	for _, l := range xl {
		if filter(l) {
			res = append(res, l)
		}
	}
	return res
}

func (xs Strings) Keep(filter func(string) bool) Strings {
	var res Strings
	for _, s := range xs {
		if filter(s) {
			res = append(res, s)
		}
	}
	return res
}
