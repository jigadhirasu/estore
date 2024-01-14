package main

import "github.com/open-policy-agent/opa/rego"

func main() {
	rego.Query("data.example.allow")
}
