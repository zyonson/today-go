const CACHE_KEY = "allPlaces";
const CACHE_TTL_MS = 60 * 60 * 1000;

const cache = sessionStorage.getItem(CACHE_KEY);
let allPlaces = {};

if (!cache || Date.now() - JSON.parse(cache).timestamp > CACHE_TTL_MS) {
  fetch("/api/data")
    .then(async (res) => {
      const contentType = res.headers.get("content-type") || "";

      if (!res.ok) {
        if (contentType.includes("application/json")) {
          const errJson = await res.json();
          throw new Error(errJson.error || "データ取得に失敗しました");
        }
        const errText = await res.text();
        throw new Error(errText || "データ取得に失敗しました");
      }

      if (!contentType.includes("application/json")) {
        throw new Error("サーバーからのレスポンスがJSON形式ではありません");
      }

      return res.json();
    })
    .then((data) => {
      const todayEvent = data?.eventList?.[0];
      if (!todayEvent) throw new Error("本日のイベントが見つかりませんでした");

      allPlaces = data?.places?.places || {};
      const apiKey = data?.apiKey;
      if (!apiKey) throw new Error("APIキーが取得できません");

      sessionStorage.setItem(
        CACHE_KEY,
        JSON.stringify({
          timestamp: Date.now(),
          storeData: allPlaces,
          todayEvent,
          apiKey,
        }),
      );

      showData(todayEvent, apiKey);
    })
    .catch((err) => {
      console.error(err);
      alert(err.message);
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
