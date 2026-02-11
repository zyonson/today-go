let allPlaces = [];

fetch("/api/data")
  .then((res) => res.json())
  .then((data) => {
    const event = data.eventList[0];
    const allStoreName = "飲食店";
    allPlaces = data.places.places;
    const allPlace = allPlaces[allStoreName] || [];
    document.querySelector("#event-title strong").textContent = event.summary;
    document.getElementById("destination").textContent =
      event.location || "場所情報なし";

    renderPlaces(allPlace);

    // Google Maps API 読み込み
    const script = document.createElement("script");
    script.src = `https://maps.googleapis.com/maps/api/js?key=${data.apiKey}&callback=initMap`;
    script.async = true;
    script.defer = true;
    document.body.appendChild(script);
  });
