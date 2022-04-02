package store

// Added with Multiple Nested Types in the Same Struct
type Crew struct {
	Captain, FirstOfficer string
}

type RentalBoat struct {
	*Boat
	IncludeCrew bool
	*Crew
}

// Changed with Multiple Nested Types in the Same Struct
// func NewRentalBoat(name string, price float64, capacity int, motorized, crewed bool) *RentalBoat {
// 	return &RentalBoat{
// 		NewBoat(name, price, capacity, motorized),
// 		crewed,
// 	}
// }
func NewRentalBoat(name string, price float64, capacity int,
	motorized, crewed bool, captain, firstOfficer string) *RentalBoat {
	return &RentalBoat{
		NewBoat(name, price, capacity, motorized),
		crewed,
		&Crew{captain, firstOfficer},
	}
}
