package controller

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

// Index of incidents
// GET
func (c *Controller) Index(ctx context.Context) (incidents []*Incident, err error) {
	return []*Incident{}, nil
}

// Show incident
// GET :id
func (c *Controller) Show(ctx context.Context, id int) (incident *Incident, err error) {
	return &Incident{
		ID: id,
	}, nil
}
