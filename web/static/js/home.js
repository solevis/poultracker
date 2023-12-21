let COLLECTIONS_HISTORY_RAW_DATA = {};

const onDeleteButtonClick = (event) => {
  if (!confirm("Do you really want to do away with the eggs of the day?")) {
    return;
  }

  resetError();

  const element = event.target;
  const id = parseInt(element.dataset.id);

  fetch(`/api/collections/${id}`, {
    method: "DELETE",
  })
    .then((response) => {
      if (!response.ok) {
        return Promise.reject(response);
      }

      const collectionsNumberInput = document.getElementById(
        "collections-number-input"
      );
      collectionsNumberInput.value = null;

      const deleteButton = document.getElementById("delete-button");
      deleteButton.classList.add("hidden");

      element.dataset.id = "0";

      const saveButton = document.getElementById("save-button");
      saveButton.dataset.id = "0";
      saveButton.innerText = "Save";

      // update table
      loadCollectionsHistoryTable();
    })
    .catch(() => showError("An error has occurred."));
};

const onSaveButtonClick = (event) => {
  const element = event.target;
  const id = parseInt(element.dataset.id);

  resetError();

  const collectionsNumberInput = document.getElementById(
    "collections-number-input"
  );
  const todayCollectionsNumber = Number(collectionsNumberInput.value);

  const isInteger = Number.isInteger(todayCollectionsNumber);
  const isPositive = todayCollectionsNumber > 0;
  const isPositiveInteger = isInteger && isPositive;

  if (!isPositiveInteger) {
    showError("Please enter a valid number of eggs.");
    return;
  }

  if (id === 0) {
    let laidDate = formatDate(new Date());
    requestPost(laidDate, todayCollectionsNumber, showError, null);
  } else {
    let laidDate = formatDate(new Date());
    requestPut(laidDate, id, todayCollectionsNumber, showError, null);
  }
};

const onPastSaveButtonClick = (event) => {
  resetPastError();

  const collectionsNumberInput = document.getElementById(
    "past-collections-number-input"
  );
  const pastCollectionsNumber = Number(collectionsNumberInput.value);

  let isInteger = Number.isInteger(pastCollectionsNumber);
  let isPositive = pastCollectionsNumber > 0;
  let isPositiveInteger = isInteger && isPositive;

  if (!isPositiveInteger) {
    showPastError("Please enter a valid number of eggs.");
    return;
  }

  const collectionsDateInput = document.getElementById(
    "past-collections-date-input"
  );
  let laidDate = collectionsDateInput.value;
  if (!laidDate) {
    showPastError("Please enter a valid date.");
    return;
  }

  requestPost(
    laidDate,
    pastCollectionsNumber,
    showPastError,
    closePastCollectionsModal
  );
};

const closePastCollectionsModal = () => {
  const pastAddModal = document.getElementById("past-add-modal");
  closeModal(pastAddModal);
};

const loadCollectionsHistoryTable = () => {
  fetch("/api/collections", {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
  })
    .then((response) => {
      if (!response.ok) {
        return Promise.reject(response);
      }

      return response.json();
    })
    .then((json) => (COLLECTIONS_HISTORY_RAW_DATA = json))
    .then(() => {
      const timeline = document.getElementById("timeline");

      // wipe previous items
      timeline.innerHTML = "";

      // no data, hide the panel
      const historyPanel = document.getElementById("history-panel");
      const navChartItems = document.querySelectorAll(".navbar-link-chart");

      if (!COLLECTIONS_HISTORY_RAW_DATA) {
        historyPanel.classList.add("hidden");
        navChartItems.forEach((link) => {
          link.classList.add("hidden");
        });

        return;
      } else {
        historyPanel.classList.remove("hidden");
        navChartItems.forEach((link) => {
          link.classList.remove("hidden");
        });
      }

      // add items
      COLLECTIONS_HISTORY_RAW_DATA.slice()
        .reverse()
        .forEach((collection) => {
          let timelineItem = createTimelineItemElement(
            collection.laidDate,
            collection.number
          );

          timeline.appendChild(timelineItem);
        });
    })
    .catch(() => showError("An error has occurred."));
};

const createTimelineItemElement = (laidDate, number) => {
  const timelineItem = document.createElement("div");

  timelineItem.classList.add(
    "flex",
    "items-center",
    "w-full",
    "my-6",
    "-ml-1.5"
  );

  timelineItem.innerHTML = `
    <div class="flex items-center w-full my-6 -ml-1.5">
      <div class="w-1/12 z-10">
        <div class="w-3.5 h-3.5 bg-gray-800 rounded-full"></div>
      </div>
      <div class="w-11/12">
        <p class="text-sm">${number} eggs.</p>
        <p class="text-xs text-gray-500">${laidDate}</p>
      </div>
    </div>
  `;

  return timelineItem;
};

const requestPost = (
  laidDate,
  collectionsNumber,
  showErrorFunction,
  successCallback
) => {
  fetch("/api/collections", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      laidDate,
      number: collectionsNumber,
    }),
  })
    .then((response) => {
      if (!response.ok) {
        return Promise.reject(response);
      }

      return response.json();
    })
    .then((json) => {
      const deleteButton = document.getElementById("delete-button");
      deleteButton.classList.remove("hidden");
      deleteButton.dataset.id = json.id;

      const saveButton = document.getElementById("save-button");
      saveButton.dataset.id = json.id;
      saveButton.innerText = "Edit";

      // update table
      loadCollectionsHistoryTable();

      if (successCallback) {
        successCallback();
      }
    })
    .catch((response) => {
      console.log(response.status, response.statusText);
      showErrorFunction("An error has occurred.");
    });
};

const requestPut = (
  laidDate,
  id,
  collectionsNumber,
  showErrorFunction,
  successCallback
) => {
  fetch(`/api/collections`, {
    method: "PUT",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id,
      laidDate,
      number: collectionsNumber,
    }),
  })
    .then((response) => {
      if (!response.ok) {
        return Promise.reject(response);
      }

      // update table
      loadCollectionsHistoryTable();

      if (successCallback) {
        successCallback();
      }
    })
    .catch((response) => {
      console.log(response.status, response.statusText);
      showErrorFunction("An error has occurred.");
    });
};
