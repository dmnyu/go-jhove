package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

var testFile = filepath.Join("test", "jhove-wave.json")
var jSig JhoveSig
var b []byte

func TestJhove(t *testing.T) {
	t.Run("test get a signature", func(t *testing.T) {
		var err error
		b, err = os.ReadFile(testFile)
		if err != nil {
			t.Error(err)
		}
		jSig = JhoveSig{}
		if err := json.Unmarshal(b, &jSig); err != nil {
			panic(err)
		}
		want := "WAVE-hul"
		got := jSig.String()
		if want != got {
			t.Errorf("Wanted %s, got %s", want, got)
		}
	})

	t.Run("test parse a wave file", func(t *testing.T) {
		jAny, err := Parse(jSig.String(), b)
		if err != nil {
			t.Error(err)
		}

		jWave := jAny.(*JhoveWave)

		want := "Jhove"
		got := jWave.Jhove.Name
		if want != got {
			t.Errorf("Wanted %s, got %s", want, got)
		}
	})
}
