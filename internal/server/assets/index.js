let app = document.getElementById("app");
if (app) {
    let nav = document.querySelector("#app > nav");

    // On page load, ensure that the last loaded tab is restored
    document.addEventListener("DOMContentLoaded", function (evt) {
        let t = document.location.hash;
        if (t === "") {
            return;
        }

        // Strip the "#tab-" prefix
        t = t.slice(5);

        let btns = nav.querySelectorAll("button");
        btns.forEach(element => {
            if (element.dataset.target === t) {
                element.click();
            }
        });
    });

    // Display the tab according to the clicked button
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
        document.location.hash = "#tab-" + t;
    });
}
