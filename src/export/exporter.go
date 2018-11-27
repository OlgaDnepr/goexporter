package export

// Exporter is a generic interface which must be satisfied by defined custom types to export
type Exporter interface {
	Export() interface{}
}

// DaysOfWeek is an enum of days of week
type DaysOfWeek int

// These constants define a list of possible DaysOfWeek values
const (
	_ DaysOfWeek = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var daysOfWeekText = map[DaysOfWeek]string{
	DaysOfWeek(0): "",
	Sunday:        "Sunday",
	Monday:        "Monday",
	Tuesday:       "Tuesday",
	Wednesday:     "Wednesday",
	Thursday:      "Thursday",
	Friday:        "Friday",
	Saturday:      "Saturday",
}

// Export defines a special format of displaying DaysOfWeek values while exporting
func (d DaysOfWeek) Export() interface{} {
	return daysOfWeekText[d]
}
