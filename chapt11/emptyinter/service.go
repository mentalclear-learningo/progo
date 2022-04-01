package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string // makes this one non-comparable. Change is made.
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
