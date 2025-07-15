package model

type AddressModel struct {
	Pref      string
	Street    []int
	IntValues []int
	Date      *Date1
	Line1     *string
}

type Date1 struct {
	Year string
}
