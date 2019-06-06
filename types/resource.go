package types

type Resource struct {
	ID            string
	Status        string
	TransitStatus string
	CreateTime    int64
	TransitTime   int64
	UpdateTime    int64
}
