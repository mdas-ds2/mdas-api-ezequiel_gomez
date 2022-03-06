package shared

type EventPayloadModel struct {
	AggregateId string `json:"aggregate"`
	EventType   string `json:"eventType"`
}
