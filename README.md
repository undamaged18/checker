# Checker
A Go package for validating formats

_please note that this package is currently in early beta testing and is subject to potential breaking changes_

### Email format validator
```
func main() {
    if err := checker.Email.Format("valid@example.com"); err != nil {
        fmt.Println(err) // err = "invalid email format"
    }
}
```

### Email Host validator
```
func main() {
    if err := checker.Email.Host("invalid@hostdoesnotexist.com"); err != nil {
        fmt.Println(err) // err = "domain search returned 0 results"
    }
}
```