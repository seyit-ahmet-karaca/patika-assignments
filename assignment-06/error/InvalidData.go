package error

type InvalidData struct {
	DataField string `json:"InvalidDataFields :"`
}

func (i *InvalidData) Error() string {
	return i.DataField
}