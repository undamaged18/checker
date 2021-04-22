# Checker
**Version 0.1.0**\
**Release Date: 22/04/2021**

A Go package for validating formats of email addresses, phone numbers and UK postcodes

## Email Addresses
### Format validator
Validates the email address based on the [HTML 5.1 W3C Recommendations for Form validation](https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemail)
```
func main() {
    if err := checker.Email.Format("valid@example.com"); err != nil {
        fmt.Println(err) // err = "invalid email format"
    }
}
```

### Host validator
Validates that the host part of the email address is for a valid domain by checking that an MX record exists for the domain
```
func main() {
    if err := checker.Email.Host("invalid@hostdoesnotexist.com"); err != nil {
        fmt.Println(err) // err = "domain search returned 0 results"
    }
}
```

## Phone Numbers
Phone number validation is based around UK phone numbers and allows for common formats such as `0800`, `+44800`, `+44 800` and `(0123) 456 7890`
### Format validator
Format validator uses [ttacon / libphonenumber](https://github.com/ttacon/libphonenumber) Golang port of the [google / libphonenumber](https://github.com/google/libphonenumber)

**National format**
```
func main() {
    num, err := checker.Phone().Format("+44 7777 555 555", "GB", checker.NATIONAl)
    if err != nil {
        fmt.Println(err) // err = "invalid phone number format"
    }
    fmt.Println(num) // returns phone number in the format XXXXX XXXXXX
}
```
**International format**
```
func main() {
    num, err := checker.Phone().Format("+44 7777 555 555", "GB", checker.INTERNATIONAL)
    if err != nil {
        fmt.Println(err) // err = "invalid phone number format"
    }
    fmt.Println(num) // returns phone number in the format +XX XXXX XXXXXX
}
```

### Mobile Checker
Check that the phone number supplied is a UK Mobile number
```
func main() {
    if checker.Phone.Mobile("07000000000") {
        fmt.Println("Number is a UK mobile number")
    }
}
```

### Premium Rate Checker
Check that the phone number supplied is a UK Premium rate number
```
func main() {
    if checker.Phone.Premium("+44 845 0000 000") {
        fmt.Println("Number is a UK premium rate number")
    }
}
```

### Freephone Rate Checker
Check that the phone number supplied is a UK Freephone number
```
func main() {
    if checker.Phone.FreePhone("+44 800 0000 000") {
        fmt.Println("Number is a UK Freephone rate number")
    }
}
```

## Postcodes
### Format validator
```
func main() {
    err := checker.Postcode().Format("AA1 1BB")
    if err != nil {
        fmt.Println(err) // returns invalid UK postcode format
    }
}
```