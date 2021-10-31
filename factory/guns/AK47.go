package guns

type AK47 struct {
	Gun
}

func NewAK47() IGun {
	return &AK47{
		Gun: Gun{
			Name: "AK47",
			Power: 4,
		},
	}
}


