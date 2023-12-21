document.addEventListener("DOMContentLoaded", function () {
  // Select all links with the class 'navbar-link'
  const links = document.querySelectorAll(".navbar-link");

  // Get the current page's full URL
  const currentUrl = window.location.href;

  links.forEach((link) => {
    // Check if the link's href matches the current URL
    if (link.href === currentUrl) {
      // Add classes to highlight the active link
      link.classList.remove("text-gray-300");
      link.classList.remove("hover:bg-gray-700");
      link.classList.remove("hover:text-white");

      link.classList.add("bg-gray-900");
      link.classList.add("text-white");
    } else {
      // Remove classes highlighting the previous link
      link.classList.remove("bg-gray-900");
      link.classList.remove("text-white");

      link.classList.add("text-gray-300");
      link.classList.add("hover:bg-gray-700");
      link.classList.add("hover:text-white");
    }
  });

  const navbarOpenMenuButton = document.getElementById(
    "navbar-open-menu-button"
  );

  navbarOpenMenuButton.addEventListener("click", (event) => {
    const button = event.target;
    if (button.getAttribute("data-opened") === "true") {
      document.getElementById("navbar-icon-closed").classList.add("block");
      document.getElementById("navbar-icon-closed").classList.remove("hidden");

      document.getElementById("navbar-icon-opened").classList.remove("block");
      document.getElementById("navbar-icon-opened").classList.add("hidden");

      document.getElementById("navbar-mobile-menu").classList.add("hidden");
      document.getElementById("navbar-mobile-menu").classList.remove("block");

      button.setAttribute("data-opened", "false");
    } else {
      document.getElementById("navbar-icon-opened").classList.add("block");
      document.getElementById("navbar-icon-opened").classList.remove("hidden");

      document.getElementById("navbar-icon-closed").classList.remove("block");
      document.getElementById("navbar-icon-closed").classList.add("hidden");

      document.getElementById("navbar-mobile-menu").classList.remove("hidden");
      document.getElementById("navbar-mobile-menu").classList.add("block");

      button.setAttribute("data-opened", "true");
    }
  });
});
