const genreSelect = document.getElementById("genre-tabs");
genreSelect.addEventListener("change", (e) => {
  const genre = e.target.value;
  if (!genre) return;

  renderPlaces(allPlaces[genre] || []);
});
