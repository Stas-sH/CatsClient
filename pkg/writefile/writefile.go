package writefile

import (
	"encoding/json"
	"os"
)

func WriteFile(data map[string][]string) error {
	file, err := os.Create("out.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
