package test

import (
	"github.com/zwjlink/meterutil/meter"
	"testing"
)

func TestAccount(t *testing.T) {
	ac, err := meter.ParseAddress("0x6a797722f929e24eac9acF8b9fFa74346dB3167c")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(ac.String())

}
