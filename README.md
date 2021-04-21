# Checker
A Go package for validating formats

_please note that this package is currently in early beta testing and is subject to potential breaking changes_

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

```
func main() {
    if err := checker.Phone.Format("0300 000 0000"); err != nil {
        fmt.Println(err) // err = "invalid phone number format"
    }
}
```

### Mobile Checker

```
func main() {
    if checker.Phone.Mobile("07000000000") {
        fmt.Println("Number is a UK mobile number")
    }
}
```

### Premium Rate Checker

```
func main() {
    if checker.Phone.Premium("+44 845 0000 000") {
        fmt.Println("Number is a UK premium rate number")
    }
}
```

### Freephone Rate Checker

```
func main() {
    if checker.Phone.FreePhone("+44 800 0000 000") {
        fmt.Println("Number is a UK Freephone rate number")
    }
}
```