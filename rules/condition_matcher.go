package rules

func NewConditionMatcher() *ConditionMatcher {
	return &ConditionMatcher{}
}

type ConditionMatcher struct {
}

func (c *ConditionMatcher) Evaluate(matching string, results []bool) bool {
	valid := false

	hits := 0
	for _, checkResult := range results {
		if checkResult {
			hits++
		}
	}

	if (matching == "or" || matching == "") && hits > 0 {
		valid = true
	} else if matching == "and" && hits == len(results) {
		valid = true
	}

	//ToDo: Extend condition engine to support complex cases e.g. "1 AND (2 OR 3)"

	return valid
}