import { TabList } from "./tabs.js";
import { Search } from "./search.js";

document.addEventListener("DOMContentLoaded", function (evt) {
    // Load the game screen's tab list
    // https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Roles/tablist_role
    let tabList = document.querySelector("[role=tablist]");
    if (tabList) {
        TabList(tabList);
    }

    // Enable search filtering on the game screen
    let search = document.querySelector("[role=search]");
    if (search) {
        Search(search);
    }
});
