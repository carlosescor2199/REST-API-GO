package model

type Contacto struct {
	ID int `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	PhoneNumber int `json:"PhoneNumber"`
	Email string `json:"Email"`
}

func (this *Contacto) setContacto(id int, firstname string, lastname string, phoneNumber int, email string) {

}


