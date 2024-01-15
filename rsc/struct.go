package routeur

type Character struct {
	PersosName                           string `json:"PersosName"`
	PersosImage                          string `json:"PersosImage"`
	PersosFullName                       string `json:"PersosFullName"`
	PersosDescription                    string `json:"PersosDescription"`
	PersosEquipe                         string `json:"PersosEquipe"`
	PersosPersonalite                    string `json:"PersosPersonalite"`
	PersosApparence                      string `json:"PersosApparence"`
	PersosCapacites                      string `json:"PersosCapacités"`
	PersosHistoires                      string `json:"PersosHistoires"`
	PersosEndurance                      int    `json:"PersosEndurance"`
	PersosStamina                        int    `json:"PersosStamina"`
	PersosPhysicalAgility                int    `json:"PersosPhysicalAgility"`
	PersosShelterBuilding                int    `json:"PersosShelterBuilding"`
	PersosFireMaking                     int    `json:"PersosFireMaking"`
	PersosStrategicThinking              int    `json:"PersosStrategicThinking"`
	PersosManipulation                   int    `json:"PersosManipulation"`
	PersosMentalEndurance                int    `json:"PersosMentalEndurance"`
	PersosTeamPlayer                     int    `json:"PersosTeamPlayer"`
	PersosLeadershipSkills               int    `json:"PersosLeadershipSkills"`
	PersosIndividualChallengePerformance int    `json:"PersosIndividualChallengePerformance"`
	PersosTeamChallengeContribution      int    `json:"PersosTeamChallengeContribution"`
	PersosAdaptability                   int    `json:"PersosAdaptability"`
}
