package phone

type PhoneNumber struct {
	ID     int    `db:"number_id" json:"number_id"`
	Number string `db:"number" json:"number"`
}
