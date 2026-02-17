const cache = sessionStorage.getItem("allPlaces");
let allPlaces = {};

if (!cache || Date.now() - JSON.parse(cache).timestamp > 60 * 60 * 1000) {
  fetch("/api/data")
    .then((res) => res.json())
    .then((data) => {
      const todayEvent = data.eventList[0];
      allPlaces = data.places.places;
      const apiKey = data.apiKey;

      // キャッシュに保存
      sessionStorage.setItem(
        "allPlaces",
        JSON.stringify({
          timestamp: Date.now(),
          storeData: allPlaces,
          todayEvent: todayEvent,
          apiKey: apiKey,
        }),
      );
      showData(todayEvent, apiKey);
    });
} else {
  // キャッシュから取得
  const parsed = JSON.parse(cache);
  allPlaces = parsed.storeData;
  const todayEvent = parsed.todayEvent;
  const apiKey = parsed.apiKey;
  showData(todayEvent, apiKey);
}
function showData(todayEvent, apiKey) {
  const allStoreName = "飲食店";
  const allPlace = allPlaces[allStoreName] || [];

  document.querySelector("#event-title strong").textContent =
    todayEvent.summary || "イベント名なし";
  document.getElementById("destination").textContent =
    todayEvent.location || "場所情報なし";

  renderPlaces(allPlace);

  const script = document.createElement("script");
  script.src = `https://maps.googleapis.com/maps/api/js?key=${apiKey}&callback=initMap`;
  script.async = true;
  script.defer = true;
  document.body.appendChild(script);
}
