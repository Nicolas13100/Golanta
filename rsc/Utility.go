package routeur

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, err := template.New(tmplName).ParseFiles("Template/" + tmplName + ".html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AddCharacterToFile(newCharacter Character, filename string) (string, error) {
	// Read existing data
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Unmarshal JSON data into a slice of Character
	var characters []Character
	err = json.Unmarshal(data, &characters)
	if err != nil {
		return "", err
	}

	// Check if there are already 5 characters
	if len(characters) >= 5 {
		return "Maximum number of characters reached", errors.New("maximum number of characters reached")
	}

	// Check if PersosFullName is already given
	for _, existingCharacter := range characters {
		if existingCharacter.PersosFullName == newCharacter.PersosFullName {
			return "PersosFullName already exists", errors.New("PersosFullName already exists")
		}
	}

	// Add the new character
	characters = append(characters, newCharacter)

	// Marshal the updated data
	updatedData, err := json.MarshalIndent(characters, "", "  ")
	if err != nil {
		return "", err
	}

	// Write the updated data back to the file
	err = os.WriteFile(filename, updatedData, 0644)
	if err != nil {
		return "", err
	}

	fmt.Println("Character added successfully!")

	return "Character added successfully!", nil
}

func validateTotalPoints(endurance, stamina, physicalAgility, shelterBuilding, fireMaking, strategicThinking, manipulation, mentalEndurance, teamPlayer, leadershipSkills, individualChallengePerformance, teamChallengeContribution, adaptability int) (string, error) {
	totalPoints := endurance + stamina + physicalAgility + shelterBuilding + fireMaking + strategicThinking + manipulation + mentalEndurance + teamPlayer + leadershipSkills + individualChallengePerformance + teamChallengeContribution + adaptability

	if totalPoints > 400 {
		return "Total points for integer stats should not exceed 400", errors.New("total points for integer stats should not exceed 400")
	}

	return "", nil
}
