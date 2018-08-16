'use strict';

window.app = {
    _game: {},
    _ui: {},
    _ws: null,

    init: function() {
        this._ui["splash"] = document.getElementById("splash");
        this._ui["game"] = document.getElementById("game");
        this._ui["ball"] = document.getElementById("ball");
        new Connection();
        new Game();

        this._ui.splash.style.display = "none";
        this._ui.game.style.display = "inline-block";

        document.addEventListener("paintboard", (e) => { this.paintBoard(e.data.Width, e.data.Height)});
        document.addEventListener("repaint", (e) => { this.repaint(e.data)});
    },

    paintBoard: function(width, height) {
        this._ui.game.style.width = width + "px";
        this._ui.game.style.height = height + "px";
    },

    repaint: function(data) {
        this._ui.ball.style.left = data.Ball.Pos.Left + "px";
        this._ui.ball.style.top = data.Ball.Pos.Top + "px";
    }
}


function fireEvent(name, target, data) {
    var e = document.createEvent("customevent");
    e.initEvent(name, false, false);
    e.data = data;

    target.dispatchEvent(e)
}
