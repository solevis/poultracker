const formatDate = (date) => {
  let dd = date.getDate();
  let mm = date.getMonth() + 1;
  let yyyy = date.getFullYear();

  if (dd < 10) {
    dd = "0" + dd;
  }

  if (mm < 10) {
    mm = "0" + mm;
  }

  return `${yyyy}-${mm}-${dd}`;
};

const resetError = () => {
  const errorElement = document.getElementById("error-panel");
  errorElement.innerText = "";
  errorElement.classList.add("hidden");
};

const showError = (message) => {
  const errorElement = document.getElementById("error-panel");
  errorElement.innerText = message;
  errorElement.classList.remove("hidden");
};

const resetPastError = () => {
  const errorElement = document.getElementById("past-error-panel");
  errorElement.innerText = "";
  errorElement.classList.add("hidden");
};

const showPastError = (message) => {
  const errorElement = document.getElementById("past-error-panel");
  errorElement.innerText = message;
  errorElement.classList.remove("hidden");
};

function openModal(el) {
  el.classList.remove("hidden");
}

function closeModal(el) {
  el.classList.add("hidden");
}
