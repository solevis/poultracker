let COLLECTIONS_HISTORY_CHART = null;
let COLLECTIONS_HISTORY_RAW_DATA = {};

const loadCollectionsHistoryChart = () => {
  fetch("/api/collections", {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
  })
    .then((response) => response.json())
    .then((json) => (COLLECTIONS_HISTORY_RAW_DATA = json))
    .then(() => {
      if (COLLECTIONS_HISTORY_CHART !== null) {
        COLLECTIONS_HISTORY_CHART.destroy();
        COLLECTIONS_HISTORY_CHART = null;
      }

      COLLECTIONS_HISTORY_CHART = createChart();
    });
};

const getCollectionsHistoryChartData = (start, end) => {
  let days = [];

  for (
    let dt = new Date(start);
    dt <= new Date(end);
    dt.setDate(dt.getDate() + 1)
  ) {
    let number = NaN;
    let day = formatDate(new Date(dt));

    const collection = COLLECTIONS_HISTORY_RAW_DATA.find(
      (collection) => collection.laidDate === day
    );
    if (collection) {
      number = collection.number;
    }

    days.push({
      laidDate: day,
      number: number,
    });
  }

  return days;
};

function createChart() {
  let fifteenDaysAgoDate = new Date();
  fifteenDaysAgoDate.setDate(fifteenDaysAgoDate.getDate() - 21);
  const collectionsHistoryChartData = getCollectionsHistoryChartData(
    fifteenDaysAgoDate,
    new Date()
  );

  const maxYValue = Math.max(
    ...collectionsHistoryChartData.map((row) =>
      isNaN(row.number) ? 0 : row.number
    )
  );

  return new Chart(document.getElementById("myChart"), {
    type: "line",
    data: {
      labels: collectionsHistoryChartData.map((row) => row.laidDate),
      datasets: [
        {
          label: "Number of eggs",
          data: collectionsHistoryChartData.map((row) => row.number),
          borderColor: "#4f46e5",
          backgroundColor: "#4f46e5",
          fill: false,
          spanGaps: true,
          lineTension: 0.4,
          cubicInterpolationMode: "monotone",
        },
      ],
    },
    options: {
      responsive: true,
      scales: {
        y: {
          min: 0,
          max: maxYValue + 1,
        },
      },
    },
  });
}
