package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnected.")
}