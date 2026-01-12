fetch("/api/data")
  .then((res) => res.json())
  .then((data) => {
    const event = data.eventList[0];
    document.querySelector("#event-title strong").textContent = event.summary;
    document.getElementById("destination").textContent =
      event.location || "場所情報なし";
    const storeContainer = document.getElementById("store-container");
    if (data.places && Array.isArray(data.places.places)) {
      data.places.places.forEach((place) => {
        const card = document.createElement("div");
        card.className = "card mb-3";
        card.innerHTML = ` <div class="card-body"> <h5 class="card-title">${place.displayName.text}</h5> <p class="card-text">評価：${place.rating}</p> <a href="${place.mapUri}" class="btn btn-primary" target="_blank" rel="noopener noreferrer">地図を見る</a> </div> `;
        storeContainer.appendChild(card);
      });
    } else {
      console.warn("places.places が存在しません");
    }
    // Google Maps API 読み込み
    const script = document.createElement("script");
    script.src = `https://maps.googleapis.com/maps/api/js?key=${data.apiKey}&callback=initMap`;
    script.async = true;
    script.defer = true;
    document.body.appendChild(script);
  });
