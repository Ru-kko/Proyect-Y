package messagging

type QueueAction string

const (
	Delete QueueAction = "delete"
	Update QueueAction = "update"
	Save   QueueAction = "create"
)
