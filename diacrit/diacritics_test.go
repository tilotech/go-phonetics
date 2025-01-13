package diacrit_test

import (
	"sync"
	"testing"

	"github.com/tilotech/go-phonetics/diacrit"
)

func TestRemoval(t *testing.T) {
	cases := map[string]string{
		"a":                      "a",
		"ä":                      "a",
		"Ä":                      "A",
		"ß":                      "ss",
		"Č":                      "C",
		"Ł":                      "L",
		"Ä möre ȼꝋmpleᶍ ⱸxamplé": "A more complex example",
	}

	for inp, expected := range cases {
		t.Run(inp, func(t *testing.T) {
			actual := diacrit.Normalize(inp)
			if actual != expected {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})
	}
}

func TestRemovalParallel(t *testing.T) {
	wg := sync.WaitGroup{}

	expected := "A more complex example"
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			actual := diacrit.Normalize("Ä möre ȼꝋmpleᶍ ⱸxamplé")
			if actual != expected {
				t.Errorf("expected %v, got %v", expected, actual)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkNormalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diacrit.Normalize("Ä möre ȼꝋmpleᶍ ⱸxamplé")
	}
}
