package elements

type Organizer uint8

// Declare all the possible organizer for a sql query
const (
	OrganizerLimit Organizer = iota
	OrganizerOffset
	OrganizerOrder
	OrganizerGroup
	OrganizerHaving
)

func OrganizersSortOrder() [5]Organizer {
	return [...]Organizer{
		OrganizerGroup,
		OrganizerHaving,
		OrganizerOrder,
		OrganizerLimit,
		OrganizerOffset,
	}
}
