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
        if (element === evt.target) {
            element.classList.remove("outline");
            element.setAttribute("aria-selected", "true");
        } else {
            element.classList.add("outline");
            element.setAttribute("aria-selected", "false");
        }
    });

    let tabPanels = app.querySelectorAll("[role=tabpanel]");
    tabPanels.forEach(element => {
        if (element.id === controls) {
            element.classList.remove("hide");
        } else {
            element.classList.add("hide");
        }
    });
    document.location.hash = "#tab-" + controls;
}
