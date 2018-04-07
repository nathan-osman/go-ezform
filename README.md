## go-ezform

[![Build Status](https://travis-ci.org/nathan-osman/go-ezform.svg?branch=master)](https://travis-ci.org/nathan-osman/go-ezform)
[![GoDoc](https://godoc.org/github.com/nathan-osman/go-ezform?status.svg)](https://godoc.org/github.com/nathan-osman/go-ezform)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package provides an extremely simple method for parsing forms.

### Usage

Begin by creating a type to represent the inputs in your form:

```go
type RegistrationForm struct {
    Name       string `form:"required"`
    NumTickets int64  `form:"required"`
    PayLater   bool
}
```

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
var f &RegistrationForm{}
m, _ := ezform.Parse(r, f)
if len(m) != 0 {
    // process validation errors
}
```

The first return value (`m`) is a map of field names to a list of errors that were encountered during validation.