package hashset

import "strings"

var placeholder = struct{}{}

// StrHashset is a hashset of strings
type StrHashset struct {
	values map[string]struct{}
}

// NewStrHashset constructs a StrHashet with
// optional initial values
func NewStrHashset(initVals ...string) StrHashset {
	numVals := 0
	if initVals != nil {
		numVals = len(initVals)
	}

	set := StrHashset{
		values: make(map[string]struct{}, len(initVals)),
	}

	if numVals > 0 {
		set.Add(initVals...)
	}

	return set
}

// Add inserts one or many strings into the hashset.
// Duplicate entries will be automatically ignored.
func (set *StrHashset) Add(items ...string) {
	for _, item := range items {
		set.values[item] = placeholder
	}
}

// Size returns the number of values currently
// held within the hashset
func (set StrHashset) Size() int {
	return len(set.values)
}

// IsEmpty returns true if there are no
// values in the hashset, false otherwise
func (set StrHashset) IsEmpty() bool {
	return len(set.values) == 0
}

// Values returns a slice of all entries
// within the hashset
func (set StrHashset) Values() []string {
	values := make([]string, set.Size())

	count := 0
	for item := range set.values {
		values[count] = item
		count++
	}

	return values
}

// Equals returns true if every entry in two
// hashsets are the same
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

// Contains returns true if the set contains
// the input string
func (set StrHashset) Contains(val string) bool {
	_, found := set.values[val]
	return found
}

// ContainsSubstr returns true if any entry in the
// set contains the given string entry
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

// Clear removes all entries from the existing hashset.
// This is a constant time operation.
func (set *StrHashset) Clear() {
	set.values = make(map[string]struct{})
}
