package models

type PayTicket struct {
	Date         string
	Amount       int
	Currency     string
	TrainID      string
	RefrenceID   string
	ClassBooking string
}

type PayCustomerDetails struct {
	Name  string
	Email string
}

type PayNotes struct {
	UserID        string
	DateOfBooking string
	ClassBooked   string
	TrainId       string
}
