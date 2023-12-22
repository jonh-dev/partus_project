package tests

import (
	"testing"
)

type TestRunner interface {
	RunTests(t *testing.T)
}
