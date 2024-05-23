// Fonction pour remplacer un lien spécifique
function remplacerLien() {
  // Ciblez l'élément contenant le lien que vous souhaitez modifier
  var lien = document.querySelector(".button wc-backward");

  if (lien) {
    // Remplacez l'URL actuelle par votre nouvelle URL
    lien.href = "https://nouvelle-url.com";
  }
}

// Exécutez la fonction lorsque le document est chargé
document.addEventListener("DOMContentLoaded", function () {
  remplacerLien();
});
