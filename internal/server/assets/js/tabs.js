export function TabList(elem) {
    if (!elem) {
        console.warn("Empty element provided, cannot load tab list");
        return;
    }

    // Display the tab according to the clicked button
    elem.addEventListener("click", evtTabClick);

    let btns = elem.querySelectorAll("[role=tab]");
    if (btns.length === 0) {
        console.warn("Tab list did not have any tabs to load");
        return;
    }

    let t = document.location.hash;
    if (t.startsWith("#tab-")) {
        // Strip the "#tab-" prefix from the URL to find the id
        t = t.slice(5);
    }

    let clicked = false;
    btns.forEach(element => {
        if (element.getAttribute("aria-controls") === t) {
            element.click();
            clicked = true;
        }
    });

    // Default to the first tab
    if (!clicked) {
        btns[0].click();
    }
}

function evtTabClick(evt) {
    evt.stopPropagation();

    let controls = evt.target.getAttribute("aria-controls")
    if (!controls) {
        return
    }

    console.debug("Tab panel '" + controls + "' loading");
    let tabList = evt.target.closest("[role=tablist]");
    let btns = tabList.querySelectorAll("[role=tab]");
    btns.forEach(element => {
        let selected = element === evt.target;
        element.classList.toggle("outline", !selected);
        element.setAttribute("aria-selected", selected);
    });

    let tabPanels = app.querySelectorAll("[role=tabpanel]");
    tabPanels.forEach(element => {
        element.classList.toggle("hide", element.id !== controls);
    });
    document.location.hash = "#tab-" + controls;
}
