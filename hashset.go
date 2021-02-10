package hashset

import "strings"

var placeholder = struct{}{}

type StrHashset struct {
	values map[string]struct{}
}

func NewStrHashset(initVals ...string) StrHashset {
	numVals := len(initVals)
	set := StrHashset{
		values: make(map[string]struct{}, len(initVals)),
	}

	if numVals > 0 {
		set.Add(initVals...)
	}

	return set
}

func (set StrHashset) GetSingleValue() string {
	retVal := ""
	for val := range set.values {
		retVal = val
		break
	}

	return retVal
}

func (set *StrHashset) Add(items ...string) {
	for _, item := range items {
		set.values[item] = placeholder
	}
}

func (set StrHashset) Size() int {
	return len(set.values)
}

func (set StrHashset) IsEmpty() bool {
	return len(set.values) == 0
}

func (set StrHashset) Values() []string {
	values := make([]string, set.Size())

	count := 0
	for item := range set.values {
		values[count] = item
		count++
	}

	return values
}

func (set StrHashset) Equals(set2 StrHashset) bool {
	if set.Size() != set2.Size() {
		return false
	}

	for val := range set.values {
		if !set2.Contains(val) {
			return false
		}
	}

	return true
}

func (set StrHashset) Contains(val string) bool {
	_, found := set.values[val]
	return found
}

func (set StrHashset) ContainsSubstr(substr string) bool {
	for val := range set.values {
		if strings.Contains(val, substr) {
			return true
		}
	}

	return false
}

func (set StrHashset) ContainsSubstrIgnoreCase(substr string) bool {
	substr = strings.ToLower(substr)

	for val := range set.values {
		if strings.Contains(strings.ToLower(val), substr) {
			return true
		}
	}

	return false
}

func (set *StrHashset) Clear() {
	set.values = make(map[string]struct{})
}
