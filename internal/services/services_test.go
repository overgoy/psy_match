package services

import (
	"testing"
)

func TestGetCompatibilityData(t *testing.T) {
	tests := []struct {
		myMBTI, otherMBTI   string
		expectedPercentage  int
		expectedDescription string
	}{
		{
			myMBTI: "ENTP", otherMBTI: "ENTP",
			expectedPercentage:  IdentityRelation,
			expectedDescription: "",
		},
		{
			myMBTI: "ENTP", otherMBTI: "ISFP",
			expectedPercentage:  DualityRelation,
			expectedDescription: "",
		},
		{
			myMBTI: "ENTP", otherMBTI: "ISFJ",
			expectedPercentage:  ConflictRelation,
			expectedDescription: "",
		},
		{
			myMBTI: "ENTP", otherMBTI: "ENTJ",
			expectedPercentage:  QuasiIdentityRelation,
			expectedDescription: "",
		},
		{
			myMBTI: "ESFJ", otherMBTI: "ENTP",
			expectedPercentage:  ActivationRelation,
			expectedDescription: "",
		},
	}

	for _, test := range tests {
		t.Run(test.myMBTI+"_"+test.otherMBTI, func(t *testing.T) {
			percentage, description := GetCompatibilityData(test.myMBTI, test.otherMBTI)
			if percentage != test.expectedPercentage {
				t.Errorf("Ожидаемый процент %d, получено %d", test.expectedPercentage, percentage)
			}
			if description != test.expectedDescription {
				t.Errorf("Ожидаемое описание '%s', получено '%s'", test.expectedDescription, description)
			}
		})
	}
}

func TestGetZodiacCompatibility(t *testing.T) {
	tests := []struct {
		myZodiac, otherZodiac string
		expectedCompatibility int
	}{
		{
			myZodiac: "Овен", otherZodiac: "Овен",
			expectedCompatibility: 20,
		},
		{
			myZodiac: "Овен", otherZodiac: "Телец",
			expectedCompatibility: 5,
		},
		{
			myZodiac: "Близнецы", otherZodiac: "Водолей",
			expectedCompatibility: 20,
		},
		{
			myZodiac: "Рак", otherZodiac: "Скорпион",
			expectedCompatibility: 20,
		},
		{
			myZodiac: "Дева", otherZodiac: "Весы",
			expectedCompatibility: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.myZodiac+"_"+test.otherZodiac, func(t *testing.T) {
			compatibility := GetZodiacCompatibility(test.myZodiac, test.otherZodiac)
			if compatibility != test.expectedCompatibility {
				t.Errorf("Ожидаемая совместимость %d, получено %d", test.expectedCompatibility, compatibility)
			}
		})
	}
}
