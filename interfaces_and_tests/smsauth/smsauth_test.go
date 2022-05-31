package smsauth

import (
	"testing"
)

func TestVerifySMS(t *testing.T) {
	tests := []struct {
		description     string
		inputChallenger stubChallenger
		inputPrompter   stubPrompter
		expected        bool
	}{
		{
			"success sms verification",
			stubChallenger{phoneNumber: "77773080226", code: "1235"},
			stubPrompter("1235"),
			true,
		},
		{
			"invalid sms verification",
			stubChallenger{phoneNumber: "77773080226", code: "1235"},
			stubPrompter("5421"),
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := VerifySMS(test.inputChallenger, test.inputPrompter)
			if got != test.expected {
				t.Fatalf("expected %v got %v\n", test.expected, got)
			}
		})
	}

}

type stubChallenger struct {
	phoneNumber string
	code        string
}

func (c stubChallenger) Challenge() (string, string) {
	return c.code, c.phoneNumber
}

type stubPrompter string

func (p stubPrompter) Prompt(_ string) string {
	return string(p)
}
