package main

import "fmt"

type JhoveSig struct {
	Jhove struct {
		RepInfo []JRepInfo `json:"repInfo"`
	} `json:"jhove"`
}

type JRepInfo struct {
	SigMatch []string `json:"sigMatch"`
}

func (jhoveSig JhoveSig) String() string {
	return jhoveSig.Jhove.RepInfo[0].SigMatch[0]
}

type JhoveNil struct{}

func Parse(sig string, b []byte) (any, error) {
	switch sig {
	case "WAVE-hul":
		{
			jWave, err := ParseWave(b)
			if err != nil {
				return JhoveNil{}, err
			}
			return jWave, nil
		}

	default:
		return JhoveNil{}, fmt.Errorf("no matching signature for %s", sig)
	}

}
