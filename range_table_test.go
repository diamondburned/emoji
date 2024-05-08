package emoji_test

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/alecthomas/assert/v2"
	"libdb.so/go-emoji"
	"libdb.so/go-emoji/data"
)

// Sample of emoji newly introduced by version
var samplesByPropertyAndVersion = map[data.Property]map[emoji.Version]string{
	data.Emoji: {
		emoji.V1:  "😀😃😄", // 1F600, 1F603, 1F604
		emoji.V2:  "🗨",   // 1F5E8
		emoji.V3:  "🤣🤥🤤", // 1F923, 1F925, 1F924
		emoji.V4:  "♀♂⚕", // 2640, 2642, 2695
		emoji.V5:  "🤩🤪🤭", // 1F929, 1F92A, 1F92D
		emoji.V11: "🥰🥵🥶", // 1F970, 1F975, 1F976
		emoji.V12: "🥱🤎🤍", // 1F971, 1F90E, 1F90D
	},
}

// Combined sample of specified version and all versions below it
func combinedSample(prop data.Property, v emoji.Version) string {
	samples := samplesByPropertyAndVersion[prop]
	sample := ""
	for _, v2 := range emoji.AllVersions {
		if v2 >= v {
			break
		}
		sample += samples[v2]
	}
	return sample
}

func TestRangeTables(t *testing.T) {
	for _, prop := range data.AllProperties {
		for _, v := range emoji.AllVersions {
			t.Run(fmt.Sprintf("%v_%v", prop, v), func(t *testing.T) {
				rt := v.RangeTable(prop)
				sample := combinedSample(prop, v)
				for _, r := range sample {
					inRange := unicode.In(r, rt)
					assert.True(t, inRange, "expected %v (%X) to be in %v range for %v, but was not", string(r), prop, v, r)
				}
			})
		}
	}
}
