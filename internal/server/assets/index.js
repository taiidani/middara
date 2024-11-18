let app = document.getElementById("app");
if (app) {
    let nav = document.querySelector("#app > nav");

    nav.addEventListener("click", function (evt) {
        evt.stopPropagation();

        let btns = nav.querySelectorAll("button");
        btns.forEach(element => {
            if (element === evt.target) {
                element.className = "";
                element.ariaCurrent = "true";
            } else {
                element.className = "outline";
                element.ariaCurrent = "false";
            }
        });

        let t = evt.target.dataset.target;
        let tabs = app.querySelectorAll(".tab");
        tabs.forEach(element => {
            if (element.id === t) {
                element.classList.remove("hide");
            } else {
                element.classList.add("hide");
            }
        });
    });
}
