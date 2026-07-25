package training

import (
	"time"

	"github.com/mason-munnik/road-cycling-trainer/rider"
	"github.com/mason-munnik/road-cycling-trainer/workout"
)

type FitnessProfile interface {
	FitnessLevel() rider.FitnessLevel
}

type Odometer interface {
	LogMiles(miles float64)
	TotalMiles() float64
}

type SessionStatus string

const (
	Scheduled SessionStatus = "scheduled"
	Pushed    SessionStatus = "pushed"    // sent to wearable, awaiting completion
	Completed SessionStatus = "completed" // activity synced back
	Skipped   SessionStatus = "skipped"
)

type Session struct {
	Date        time.Time
	Workout     workout.Workout
	Status      SessionStatus
	ExternalRef string // Garmin's workout/schedule ID once pushed
}

type TrainingPlan struct {
	GoalEvent *time.Time // nil for open-ended (rolling) plans
	Sessions  []Session
}

type WeekPlan struct {
	SessionCount   int
	PrimaryZone    workout.Zone
	SecondaryZone  workout.Zone // optional, zero value = unused
	IsRecoveryWeek bool
}

// Generator produces the week-by-week shape of a plan; turning that shape
// into actual Sessions/Workouts is a separate, strategy-agnostic step.
type Generator interface {
	Generate(profile FitnessProfile) []WeekPlan
}

// PhasedGenerator splits the weeks between now and GoalEvent into
// base/build/peak/taper phases.
type PhasedGenerator struct {
	GoalEvent time.Time
}

func (g PhasedGenerator) Generate(profile FitnessProfile) []WeekPlan {
	panic("not implemented")
}

// RollingGenerator has no fixed end date: a sawtooth progression (build
// weeks followed by a recovery week) bounded by Horizon weeks at a time.
type RollingGenerator struct {
	Horizon int
}

func (g RollingGenerator) Generate(profile FitnessProfile) []WeekPlan {
	panic("not implemented")
}
