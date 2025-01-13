// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics_test

import (
	"testing"

	"github.com/tilotech/go-phonetics"
)

func TestMetaphoneEmptyString(t *testing.T) {
	if phonetics.EncodeMetaphone("") != "" {
		t.Errorf("Encode with empty string should return empty string")
	}
}

func TestMetaphoneEncode(t *testing.T) {
	assertMetaphoneEquals(t, "Donald", "TNLT")
	assertMetaphoneEquals(t, "Zach", "SX")
	assertMetaphoneEquals(t, "Campbel", "KMPBL")
	assertMetaphoneEquals(t, "Cammmppppbbbeeelll", "KMPBL")
	assertMetaphoneEquals(t, "David", "TFT")
	assertMetaphoneEquals(t, "Wat", "WT")
	assertMetaphoneEquals(t, "What", "WT")
	assertMetaphoneEquals(t, "Gaspar", "KSPR")
	assertMetaphoneEquals(t, "ggaspar", "KSPR")
}

func assertMetaphoneEquals(t *testing.T, source string, target string) {
	if phonetics.EncodeMetaphone(source) != target {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", source, phonetics.EncodeMetaphone(source), target)
	}
}

func BenchmarkEncodeMetaphone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonetics.EncodeMetaphone("Cammmppppbbbeeelll")
	}
}
