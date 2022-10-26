package incidents

import (
	context "context"
)

// Controller for incidents
type Controller struct {
}

// Incident struct
type Incident struct {
	ID int `json:"id"`
}
