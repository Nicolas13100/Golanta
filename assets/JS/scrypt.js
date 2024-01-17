document.addEventListener("DOMContentLoaded", function () {
  const totalPointsInput = document.getElementById("totalPoints");
  const statSliders = document.querySelectorAll(".stat-slider");
  const statInputs = document.querySelectorAll(".stat-input");

  statSliders.forEach(function (slider, index) {
    slider.addEventListener("input", function () {
      updateTotalPoints();
      updateStatInputValue(index);
      limitTotalPoints(); // Add this function call to limit total points
    });
  });

  function updateTotalPoints() {
    let totalPoints = 0;

    statSliders.forEach(function (slider) {
      totalPoints += parseInt(slider.value) || 0;
    });

    totalPointsInput.value = totalPoints;
  }

  function updateStatInputValue(index) {
    statInputs[index].value = statSliders[index].value;
  }

  function limitTotalPoints() {
    const maxTotalPoints = 400;

    if (parseInt(totalPointsInput.value) > maxTotalPoints) {
      // Block all sliders if total points exceed the maximum
      statSliders.forEach(function (slider) {
        slider.value = 0; // Set each slider's value to 0
        updateStatInputValue(Array.from(statSliders).indexOf(slider));
      });

      updateTotalPoints();

      // Disable scrolling when total points exceed the maximum
      document.body.style.overflowY = "hidden";
    } else {
      // Enable scrolling otherwise
      document.body.style.overflowY = "auto";
    }
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