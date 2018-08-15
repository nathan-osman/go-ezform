## go-ezform

[![Build Status](https://travis-ci.org/nathan-osman/go-ezform.svg?branch=master)](https://travis-ci.org/nathan-osman/go-ezform)
[![Coverage Status](https://coveralls.io/repos/github/nathan-osman/go-ezform/badge.svg?branch=master)](https://coveralls.io/github/nathan-osman/go-ezform?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nathan-osman/go-ezform)](https://goreportcard.com/report/github.com/nathan-osman/go-ezform)
[![GoDoc](https://godoc.org/github.com/nathan-osman/go-ezform?status.svg)](https://godoc.org/github.com/nathan-osman/go-ezform)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package provides an extremely simple method for parsing forms.

### Usage

Begin by importing the following packages:

```go
import (
    "github.com/nathan-osman/ezforms"
    "github.com/nathan-osman/ezforms/fields"
    "github.com/nathan-osman/ezforms/validators"
)
```

#### Creating a Form

Now create a type to represent the inputs in your form:

```go
type RegistrationForm struct {
    Name          *fields.String
    Age           *fields.Integer
    AcceptLicense *fields.Boolean
}
```

To simplify creation, add a function that initializes each field:

```go
func NewRegistrationForm() *RegistrationForm {
    return &RegistrationForm{
        Name:          fields.NewString(&validators.NonEmpty{}),
        Age:           fields.NewInteger(),
        AcceptLicense: fields.NewBoolean(&validators.Required{}),
    }
}
```

In the example above, the `NonEmpty` validator ensures that a non-empty string is provided for `Name` and the `Required` validator ensures that `AcceptLicense` is set to `true`.

#### Writing HTML

The HTML for the form above would resemble the following:

```html
<form method="post">
    <input type="text" placeholder="Name" name="Name"><br>
    <input type="number" placeholder="Age" name="Age"><br>
    <input type="checkbox" name="AcceptLicense"> Pay Later
    <button type="submit">Submit</button>
</form>
```

#### Validating a Request

To validate a POST request against the form, use the `ezform.Validate()` function:

```go
func registrationPage(w http.ResponseWriter, r *http.Request) {
    f := NewRegistrationForm()
    if err := ezform.Validate(r, f); err  {
        if err == ezform.ErrInvalid {
            // form contains invalid data
        } else {
            // internal error
        }
    }
    //...
}
```

After successful validation, each field will be set to the appropriate value from the request using the field's native type, which can be retrieved with the `Value()` method. For example, to retrieve the `Age` field from the example above:

```go
age := f.Age.Value()
```
