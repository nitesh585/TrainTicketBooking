package models

type PayTicket struct {
	Date         string
	Amount       int
	Currency     string
	TrainID      string
	RefrenceID   string
	ClassBooking string
	Name         string
	Email        string
	User_id      string
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

type PaymentResponse struct {
	Amount   int
	Currency string
	Customer struct {
		Name  string
		Email string
	}
	PaymentLink string
	ExpireTime  string
	Notes       struct {
		User_id       string
		DateOfBooking string
		ClassBooked   string
		TrainId       string
	}
}
