package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
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
	{"Andy", "+91000000000", "2000-01-01"},
	{"Snehal", "+91000000000", "2000-12-17"},
}

type Message struct {
	Name              string
	IsBirthday        bool
	PurchaseThreshold int
	DailyDiscount     int
	BirthdayDiscount  int
}

func LoadTemplate(filePath string, templateName string) (*template.Template, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	bytes, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return template.New(templateName).Parse(string(bytes))
}

func main() {
	// Marketing emailer => if today is *NOT* a subscriber b'day: 10% off; 15% off
	todayPurchaseThreshold := 10000
	dailyDiscount := 10

	birthdayDiscount := 5

	currentMonthDay := time.Now().Format("01-02")

	// tmpl, err := LoadTemplate("/Users/gaurav/Developer/Training/Go/examples/18-templating/after/message.tpl", "promotion-email")
	// tmpl, err := LoadTemplate("./message.tpl", "promotion-sms")

	tmpl, err := LoadTemplate("./email_message.html.tpl", "promotion-email")

	// tmpl, err := LoadTemplate("./launch.tpl", "launch-email")

	if err != nil {
		fmt.Println("Unable to parse the template:", err)
		os.Exit(1)
	}

	for _, subscriber := range subscribers {
		doB := strings.Split(subscriber.DoB, "-")
		monthDay := doB[1] + "-" + doB[2]

		sub := Message{
			Name:              subscriber.Name,
			IsBirthday:        currentMonthDay == monthDay,
			PurchaseThreshold: todayPurchaseThreshold,
			DailyDiscount:     dailyDiscount,
			BirthdayDiscount:  birthdayDiscount,
		}

		err = tmpl.Execute(os.Stdout, sub)

		if err != nil {
			fmt.Println("Unable to execute the template:", err)
			os.Exit(1)
		}
	}
}
