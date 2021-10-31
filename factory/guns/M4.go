package guns

type M4 struct {
	Gun
}

func NewM4() IGun {
	return &M4{
		Gun: Gun{
			Name: "Maverick",
			Power: 4,
		},
	}
}


