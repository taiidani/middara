import { TabList } from "./tabs.js";

document.addEventListener("DOMContentLoaded", function (evt) {
    // Load the game screen's tab list
    // https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Roles/tablist_role
    let tabList = document.querySelector("[role=tablist]");
    if (tabList) {
        TabList(tabList);
    }
});
