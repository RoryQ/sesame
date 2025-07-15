package domain

type Address struct {
	Pref         string
	Street       string
	StringValues []string
	Date         Date1
	Line1        *string
}

type Date1 struct {
	Year int
}

func (a *Address) GetLine1() string {
	if a != nil && a.Line1 != nil {
		return *a.Line1
	}
	return ""
}
