function initMap() {
  const destination = document.getElementById("destination").textContent;

  navigator.geolocation.getCurrentPosition(
    function (position) {
      const origin = {
        lat: position.coords.latitude,
        lng: position.coords.longitude,
      };

      const map = new google.maps.Map(document.getElementById("map"), {
        zoom: 14,
        center: origin,
      });

      const directionsService = new google.maps.DirectionsService();
      const directionsRenderer = new google.maps.DirectionsRenderer();
      directionsRenderer.setMap(map);

      directionsService.route(
        {
          origin: origin,
          destination: destination,
          travelMode: google.maps.TravelMode.DRIVING,
        },
        function (response, status) {
          if (status === "OK") {
            directionsRenderer.setDirections(response);

            const leg = response.routes[0].legs[0];
            document.getElementById("duration").textContent = leg.duration.text;

            const steps = leg.steps;
            const drivingInfo = steps.map(
              (step) => `<div>${step.instructions}</div>`
            );
            document.getElementById("driving").innerHTML =
              drivingInfo.join("<hr>");
          } else {
            alert("ルート検索に失敗しました: " + status);
          }
        }
      );
    },
    function (error) {
      alert("現在地の取得に失敗しました: " + error.message);
    }
  );
}
