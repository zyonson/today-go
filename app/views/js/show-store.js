function renderPlaces(places) {
  const storeContainer = document.getElementById("store-container");
  storeContainer.innerHTML = "";

  if (Array.isArray(places) && places.length > 0) {
    places.forEach((place) => {
      const card = document.createElement("div");
      card.className = "card mb-3";
      card.innerHTML = `
        <div class="card-body">
          <h5 class="card-title">${place.displayName.text}</h5>
          <p class="card-text">評価：${place.rating}</p>
          <a href="${place.googleMapsUri}" class="btn btn-primary" target="_blank" rel="noopener noreferrer">地図を見る</a>
        </div> `;
      storeContainer.appendChild(card);
    });
  } else {
    storeContainer.innerHTML = "<p>店舗情報が見つかりませんでした。</p>";
  }
}
