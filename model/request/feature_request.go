package request

import (
	"time"
)

type FeatureRequest struct {
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
