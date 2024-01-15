package routeur

import (
	"encoding/json"
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

func AddCharacterToFile(newCharacter Character, filename string) error {
	// Read existing data
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal JSON data into a slice of Character
	var characters []Character
	err = json.Unmarshal(data, &characters)
	if err != nil {
		return err
	}

	// Add the new character
	characters = append(characters, newCharacter)

	// Marshal the updated data
	updatedData, err := json.MarshalIndent(characters, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated data back to the file
	err = os.WriteFile(filename, updatedData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Character added successfully!")

	return nil
}
