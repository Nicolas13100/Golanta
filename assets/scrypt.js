document.addEventListener("DOMContentLoaded", function () {
    const totalPointsInput = document.getElementById("totalPoints");
    const statInputs = document.querySelectorAll(".stat-input");
  
    statInputs.forEach(function (input) {
      input.addEventListener("input", function () {
        updateTotalPoints();
      });
    });
  
    function updateTotalPoints() {
      let totalPoints = 0;
  
      statInputs.forEach(function (input) {
        totalPoints += parseInt(input.value) || 0;
      });
  
      totalPointsInput.value = totalPoints;
    }
  });