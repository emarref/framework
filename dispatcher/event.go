package dispatcher

// Event represents something that has occurred
type Event struct {
	Subject   interface{}
	Arguments map[string]interface{}
}

// NewEvent creates anew returns a new event for a subject
func NewEvent(subject interface{}) Event {
	event := Event{subject, make(map[string]interface{}, 0)}
	return event
}
