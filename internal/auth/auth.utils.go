package auth

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type AuthUtils interface {
	IsStudentIdInMap(studentId string) bool
}

type authUtilsImpl struct {
	staffStudentIdMap map[string]interface{}
}

func NewAuthUtils() AuthUtils {
	staffStudentIdMap, err := extractMapFromFile("./config/staffs/staff.json")
	if err != nil {
		panic(err)
	}

	return &authUtilsImpl{
		staffStudentIdMap: staffStudentIdMap,
	}
}

func IsEmailChulaStudent(email string) bool {
	studentId := extractStudentIdFromEmail(email)
	if len(studentId) != 10 {
		return false
	}

	year, err := strconv.ParseInt(studentId[:2], 10, 64)
	if err != nil {
		return false
	}

	return year <= 67 && strings.HasSuffix(email, "@student.chula.ac.th")
}

func (u *authUtilsImpl) IsStudentIdInMap(email string) bool {
	studentId := extractStudentIdFromEmail(email)

	_, ok := u.staffStudentIdMap[studentId]
	return ok
}

func extractStudentIdFromEmail(email string) string {
	// Example: "6932203021@student.chula.ac.th" -> "6932203021"
	return email[:10]
}

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
