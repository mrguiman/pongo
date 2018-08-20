"use strict";

class Game {
    constructor() {
        this._ui = {
            container: document.getElementById("game"),
            ball: document.getElementById("ball"),
            players: {
                1: document.getElementById("player" + 1),
                2: document.getElementById("player" + 2)
            }
        }

        document.addEventListener("onReady", () => {
            this._ui.container.style.display = "inline-block";
        });

        document.addEventListener("onInit", (e) => {
            this.myPlayerID = e.data.MyPlayerID;
            this.opponentID = this.myPlayerID == 1 ? 2 : 1;

            this.paintBoard(e.data.Game.Width, e.data.Game.Height);
            this.paintBall(e.data.Game.Ball);
            this.paintPlayer(e.data.Game.Players[this.myPlayerID - 1]);
        });

        document.addEventListener("updategame", (e) => {
            console.log(e.data);
            this.paintBall(e.data.Ball);

            let opponent = e.data.Players[this.opponentID - 1];
            opponent && this.paintPlayer(opponent);
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
