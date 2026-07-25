package garmin

import (
	"time"

	"github.com/mason-munnik/road-cycling-trainer/workout"
)

// PushableSession is what garmin needs from a session to push it.
type PushableSession interface {
	Workout() workout.Workout
	MarkPushed(externalRef string)
}

type CompletedActivity struct {
	ExternalRef               string // matches Session.ExternalRef
	Duration                  time.Duration
	Distance                  float64
	AvgPower, NormalizedPower float64
	AvgHeartRate              float64
	AvgCadence                float64
	// per-second stream data intentionally omitted for MVP — summary only first
}

type Client struct {
	// OAuth token/credential fields — shape depends on Connect Developer
	// Program approval details
}

func (c *Client) PushWorkout(w workout.Workout) (externalRef string, err error) {
	panic("not implemented")
}

func (c *Client) FetchActivity(externalRef string) (CompletedActivity, error) {
	panic("not implemented")
}
