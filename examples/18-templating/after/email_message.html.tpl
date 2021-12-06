<html>
<body>
Hi {{.Name}},
</br>
{{if .IsBirthday}}
Happy birthday to you!
{{end}}
We have a special offer running in store to give you. And if you purchase above INR {{.PurchaseThreshold}} amount we will give you a {{.DailyDiscount}}% discount.
</br>
{{if .IsBirthday}}
Also, on top of that, to make your day special we also offer a {{.BirthdayDiscount}} additional discount on any purchase.
{{end}}
</body>
</html>
