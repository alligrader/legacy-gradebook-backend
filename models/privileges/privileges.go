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

func (r *Role) fetchPrivileges() {
	// Sets the privileges on this object
}

func (r *Role) Can(a Action, s Status) bool {
	if r.Privileges == nil {
		r.fetchPrivileges()
	}
	if statusSlice, ok := r.Privileges[a]; ok {
		// search statuses for the status. If not,
		statuses := ToInts(statusSlice)
		if !sort.IntsAreSorted(statuses) {
			sort.Ints(statuses)
		}

		if location := sort.SearchInts(statuses, int(s)); location < len(statuses) && statuses[location] == int(s) {
			return true
		}
		return false
	}
	return false
}

func ToInts(s []Status) []int {
	a := []int{}
	for _, status := range s {
		a = append(a, int(status))
	}
	return a
}
