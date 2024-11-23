export function Search(elem) {
    if (!elem) {
        console.warn("Empty element provided, cannot load tab list");
        return;
    }

    // Perform the search upon every typed character
    elem.addEventListener("keyup", evtChange);

    console.debug("Search bar loaded");
}

function evtChange(evt) {
    evt.stopPropagation();

    let filter = evt.target.value.toLowerCase();
    console.debug("Searched for '" + filter + "'");

    let controls = evt.target.getAttribute("aria-controls");
    if (!controls) {
        return
    }

    let content = document.getElementById(controls);
    if (!content) {
        console.warn("No element found for search target '" + controls + "'");
        return
    }

    let lists = content.querySelectorAll("dl");
    lists.forEach(list => {
        let terms = list.querySelectorAll("dt");
        terms.forEach(term => {
            let found = term.innerText.toLowerCase().includes(filter);

            // Toggle visibility for both the dt and the dd
            term.classList.toggle("hide", !found);
            term.nextElementSibling.classList.toggle("hide", !found);
        });
    });
}
