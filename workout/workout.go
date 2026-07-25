package workout

type StepType string

const (
	Warmup   StepType = "warmup"
	Interval StepType = "interval"
	Recovery StepType = "recovery"
	Cooldown StepType = "cooldown"
	Rest     StepType = "rest"
)

type DurationType string

const (
	DurationTime     DurationType = "time"     // seconds
	DurationDistance DurationType = "distance" // meters
	DurationOpen     DurationType = "open"     // lap-button / manual advance
)

type TargetType string

const (
	TargetNone      TargetType = "none"
	TargetZone      TargetType = "zone" // qualitative, no device value yet
	TargetPower     TargetType = "power"
	TargetHeartRate TargetType = "heart_rate"
	TargetCadence   TargetType = "cadence"
)

// Zone is a qualitative effort level, used until Rider carries an FTP or
// max/resting heart rate to convert to concrete power/heart-rate numbers.
type Zone string

const (
	ZoneRecovery  Zone = "recovery"
	ZoneEndurance Zone = "endurance"
	ZoneTempo     Zone = "tempo"
	ZoneThreshold Zone = "threshold"
	ZoneVO2Max    Zone = "vo2max"
	ZoneAnaerobic Zone = "anaerobic"
)

type Target struct {
	Type      TargetType
	Zone      Zone    // set when Type == TargetZone
	Low, High float64 // set when Type == TargetPower/TargetHeartRate/TargetCadence
}

type Step struct {
	Type          StepType
	DurationType  DurationType
	DurationValue float64
	Target        Target
	Repeat        int    // >1 for interval blocks
	Steps         []Step // nested steps when Repeat > 1
}

type Workout struct {
	Name  string
	Sport string // "cycling" for now
	Steps []Step
}
