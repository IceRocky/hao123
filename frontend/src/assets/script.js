document.addEventListener("DOMContentLoaded", function () {
    const tabButtons = document.querySelectorAll(".tab-button");
    const tabContents = document.querySelectorAll(".tab-content");

    tabButtons.forEach((button) => {
      button.addEventListener("click", () => {
        // Reset all tabs
        tabButtons.forEach((btn) => {
          btn.querySelector("span").classList.add("hidden");
          btn.classList.remove("text-green-500");
        });
        tabContents.forEach((content) => content.classList.add("hidden"));

        // Activate clicked tab
        button.querySelector("span").classList.remove("hidden");
        button.classList.add("text-green-500");
        const tabId = button.getAttribute("data-tab");
        document.getElementById(tabId).classList.remove("hidden");
      });
    });

    // Activate the first tab by default
    tabButtons[0].querySelector("span").classList.remove("hidden");
    tabButtons[0].classList.add("text-green-500");
    tabContents[0].classList.remove("hidden");
  });

  document.addEventListener("DOMContentLoaded", function () {
    const tabButtons = document.querySelectorAll(".tab-btn");
    const tabPanels = document.querySelectorAll(".tab-panel");

    tabButtons.forEach((btn) => {
      btn.addEventListener("click", () => {
        // Reset all tabs and hide all panels
        tabButtons.forEach((button) => {
          button.querySelector(".active-indicator").classList.add("hidden");
        });
        tabPanels.forEach((panel) => panel.classList.add("hidden"));

        // Activate clicked tab and show the corresponding panel
        btn.querySelector(".active-indicator").classList.remove("hidden");
        const tabId = btn.getAttribute("data-tab");
        document.getElementById(tabId).classList.remove("hidden");
      });
    });

    // Activate the first tab by default
    tabButtons[0]
      .querySelector(".active-indicator")
      .classList.remove("hidden");
    tabPanels[0].classList.remove("hidden");
  });

  document.addEventListener("DOMContentLoaded", () => {
const showButton = document.getElementById("showButton");
const hideButton = document.getElementById("hideButton");
const contentDiv = document.getElementById("contentDiv");

// Show content
showButton.addEventListener("click", (e) => {
e.preventDefault();
contentDiv.classList.remove("hidden");
showButton.classList.add("hidden");
});

// Hide content
hideButton.addEventListener("click", (e) => {
e.preventDefault();
contentDiv.classList.add("hidden");
showButton.classList.remove("hidden");
});
});

document.querySelectorAll('.accordion-header').forEach(header => {
    header.addEventListener('click', () => {
        const content = header.nextElementSibling;
        document.querySelectorAll('.accordion-content').forEach(item => {
            if (item !== content) {
                item.classList.add('hidden');
            }
        });
        content.classList.toggle('hidden');
    });
});


document.addEventListener("DOMContentLoaded", function () {
    const tabButtons = document.querySelectorAll(".tab-button-2");
    const tabContents = document.querySelectorAll(".tab-content-2");

    tabButtons.forEach((button) => {
      button.addEventListener("click", () => {
        // Reset all tabs
        tabButtons.forEach((btn) => {
          btn.querySelector("span").classList.add("hidden");
          btn.classList.remove("text-green-500");
        });
        tabContents.forEach((content) => content.classList.add("hidden"));

        // Activate clicked tab
        button.querySelector("span").classList.remove("hidden");
        button.classList.add("text-green-500");
        const tabId = button.getAttribute("data-tab");
        document.getElementById(tabId).classList.remove("hidden");
      });
    });

    // Activate the first tab by default
    tabButtons[0].querySelector("span").classList.remove("hidden");
    tabButtons[0].classList.add("text-green-500");
    tabContents[0].classList.remove("hidden");
  });


  document.addEventListener("DOMContentLoaded", function () {
    const tabButtons = document.querySelectorAll(".tab-button-3");
    const tabContents = document.querySelectorAll(".tab-content-3");

    tabButtons.forEach((button) => {
      button.addEventListener("click", () => {
        // Reset all tabs
        tabButtons.forEach((btn) => {
          btn.querySelector("span").classList.add("hidden");
          btn.classList.remove("text-green-500");
        });
        tabContents.forEach((content) => content.classList.add("hidden"));

        // Activate clicked tab
        button.querySelector("span").classList.remove("hidden");
        button.classList.add("text-green-500");
        const tabId = button.getAttribute("data-tab");
        document.getElementById(tabId).classList.remove("hidden");
      });
    });

    // Activate the first tab by default
    tabButtons[0].querySelector("span").classList.remove("hidden");
    tabButtons[0].classList.add("text-green-500");
    tabContents[0].classList.remove("hidden");
  });


const slider = document.getElementById('slider');
const slides = slider.children;
let currentIndex = 0;

document.getElementById('prevBtn').addEventListener('click', () => {
  currentIndex = (currentIndex - 1 + slides.length) % slides.length;
  slider.style.transform = `translateX(-${currentIndex * 100}%)`;
});

document.getElementById('nextBtn').addEventListener('click', () => {
  currentIndex = (currentIndex + 1) % slides.length;
  slider.style.transform = `translateX(-${currentIndex * 100}%)`;
});


// document.addEventListener("DOMContentLoaded", function () {
//   const dropdownButton = document.getElementById("dropdown-button");
//   const dropdown = document.getElementById("dropdown");

//   dropdownButton.addEventListener("click", function () {
//     dropdown.classList.toggle("hidden");
//   });
// });
