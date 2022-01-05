package cirks

import "strconv"

type Rule struct {
	to       int
	location string
	name     string
}

func getRules() map[int]Rule {
	var rules = make(map[int]Rule)

	rules[3] = Rule{
		to:       19,
		location: "up",
		name:     "cangoroo",
	}

	rules[9] = Rule{
		to:       11,
		location: "up",
		name:     "seal",
	}

	rules[14] = Rule{
		to:       30,
		location: "up",
		name:     "bear",
	}

	rules[16] = Rule{
		to:       1,
		location: "down",
		name:     "girl",
	}

	rules[25] = Rule{
		to:       7,
		location: "down",
		name:     "pig",
	}

	rules[34] = Rule{
		to:       97,
		location: "up",
		name:     "lion",
	}

	rules[36] = Rule{
		to:       45,
		location: "up",
		name:     "dog",
	}

	rules[40] = Rule{
		to:       58,
		location: "up",
		name:     "rooster",
	}

	rules[50] = Rule{
		to:       32,
		location: "up",
		name:     "dog",
	}

	rules[52] = Rule{
		to:       68,
		location: "up",
		name:     "dog",
	}

	rules[53] = Rule{
		to:       8,
		location: "down",
		name:     "monkey",
	}

	rules[57] = Rule{
		to:       63,
		location: "up",
		name:     "rabbit",
	}

	rules[59] = Rule{
		to:       80,
		location: "up",
		name:     "horse",
	}

	rules[65] = Rule{
		to:       83,
		location: "up",
		name:     "dows",
	}

	rules[71] = Rule{
		to:       51,
		location: "down",
		name:     "cat",
	}

	rules[72] = Rule{
		to:       94,
		location: "up",
		name:     "tiger",
	}

	rules[87] = Rule{
		to:       67,
		location: "down",
		name:     "monkey",
	}

	rules[92] = Rule{
		to:       108,
		location: "up",
		name:     "goose",
	}

	rules[96] = Rule{
		to:       116,
		location: "up",
		name:     "man",
	}

	rules[102] = Rule{
		to:       81,
		location: "down",
		name:     "man",
	}

	rules[107] = Rule{
		to:       23,
		location: "down",
		name:     "goat",
	}

	rules[112] = Rule{
		to:       90,
		location: "down",
		name:     "clown",
	}

	rules[117] = Rule{
		to:       98,
		location: "down",
		name:     "lady",
	}

	rules[119] = Rule{
		to:       101,
		location: "down",
		name:     "cat",
	}

	return rules
}

func getTo(rule Rule) int {
	return rule.to
}

func getRuleName(rule Rule) string {
	return rule.location + " " + rule.name + " to " + strconv.Itoa(rule.to)
}
