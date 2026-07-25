package bike

type Bike struct {
	Make         string
	Model        string
	WeightLbs    float64
	MileageTotal float64
}

func (b *Bike) LogMiles(miles float64) {
	b.MileageTotal += miles
}

func (b *Bike) TotalMiles() float64 {
	return b.MileageTotal
}
