package routeur

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func RUN() {
	// Set up your other handlers
	http.HandleFunc("/", errorHandler)
	http.HandleFunc("/Home", indexHandler)
	http.HandleFunc("/CreatChar", CreatHandler)
	http.HandleFunc("/CreatChar/Gestion", CreaGestionHandler)
	http.HandleFunc("/ModifyChar", ModifyHandler)
	http.HandleFunc("/ModifyChar/Gestion", ModifyGestionHandler)
	http.HandleFunc("/DeletChar", DeleteHandler)
	http.HandleFunc("/CharList", ListHandler)
	http.HandleFunc("/CharDisplay", CharDisplayHandler)
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	// Print statement indicating server is running
	fmt.Println("Server is running on :8080 http://localhost:8080/Home")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "404", nil)
}

func CharDisplayHandler(w http.ResponseWriter, r *http.Request) {
	// Read data from data.json
	data, err := os.ReadFile("data.json")
	if err != nil {
		http.Error(w, "Error reading data.json", http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON data into a slice of characters
	var characters []Character
	if err := json.Unmarshal(data, &characters); err != nil {
		http.Error(w, "Error parsing data.json", http.StatusInternalServerError)
		return
	}

	// Extract the "fullname" parameter from the URL query
	fullname := r.URL.Query().Get("fullname")

	// Find the character with the specified fullname in the characters data
	var selectedCharacter Character
	for _, char := range characters {
		if char.PersosFullName == fullname {
			selectedCharacter = char
			break
		}
	}

	// Check if the character was found
	if selectedCharacter.PersosFullName == "" {
		// Handle the case where the character is not found (e.g., show an error message)
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	// Pass the selected character to the renderTemplate function
	renderTemplate(w, "selectedChar", selectedCharacter)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func CreatHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "newChar", nil)
}

func CreaGestionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	// Convert string values to integers
	endurance, err := strconv.Atoi(r.FormValue("PersosEndurance"))
	if err != nil {
		http.Error(w, "Error converting PersosEndurance to int", http.StatusBadRequest)
		return
	}

	stamina, err := strconv.Atoi(r.FormValue("PersosStamina"))
	if err != nil {
		http.Error(w, "Error converting PersosStamina to int", http.StatusBadRequest)
		return
	}
	PhysicalAgility, err := strconv.Atoi(r.FormValue("PersosPhysicalAgility"))
	if err != nil {
		http.Error(w, "Error converting PersosPhysicalAgility to int", http.StatusBadRequest)
		return
	}
	ShelterBuilding, err := strconv.Atoi(r.FormValue("PersosShelterBuilding"))
	if err != nil {
		http.Error(w, "Error converting PersosShelterBuilding to int", http.StatusBadRequest)
		return
	}
	FireMaking, err := strconv.Atoi(r.FormValue("PersosFireMaking"))
	if err != nil {
		http.Error(w, "Error converting PersosFireMaking to int", http.StatusBadRequest)
		return
	}

	StrategicThinkin, err := strconv.Atoi(r.FormValue("PersosStrategicThinking"))
	if err != nil {
		http.Error(w, "Error converting PersosStrategicThinking to int", http.StatusBadRequest)
		return
	}
	Manipulation, err := strconv.Atoi(r.FormValue("PersosManipulation"))
	if err != nil {
		http.Error(w, "Error converting PersosManipulation to int", http.StatusBadRequest)
		return
	}

	MentalEndurance, err := strconv.Atoi(r.FormValue("PersosMentalEndurance"))
	if err != nil {
		http.Error(w, "Error converting PersosMentalEndurance to int", http.StatusBadRequest)
		return
	}
	TeamPlayer, err := strconv.Atoi(r.FormValue("PersosTeamPlayer"))
	if err != nil {
		http.Error(w, "Error converting PersosTeamPlayer to int", http.StatusBadRequest)
		return
	}

	LeadershipSkills, err := strconv.Atoi(r.FormValue("PersosLeadershipSkills"))
	if err != nil {
		http.Error(w, "Error converting PersosLeadershipSkills to int", http.StatusBadRequest)
		return
	}
	IndividualChallengePerformance, err := strconv.Atoi(r.FormValue("PersosIndividualChallengePerformance"))
	if err != nil {
		http.Error(w, "Error converting PersosIndividualChallengePerformance to int", http.StatusBadRequest)
		return
	}
	TeamChallengeContribution, err := strconv.Atoi(r.FormValue("PersosTeamChallengeContribution"))
	if err != nil {
		http.Error(w, "Error converting PersosTeamChallengeContribution to int", http.StatusBadRequest)
		return
	}
	Adaptability, err := strconv.Atoi(r.FormValue("PersosAdaptability"))
	if err != nil {
		http.Error(w, "Error converting PersosAdaptability to int", http.StatusBadRequest)
		return
	}

	message, err := validateTotalPoints(endurance, stamina, PhysicalAgility, ShelterBuilding, FireMaking, StrategicThinkin, Manipulation, MentalEndurance, TeamPlayer, LeadershipSkills, IndividualChallengePerformance, TeamChallengeContribution, Adaptability)
	if err != nil {
		fmt.Println("Error:", err)
		renderTemplate(w, "newChar", struct{ ErrorMessage string }{ErrorMessage: message})
		return
	}

	imageFile, imageHeader, err := r.FormFile("PersosImage")
	if err != nil {
		http.Error(w, "Error retrieving PersosImage", http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	// Get the file extension
	ext := filepath.Ext(imageHeader.Filename)

	// Create the path for the image file in the /assets/IMG/ directory
	imageName := fmt.Sprintf("%s%s", r.FormValue("PersosFullName"), ext)
	imagePath := filepath.Join("assets", "IMG", "Imgpersos", imageName)
	imageSavePath := "/static/IMG/ImgPersos/" + imageName
	// Create a new file at the specified path
	outputFile, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, "Error creating image file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Copy the contents of the uploaded file to the new file
	_, err = io.Copy(outputFile, imageFile)
	if err != nil {
		http.Error(w, "Error copying image file", http.StatusInternalServerError)
		return
	}

	newChar := Character{
		PersosName:                           r.FormValue("PersosName"),
		PersosImage:                          imageSavePath,
		PersosFullName:                       r.FormValue("PersosFullName"),
		PersosDescription:                    r.FormValue("PersosDescription"),
		PersosEquipe:                         r.FormValue("PersosEquipe"),
		PersosPersonalite:                    r.FormValue("PersosPersonalite"),
		PersosApparence:                      r.FormValue("PersosApparence"),
		PersosCapacites:                      r.FormValue("PersosCapacité"),
		PersosHistoires:                      r.FormValue("PersosHistoire"),
		PersosEndurance:                      endurance,
		PersosStamina:                        stamina,
		PersosPhysicalAgility:                PhysicalAgility,
		PersosShelterBuilding:                ShelterBuilding,
		PersosFireMaking:                     FireMaking,
		PersosStrategicThinking:              StrategicThinkin,
		PersosManipulation:                   Manipulation,
		PersosMentalEndurance:                MentalEndurance,
		PersosTeamPlayer:                     TeamPlayer,
		PersosLeadershipSkills:               LeadershipSkills,
		PersosIndividualChallengePerformance: IndividualChallengePerformance,
		PersosTeamChallengeContribution:      TeamChallengeContribution,
		PersosAdaptability:                   Adaptability,
	}

	// Provide the path to your data.json file
	filename := "data.json"

	message, err = AddCharacterToFile(newChar, filename)
	if err != nil {
		fmt.Println("Error:", err)
		renderTemplate(w, "newChar", struct{ ErrorMessage string }{ErrorMessage: message})
		return
	}

	http.Redirect(w, r, "/CharList", http.StatusSeeOther)
}

func ModifyHandler(w http.ResponseWriter, r *http.Request) {
	// Read data from data.json
	data, err := os.ReadFile("data.json")
	if err != nil {
		http.Error(w, "Error reading data.json", http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON data into a slice of characters
	var characters []Character
	if err := json.Unmarshal(data, &characters); err != nil {
		http.Error(w, "Error parsing data.json", http.StatusInternalServerError)
		return
	}

	// Extract the "fullname" parameter from the URL query
	fullname := r.URL.Query().Get("fullname")

	// Find the character with the specified fullname in the characters data
	var selectedCharacter Character
	for _, char := range characters {
		if char.PersosFullName == fullname {
			selectedCharacter = char
			break
		}
	}

	// Check if the character was found
	if selectedCharacter.PersosFullName == "" {
		// Handle the case where the character is not found (e.g., show an error message)
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}
	renderTemplate(w, "modifyChar", selectedCharacter)
}

func ModifyGestionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	// Convert string values to integers
	endurance, err := strconv.Atoi(r.FormValue("PersosEndurance"))
	if err != nil {
		http.Error(w, "Error converting PersosEndurance to int", http.StatusBadRequest)
		return
	}

	stamina, err := strconv.Atoi(r.FormValue("PersosStamina"))
	if err != nil {
		http.Error(w, "Error converting PersosStamina to int", http.StatusBadRequest)
		return
	}
	PhysicalAgility, err := strconv.Atoi(r.FormValue("PersosPhysicalAgility"))
	if err != nil {
		http.Error(w, "Error converting PersosPhysicalAgility to int", http.StatusBadRequest)
		return
	}
	ShelterBuilding, err := strconv.Atoi(r.FormValue("PersosShelterBuilding"))
	if err != nil {
		http.Error(w, "Error converting PersosShelterBuilding to int", http.StatusBadRequest)
		return
	}
	FireMaking, err := strconv.Atoi(r.FormValue("PersosFireMaking"))
	if err != nil {
		http.Error(w, "Error converting PersosFireMaking to int", http.StatusBadRequest)
		return
	}

	StrategicThinkin, err := strconv.Atoi(r.FormValue("PersosStrategicThinking"))
	if err != nil {
		http.Error(w, "Error converting PersosStrategicThinking to int", http.StatusBadRequest)
		return
	}
	Manipulation, err := strconv.Atoi(r.FormValue("PersosManipulation"))
	if err != nil {
		http.Error(w, "Error converting PersosManipulation to int", http.StatusBadRequest)
		return
	}

	MentalEndurance, err := strconv.Atoi(r.FormValue("PersosMentalEndurance"))
	if err != nil {
		http.Error(w, "Error converting PersosMentalEndurance to int", http.StatusBadRequest)
		return
	}
	TeamPlayer, err := strconv.Atoi(r.FormValue("PersosTeamPlayer"))
	if err != nil {
		http.Error(w, "Error converting PersosTeamPlayer to int", http.StatusBadRequest)
		return
	}

	LeadershipSkills, err := strconv.Atoi(r.FormValue("PersosLeadershipSkills"))
	if err != nil {
		http.Error(w, "Error converting PersosLeadershipSkills to int", http.StatusBadRequest)
		return
	}
	IndividualChallengePerformance, err := strconv.Atoi(r.FormValue("PersosIndividualChallengePerformance"))
	if err != nil {
		http.Error(w, "Error converting PersosIndividualChallengePerformance to int", http.StatusBadRequest)
		return
	}
	TeamChallengeContribution, err := strconv.Atoi(r.FormValue("PersosTeamChallengeContribution"))
	if err != nil {
		http.Error(w, "Error converting PersosTeamChallengeContribution to int", http.StatusBadRequest)
		return
	}
	Adaptability, err := strconv.Atoi(r.FormValue("PersosAdaptability"))
	if err != nil {
		http.Error(w, "Error converting PersosAdaptability to int", http.StatusBadRequest)
		return
	}

	message, err := validateTotalPoints(endurance, stamina, PhysicalAgility, ShelterBuilding, FireMaking, StrategicThinkin, Manipulation, MentalEndurance, TeamPlayer, LeadershipSkills, IndividualChallengePerformance, TeamChallengeContribution, Adaptability)
	if err != nil {
		fmt.Println("Error:", err)
		renderTemplate(w, "newChar", struct{ ErrorMessage string }{ErrorMessage: message})
		return
	}

	imageFile, imageHeader, err := r.FormFile("PersosImage")
	if err != nil {
		http.Error(w, "Error retrieving PersosImage", http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	// Get the file extension
	ext := filepath.Ext(imageHeader.Filename)

	// Create the path for the image file in the /assets/IMG/ directory
	imageName := fmt.Sprintf("%s%s", r.FormValue("PersosFullName"), ext)
	imagePath := filepath.Join("assets", "IMG", "ImgPersos", imageName)
	imageSavePath := "/static/IMG/ImgPersos/" + imageName
	// Create a new file at the specified path
	outputFile, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, "Error creating image file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Copy the contents of the uploaded file to the new file
	_, err = io.Copy(outputFile, imageFile)
	if err != nil {
		http.Error(w, "Error copying image file", http.StatusInternalServerError)
		return
	}

	modifiedChar := Character{
		PersosName:                           r.FormValue("PersosName"),
		PersosImage:                          imageSavePath,
		PersosFullName:                       r.FormValue("PersosFullName"),
		PersosDescription:                    r.FormValue("PersosDescription"),
		PersosEquipe:                         r.FormValue("PersosEquipe"),
		PersosPersonalite:                    r.FormValue("PersosPersonalite"),
		PersosApparence:                      r.FormValue("PersosApparence"),
		PersosCapacites:                      r.FormValue("PersosCapacité"),
		PersosHistoires:                      r.FormValue("PersosHistoire"),
		PersosEndurance:                      endurance,
		PersosStamina:                        stamina,
		PersosPhysicalAgility:                PhysicalAgility,
		PersosShelterBuilding:                ShelterBuilding,
		PersosFireMaking:                     FireMaking,
		PersosStrategicThinking:              StrategicThinkin,
		PersosManipulation:                   Manipulation,
		PersosMentalEndurance:                MentalEndurance,
		PersosTeamPlayer:                     TeamPlayer,
		PersosLeadershipSkills:               LeadershipSkills,
		PersosIndividualChallengePerformance: IndividualChallengePerformance,
		PersosTeamChallengeContribution:      TeamChallengeContribution,
		PersosAdaptability:                   Adaptability,
	}

	// Provide the path to your data.json file
	filename := "data.json"
	fullname := r.URL.Query().Get("fullname")
	UpdateCharData(filename, fullname, modifiedChar)
	http.Redirect(w, r, "/CharList", http.StatusSeeOther)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Read data from data.json
	data, err := os.ReadFile("data.json")
	if err != nil {
		http.Error(w, "Error reading data.json", http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON data into a slice of characters
	var characters []Character
	if err := json.Unmarshal(data, &characters); err != nil {
		http.Error(w, "Error parsing data.json", http.StatusInternalServerError)
		return
	}

	// Extract the "fullname" parameter from the URL query
	fullname := r.URL.Query().Get("fullname")

	// Find the index of the character with the specified fullname in the characters slice
	var selectedIndex = -1
	for i, char := range characters {
		if char.PersosFullName == fullname {
			selectedIndex = i
			break
		}
	}

	// Check if the character was found
	if selectedIndex == -1 {
		// Handle the case where the character is not found (e.g., show an error message)
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	// Remove the selected character from the slice
	characters = append(characters[:selectedIndex], characters[selectedIndex+1:]...)

	// Marshal the modified characters slice back to JSON
	updatedData, err := json.Marshal(characters)
	if err != nil {
		http.Error(w, "Error marshaling updated data", http.StatusInternalServerError)
		return
	}

	// Write the updated JSON back to data.json
	err = os.WriteFile("data.json", updatedData, 0644)
	if err != nil {
		http.Error(w, "Error writing updated data to data.json", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/CharList", http.StatusSeeOther)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the JSON data
	var characters []Character
	err = json.NewDecoder(file).Decode(&characters)
	if err != nil {
		log.Fatal(err)
	}

	renderTemplate(w, "listChar", characters)
}
