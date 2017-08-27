package async

type Event interface {
	Process()
}

var (
	EventQueue chan Event
)

func init() {
	EventQueue = make(chan Event, 100)
}
