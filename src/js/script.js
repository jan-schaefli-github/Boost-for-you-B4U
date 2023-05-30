
// Dark Mode / Light Mode
if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    // Dark Mode
    changeFavicon('./img/favicon_white.ico');
} else {
    // Light Mode
    changeFavicon('./img/favicon_black.ico');
}

function changeFavicon(faviconPath) {
    var favicon = document.querySelector('link[rel="icon"]');
    favicon.href = faviconPath;
}

// Menu
var burgerMenu = document.querySelector('.menu');
var menu = document.querySelector('.nav-links');
var menuVisible = false;

function toggleMenu() {
  if (!menuVisible) {
    menu.style.display = 'block';
    menuVisible = true;
  } else {
    menu.style.display = 'none';
    menuVisible = false;
  }
}

function handleScreenWidthChange(screenWidth) {
  if (screenWidth <= 480) {
    menu.style.display = 'none';
    burgerMenu.addEventListener('click', toggleMenu);
  } else {
    menu.style.display = 'flex';
    burgerMenu.removeEventListener('click', toggleMenu);
    menuVisible = false;
  }
}

// Überprüfen der Bildschirmbreite bei Initialisierung
handleScreenWidthChange(window.innerWidth);

// Überprüfen der Bildschirmbreite bei Größenänderungen
window.addEventListener('resize', function() {
  handleScreenWidthChange(window.innerWidth);
});


