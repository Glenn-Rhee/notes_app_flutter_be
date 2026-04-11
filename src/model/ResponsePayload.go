package model

// Added type status
type Status string

const (
	Success Status = "Success"
	Failed  Status = "Failed"
)

type ResponsePayload struct {
	Code    int
	Status  Status
	Message string
	Data    any
}