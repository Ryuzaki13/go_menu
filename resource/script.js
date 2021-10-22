let hamburger = document.querySelector(".hamburger");
let menu = document.querySelector("nav > ul");
if (hamburger) {
    hamburger.onclick = () => {
        if (menu) {
            menu.classList.toggle("open");
        }
    }
}