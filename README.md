## go-ezform

[![Build Status](https://travis-ci.org/nathan-osman/go-ezform.svg?branch=master)](https://travis-ci.org/nathan-osman/go-ezform)
[![Coverage Status](https://coveralls.io/repos/github/nathan-osman/go-ezform/badge.svg?branch=master)](https://coveralls.io/github/nathan-osman/go-ezform?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nathan-osman/go-ezform)](https://goreportcard.com/report/github.com/nathan-osman/go-ezform)
[![GoDoc](https://godoc.org/github.com/nathan-osman/go-ezform?status.svg)](https://godoc.org/github.com/nathan-osman/go-ezform)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package provides an extremely simple method for parsing forms.

### Usage

Begin by creating a type to represent the inputs in your form:

```go
type RegistrationForm struct {
    Name       string
    NumTickets int64
    PayLater   bool
}
```

Next, create validation methods for each field you wish to validate.

```go
func (r RegistrationForm) ValidateName(v string) error {
    return ezform.IsNonZero(v)
}
```

Each validation method must conform to four requirements:

- the name of the method must be `Validate[field]`
- the method must use a value receiver
- the method must accept a single parameter of the field's type
- the method return type must be an `error`

The example validation method above ensures that `Name` has a non-zero value.

The HTML for this form might resemble the following:

```html
<form method="post">
    <input type="text" placeholder="Name" name="Name"><br>
    <input type="number" placeholder="# Tickets" name="NumTickets"><br>
    <input type="checkbox" name="PayLater"> Pay Later
    <button type="submit">Submit</button>
</form>
```

To parse and validate the form, use the following code in the handler:

```go
f := &RegistrationForm{}
fieldErrs, err := ezform.Parse(r, f, nil)
```

The first return value is a map of field names to errors that were encountered during validation. This map will be empty if no errors were encountered. The second return value will indicate any internal errors that were encountered during parsing.

### Validators

The example above introduced the `IsNonZero` validator, but there are others:

#### Contains

The `Contains` validator checks to see if the provided value is found within a slice or the keys of a map.

```go
validCountries := []string{
    "Australia",
    "Canada",
    "Spain",
}

func (r RegistrationForm) ValidateCountry(v string) error {
    return ezform.Contains(validCountries, v)
}
```

### Passing Parameters

Sometimes a validation method will need to access data from the scope invoking `Parse`. The third parameter to the function is used for this purpose:

```go
ezform.Parse(r, f, dbConn)
```

The validation methods may include an optional second parameter in their signature. This will be set to the value passed in `Parse`:

```go
func (r RegistrationForm) ValidateName(v string, dbConn *db.Conn) error {
    if !dbConn.isNameUnique {
        return errors.New("username is already in use")
    }
    return nil
}
```

If more than one parameter must be passed, a `struct` should be used.
