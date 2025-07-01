package vehicle

import "fmt"

type Vehicle interface {
	GetWheels() int
	Description() string
}

func TalkAbout(v Vehicle) string {
	return fmt.Sprintf("I really love my vehicle, it is %v and has the correct amount of wheels %v", v.Description(), v.GetWheels())
}

type Sale interface {
	Description() string
	Price() int
}

func Advertise(s Sale) string {
	if s.Price() > 0 {
		return fmt.Sprintf("I'm selling a %v for £%v", s.Description(), s.Price())
	} else {
		return "price has not been set, check back later :)"
	}
}
