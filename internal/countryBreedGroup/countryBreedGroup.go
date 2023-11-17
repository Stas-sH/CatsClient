package countryBreedGroup

import (
	"Stas-sH/clietntForCts/pkg/cat"
	"sort"
	"sync"
)

func CountryBreedGroup(cats []cat.Cat) map[string][]string {
	group := make(map[string][]string)
	for i := 0; i < len(cats); i++ {
		catCountry := cats[i].Country
		catBreed := cats[i].Breed
		if _, inMap := group[catCountry]; inMap {
			group[catCountry] = append(group[catCountry], catBreed)
		} else {
			group[catCountry] = []string{catBreed}
		}
	}
	return sortByBreedsLen(group)
}

func sortByBreedsLen(group map[string][]string) map[string][]string {
	wg := &sync.WaitGroup{}
	var mu sync.Mutex
	for key, breeds := range group {
		wg.Add(1)
		func(key string, breeds []string) {
			defer wg.Done()
			mu.Lock()
			breeds = sortBreeds(breeds)
			group[key] = breeds
			mu.Unlock()
		}(key, breeds)
	}
	wg.Wait()
	return group
}

func sortBreeds(breeds []string) []string {
	sort.Slice(breeds, func(i, j int) bool {
		return len(breeds[i]) < len(breeds[j])
	})
	return breeds
}
