package main

import (
	"Stas-sH/clietntForCts/internal/client"
	"Stas-sH/clietntForCts/internal/countryBreedGroup"
	"Stas-sH/clietntForCts/pkg/writefile"

	"log"
)

func main() {
	client := client.NewClient()
	cats, err := client.GetCats()
	if err != nil {
		log.Fatal(err)
	}

	result := countryBreedGroup.CountryBreedGroup(cats)

	if err = writefile.WriteFile(result); err != nil {
		log.Fatal(err)
	}

}
