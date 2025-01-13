package phonetics_test

import (
	"testing"

	"github.com/tilotech/go-phonetics"
)

func TestCologneEncode(t *testing.T) {
	cases := map[string]string{
		"":                    "",
		"Aeijouy":             "0",
		"H":                   "",
		"Bp":                  "1",
		"Dt":                  "2",
		"Fvwph":               "3",
		"Gkq":                 "4",
		"C":                   "8",
		"Caca":                "4",
		"Chch":                "4",
		"Ckck":                "4",
		"Clcl":                "4585",
		"Coco":                "4",
		"Cqcq":                "4",
		"Crcr":                "4787",
		"Cucu":                "4",
		"Cxcx":                "4848",
		"X":                   "48",
		"Xx":                  "4848",
		"L":                   "5",
		"Ł":                   "5",
		"Mn":                  "6",
		"R":                   "7",
		"Sz":                  "8",
		"Scsc":                "8",
		"Szsz":                "8",
		"Dcdc":                "8",
		"Dsds":                "8",
		"Dzdz":                "8",
		"Tctc":                "8",
		"Tsts":                "8",
		"Tztz":                "8",
		"Kxkx":                "4848",
		"Qxqx":                "4848",
		"Breschnew":           "17863",
		"Müller-Lüdenscheidt": "65752682",
		"Heinz Classen":       "068586",
		"小鹿":                  "",
	}

	for inp, expected := range cases {
		t.Run(inp, func(t *testing.T) {
			actual := phonetics.EncodeCologne(inp)
			if actual != expected {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})
	}
}

func BenchmarkEncodeCologne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonetics.EncodeCologne("Müller-Lüdenscheidt")
	}
}
