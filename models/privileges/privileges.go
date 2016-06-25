package privileges

import (
	"sort"
)

type (
	Role struct {
		Name        string
		Description string
		Privileges
	}

	Action     string
	Status     int
	Privileges map[Action][]Status
)

// Map converts the Privileges type to it's native type
func (p Privileges) Map() map[Action][]Status {
	return map[Action][]Status(p)
}

func (r *Role) Can(a Action, s Status) bool {
	// First, collect the statuses for each action
	stats := r.Map()[a]
	ints := asInts(stats)

	// Search the array for the value's index
	index := sort.SearchInts(ints, int(s))

	// Fail if the index is out of bounds
	if index >= len(stats) {
		return false
	}

	// Return true if the value at the index is the value we're looking for
	return stats[index] == s
}

func asInts(stats []Status) []int {
	ints := []int{}
	for _, status := range stats {
		ints = append(ints, int(status))
	}
	return ints
}
