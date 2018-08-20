'use strict';

window.app = {
    _game: {},
    _ui: {},
    _ws: null,

    init: function() {
        this._ui["splash"] = document.getElementById("splash");
        this._ui["game"] = document.getElementById("game");

        new Connection();
        new Game();

        document.addEventListener("onReady", () => {
            this._ui.splash.style.display = "none";
        })

        fireEvent("onReady", document);
    }
}


function fireEvent(name, target, data) {
    var e = document.createEvent("customevent");
    e.initEvent(name, false, false);
    e.data = data;

    target.dispatchEvent(e)
}
