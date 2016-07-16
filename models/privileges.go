package privileges

type (
	Role struct {
		Name        string
		ID          int64
		Description string
		Privileges
	}

	Action     string
	Status     int
	Privileges map[Action][]Status
)

func (r *Role) Can(a Action, s Status) bool {
	return false
}
