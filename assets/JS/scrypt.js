document.addEventListener("DOMContentLoaded", function () {
  const totalPointsInput = document.getElementById("totalPoints");
  const statSliders = document.querySelectorAll(".stat-slider");
  const statInputs = document.querySelectorAll(".stat-input");

  statSliders.forEach(function (slider, index) {
    slider.addEventListener("input", function () {
      updateStatInputValue(index);
      updateTotalPoints();
    });
  });

  function updateTotalPoints() {
    let totalPoints = 0;

    // Calculate the total points considering each slider
    statSliders.forEach(function (slider) {
      totalPoints += parseInt(slider.value) || 0;
    });

    // Check if total points exceed the limit
    if (totalPoints > 400) {
      alert("Total points cannot exceed 400. Please adjust your allocations.");

      // Determine the excess points
      const excessPoints = totalPoints - 400;

      // Distribute the excess points proportionally among all sliders
      statSliders.forEach(function (slider, index) {
        const currentValue = parseInt(slider.value) || 0;
        const newSliderValue = Math.max(currentValue - (excessPoints / statSliders.length), 0);
        slider.value = newSliderValue.toFixed(2); // Ensure two decimal places
      });
    }

    // Update the total points field
    totalPointsInput.value = totalPoints;
  }

  function updateStatInputValue(index) {
    statInputs[index].value = statSliders[index].value;
  }
});

document.addEventListener('DOMContentLoaded', function () {
  document.getElementById('PersosImage').addEventListener('change', function (e) {
    const previewImage = document.getElementById('previewImage');
    const firstImage = document.getElementById('firstImage'); // Replace 'firstImage' with the actual ID of the first image
    const fileInput = e.target;

    if (fileInput.files && fileInput.files[0]) {
      const reader = new FileReader();

      reader.onload = function (e) {
        previewImage.src = e.target.result;
        previewImage.style.display = 'flex';

        // Hide the first image
        if (firstImage) {
          firstImage.style.display = 'none';
        }
      };

      reader.readAsDataURL(fileInput.files[0]);
    } else {
      // If no file is selected, hide the preview image and show the first image
      previewImage.style.display = 'none';
      if (firstImage) {
        firstImage.style.display = 'flex';
      }
      // Adjust 'flex' to your desired display property
    }
  });
});