package housing

const (
	TypeHouse      = int(1)
	TypeAppartment = int(2)
)

type Housing struct {
	UUID          string   `json:"uuid"`
	Type          int      `json:"type"`
	RoomAmount    int      `json:"room_amount"`
	SuiteAmount   int      `json:"suite_amount"`
	BedroomAmount int      `json:"bedroom_amount"`
	ParkingSpaces int      `json:"parking_spaces"`
	Area          int      `json:"area"`
	HasCabinets   bool     `json:"has_cabinets"`
	Description   string   `json:"description"`
	Fees          *float64 `json:"fees,omitempty"`
	HasConcierge  *bool    `json:"has_concierge,omitempty"`
	Street        string   `json:"street"`
	Number        int      `json:"number"`
	District      string   `json:"district"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	CEP           string   `json:"cep"`
	Floor         *int     `json:"floor,omitempty"`
	RentValue     float64  `json:"rent_value"`
}
