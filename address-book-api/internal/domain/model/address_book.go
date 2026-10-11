// Package model contains the application entity declarations
package model

type Address struct {
	complement string
	block      string
	street     string
	number     int16
	city       string
	country    string
}
