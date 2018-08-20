"use strict";

class Game {
    constructor() {
        this._ui = {
            container: document.getElementById("game"),
            ball: document.getElementById("ball"),
            players: {
                1: null,
                2: null
            }
        }

        document.addEventListener("onReady", () => {
            this._ui.container.style.display = "inline-block";
        });

        document.addEventListener("onInit", (e) => {
            this._ui.players[e.data.MyPlayer.ID] = document.getElementById("player" + e.data.MyPlayer.ID);
            this.paintBoard(e.data.Game.Width, e.data.Game.Height);
            this.paintBall(e.data.Game.Ball);
            this.paintPlayer(e.data.MyPlayer);
        });

        document.addEventListener("updategame", (e) => {
            this.paintBall(e.data.Ball);
        });
    }

    paintBoard(width, height) {
        this._ui.container.style.width = width + "px";
        this._ui.container.style.height = height + "px";
    }

    paintBall(ball) {
        this._ui.ball.style.left = ball.Pos.X + "px";
        this._ui.ball.style.top = ball.Pos.Y + "px";
    }

    paintPlayer(player) {
        let ui = this._ui.players[player.ID];
        ui.style.display = "block";
        ui.style.width = player.Width + "px";
        ui.style.height = player.Height + "px";
        ui.style.left = player.Pos.X + "px";
        ui.style.top = player.Pos.Y + "px";
    }
}
