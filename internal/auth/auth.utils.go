package auth

import (
	"encoding/json"
	"os"
)

type marshalledJson struct {
	// Other data fields in your original JSON structure
	Staffs []string `json:"staffs"`
}

func extractMapFromFile(filePath string) (map[string]interface{}, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data marshalledJson
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	extractedMap := make(map[string]interface{})

	for _, element := range data.Staffs {
		extractedMap[element] = element
	}

	return extractedMap, nil
}

var StaffStudentIdMap, err = extractMapFromFile("staff_student_id.json")
