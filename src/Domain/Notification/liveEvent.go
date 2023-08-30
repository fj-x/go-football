package notification

const (
	StartLineup  = "LINEUP"
	Goal         = "GOAL"
	Card         = "CARD"
	Substitution = "SUBSTITUTION"
	Penalty      = "PENALTY"
)

type LiveEventTypes [5]string

type LiveEvent struct {
	Id        int32
	EventType string
	Minute    int8
	Player    string
}

func GetEventTypes() LiveEventTypes {
	return LiveEventTypes{StartLineup, Goal, Card, Substitution, Penalty}
}
