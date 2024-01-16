document.addEventListener("DOMContentLoaded", function () {
    // Endurence
    var enduranceValue = parseFloat(document.getElementById('enduranceValue').innerText);
    var enduranceProgressBar = document.getElementById('enduranceProgress');
    if (enduranceProgressBar) {
        enduranceProgressBar.style.width = (2 * enduranceValue) + '%';
    }
    // Stamina
    var staminaValue = parseFloat(document.getElementById('staminaValue').innerText);
    var staminaProgressBar = document.getElementById('staminaProgress');
    if (staminaProgressBar) {
        staminaProgressBar.style.width = (2 * staminaValue) + '%';
    }

    // Physical Agility
    var agilityValue = parseFloat(document.getElementById('agilityValue').innerText);
    var agilityProgressBar = document.getElementById('agilityProgress');
    if (agilityProgressBar) {
        agilityProgressBar.style.width = (2 * agilityValue) + '%';
    }
    // Shelter Building
    var shelterBuildingValue = parseFloat(document.getElementById('shelterBuildingValue').innerText);
    var shelterBuildingProgressBar = document.getElementById('shelterBuildingProgress');
    if (shelterBuildingProgressBar) {
        shelterBuildingProgressBar.style.width = (2 * shelterBuildingValue) + '%';
    }

    // Fire Making
    var fireMakingValue = parseFloat(document.getElementById('fireMakingValue').innerText);
    var fireMakingProgressBar = document.getElementById('fireMakingProgress');
    if (fireMakingProgressBar) {
        fireMakingProgressBar.style.width = (2 * fireMakingValue) + '%';
    }
    // Strategic Thinking
    var strategicThinkingValue = parseFloat(document.getElementById('strategicThinkingValue').innerText);
    var strategicThinkingProgressBar = document.getElementById('strategicThinkingProgress');
    if (strategicThinkingProgressBar) {
        strategicThinkingProgressBar.style.width = (2 * strategicThinkingValue) + '%';
    }

    // Manipulation
    var manipulationValue = parseFloat(document.getElementById('manipulationValue').innerText);
    var manipulationProgressBar = document.getElementById('manipulationProgress');
    if (manipulationProgressBar) {
        manipulationProgressBar.style.width = (2 * manipulationValue) + '%';
    }

    // Mental Endurance
    var mentalEnduranceValue = parseFloat(document.getElementById('mentalEnduranceValue').innerText);
    var mentalEnduranceProgressBar = document.getElementById('mentalEnduranceProgress');
    if (mentalEnduranceProgressBar) {
        mentalEnduranceProgressBar.style.width = (2 * mentalEnduranceValue) + '%';
    }

    // Team Player
    var teamPlayerValue = parseFloat(document.getElementById('teamPlayerValue').innerText);
    var teamPlayerProgressBar = document.getElementById('teamPlayerProgress');
    if (teamPlayerProgressBar) {
        teamPlayerProgressBar.style.width = (2 * teamPlayerValue) + '%';
    }
    // Leadership Skills
    var leadershipSkillsValue = parseFloat(document.getElementById('leadershipSkillsValue').innerText);
    var leadershipSkillsProgressBar = document.getElementById('leadershipSkillsProgress');
    if (leadershipSkillsProgressBar) {
        leadershipSkillsProgressBar.style.width = (2 * leadershipSkillsValue) + '%';
    }

    // Individual Challenge Performance
    var IndividualChallengePerformanceValue = parseFloat(document.getElementById('IndividualChallengePerformanceValue').innerText);
    var IndividualChallengePerformanceProgressBar = document.getElementById('IndividualChallengePerformanceProgress');
    if (IndividualChallengePerformanceProgressBar) {
        IndividualChallengePerformanceProgressBar.style.width = (2 * IndividualChallengePerformanceValue) + '%';
    }

    // Team Challenge Contribution
    var TeamChallengeContributionValue = parseFloat(document.getElementById('TeamChallengeContributionValue').innerText);
    var TeamChallengeContributionProgressBar = document.getElementById('TeamChallengeContributionProgress');
    if (TeamChallengeContributionProgressBar) {
        TeamChallengeContributionProgressBar.style.width = (2 * TeamChallengeContributionValue) + '%';
    }

    // Adaptability
    var adaptabilityValue = parseFloat(document.getElementById('adaptabilityValue').innerText);
    var adaptabilityProgressBar = document.getElementById('adaptabilityProgress');
    if (adaptabilityProgressBar) {
        adaptabilityProgressBar.style.width = (2 * adaptabilityValue) + '%';
    }
});

function confirmDelete() {
        var confirmation = confirm("Êtes-vous sûr de vouloir supprimer ce personnage?");

        if (confirmation) {
            window.location.href = "/DeletChar";
        }
    }