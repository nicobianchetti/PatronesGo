package main

type m4 struct {
	gun
}

func newM4() iGun {
	return &m4{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
