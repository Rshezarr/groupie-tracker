ymaps.ready(init);

function init() {
  myMap = new ymaps.Map("geolocation", {
    center: [51.11, 71.45],
    zoom: 2,
  });

  let coords = document.getElementsByClassName("artist_cities");

  for (let coord of coords) {
    let newCoord = coord.innerHTML.replace(/-/g, "_");
    newCoord = newCoord.replace(/_/g, " ");
    console.log(newCoord);

    let myGeocoder = ymaps.geocode(newCoord, {
      results: 1,
      prefLang: "en",
    });

    myGeocoder.then(function (res) {
      myMap.geoObjects.add(res.geoObjects);
    });
  }
}
