package rider

type FitnessLevel string

const (
	Beginner     FitnessLevel = "beginner"
	Intermediate FitnessLevel = "intermediate"
	Advanced     FitnessLevel = "advanced"
)

type Rider struct {
	HeightInches float64
	WeightLbs    float64
	Fitness      FitnessLevel
}

func (r *Rider) FitnessLevel() FitnessLevel {
	return r.Fitness
}
