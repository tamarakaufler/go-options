package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tamarakaufler/go-options/puppet"
)

func main() {
	pup := puppet.New(puppet.Features{
		Eyes:       "blue",
		HairLength: "short",
		HairStyle:  "afro",
		Height:     165,
		Weight:     57,
	})
	// fmt.Printf(">> Eyes: %s, skin: %s, hair: %+v, height: %0.2f, weight: %0.2f,\n", pup.Eyes(), pup.Skin(), pup.Hair(), pup.Height(), pup.Weight())

	// pup.Option(puppet.SetEyes("brown"))
	// pup.Option(puppet.SetSkin("bronze"))
	// pup.Option(puppet.SetHeight(150))
	// fmt.Printf(">> Eyes: %s, skin: %s, hair: %+v, height: %0.2f, weight: %0.2f,\n", pup.Eyes(), pup.Skin(), pup.Hair(), pup.Height(), pup.Weight())

	// restoreEyes := pup.Option(puppet.SetEyes("brown"))
	// restoreSkin := pup.Option(puppet.SetSkin("bronze"))
	// restoreHeight := pup.Option(puppet.SetHeight(150))
	// fmt.Printf(">> Eyes: %s, skin: %s, hair: %+v, height: %0.2f, weight: %0.2f,\n", pup.Eyes(), pup.Skin(), pup.Hair(), pup.Height(), pup.Weight())

	// pup.Option(restoreEyes)
	// pup.Option(restoreSkin)
	// pup.Option(restoreHeight)
	// fmt.Printf(">> Eyes: %s, skin: %s, hair: %+v, height: %0.2f, weight: %0.2f,\n", pup.Eyes(), pup.Skin(), pup.Hair(), pup.Height(), pup.Weight())

	fmt.Printf(">> Eyes: %s, skin: %s\n", pup.Eyes(), pup.Skin())

	// restoreAll := pup.Options(puppet.SetEyes("green"), puppet.SetSkin("fair"),
	// 	puppet.SetHairColour("ginger"), puppet.SetHeight(180), puppet.SetWeight(70))
	// fmt.Printf("<< Eyes: %s, skin: %s, hair: %+v, height: %0.2f, weight: %0.2f,\n", pup.Eyes(), pup.Skin(), pup.Hair(), pup.Height(), pup.Weight())

	// restoreAll = pup.Options(restoreAll)
	// fmt.Printf("<< Eyes: %s, skin: %s, hair: %+v, height: %0.2f, weight: %0.2f,\n", pup.Eyes(), pup.Skin(), pup.Hair(), pup.Height(), pup.Weight())

	// restoreAll = pup.Options(restoreAll)
	// fmt.Printf("<< Eyes: %s, skin: %s, hair: %+v, height: %0.2f, weight: %0.2f,\n", pup.Eyes(), pup.Skin(), pup.Hair(), pup.Height(), pup.Weight())

	restoreAll := pup.Options(puppet.SetEyes("green"), puppet.SetSkin("fair"))
	fmt.Printf("<< Eyes: %s, skin: %s\n", pup.Eyes(), pup.Skin())

	restoreAll = pup.Options(restoreAll)
	fmt.Printf("<< Eyes: %s, skin: %s\n", pup.Eyes(), pup.Skin())

	start := time.Now()
	restoreAll = pup.Options(restoreAll)
	log.Printf("Options took %s\n", time.Since(start))
	fmt.Printf("<< Eyes: %s, skin: %s\n", pup.Eyes(), pup.Skin())

	start = time.Now()
	restoreAll = pup.Options2(restoreAll)
	log.Printf("Options2 took %s\n", time.Since(start))
	fmt.Printf("== Eyes: %s, skin: %s\n", pup.Eyes(), pup.Skin())

}
