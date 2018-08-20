"use strict";

class Game {
    constructor() {
        this._ui = {
            container: document.getElementById("game"),
            ball: document.getElementById("ball")
        }

        document.addEventListener("requeststart", () => {
            this._ui.container.style.display = "inline-block";
        });

        document.addEventListener("preparegame", (e) => {
            this.paintBoard(e.data.Width, e.data.Height);
        });

        document.addEventListener("updategame", (e) => {
            this.repaint(e.data);
        });
    }

    paintBoard(width, height) {
        this._ui.container.style.width = width + "px";
        this._ui.container.style.height = height + "px";
    }

    repaint(data) {
        this._ui.ball.style.left = data.Ball.Pos.Left + "px";
        this._ui.ball.style.top = data.Ball.Pos.Top + "px";
    }
}
