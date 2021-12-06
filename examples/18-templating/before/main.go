package main

import (
	"fmt"
	"strings"
	"time"
)

type Subscriber struct {
	Name        string
	PhoneNumber string
	DoB         string
}

var subscribers = []Subscriber{
	{"Gaurav", "+91000000000", "2000-05-07"},
	{"Nikhil", "+91000000000", "2000-01-01"},
	{"Shafeeq", "+91000000000", "2000-12-06"},
}

func main() {
	// Marketing emailer => if today is *NOT* a subscriber b'day: 10% off; 25% off

	todayPurchaseThresold := 10000
	dailyDiscount := 10

	birthdayDiscount := 5

	currentMonthDay := time.Now().Format("01-02")

	for _, subscriber := range subscribers {
		doB := strings.Split(subscriber.DoB, "-")
		monthDay := doB[1] + "-" + doB[2]

		if monthDay == currentMonthDay {
			fmt.Printf(`Hi %s,

Happy birthday to you!

We have a special offer running in store to give you. And if you purchase above INR %d amount we will give you a %d%% discount.

Also, on top of that, to make your day special we also offer a %d%% additional discount on any purchase.`, subscriber.Name, todayPurchaseThresold, dailyDiscount, birthdayDiscount)
		} else {
			fmt.Printf(`Hi %s

We have a special offer running in store to give you. And if you purchase above INR %d amount we will give you a %d%% discount.`, subscriber.Name, todayPurchaseThresold, dailyDiscount)
		}
		fmt.Println()
		fmt.Println("-------------------------------------------------------------")
	}
}
