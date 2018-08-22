"use strict";

class Game {
    constructor() {
        this.myPosition = { y: 0 };
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
            this.myPosition.y = e.data.Game.Players[this.myPlayerID].Pos.Y;
            this.yBoundaries = [0, e.data.Game.Height - e.data.Game.Players[this.myPlayerID].Height]
            this.initMovementEvents();

            this.paintBoard(e.data.Game.Width, e.data.Game.Height);
            this.paintBall(e.data.Game.Ball);
            this.paintPlayer(e.data.Game.Players[this.myPlayerID]);
        });

        document.addEventListener("updategame", (e) => {
            this.paintBall(e.data.Ball);

            let opponent = e.data.Players[this.opponentID];
            !!opponent ? this.paintPlayer(opponent) : this.removePlayer(this.opponentID);
        });
    }

    initMovementEvents() {
        document.addEventListener("keypress", (e) => {
            let newPosition = this.myPosition.y;
            switch(e.keyCode) {
                case 38: // ArrowUp
                    newPosition = this.myPosition.y - 4;
                    break;
                case 40:
                    newPosition = this.myPosition.y + 4;
                    break;
            }
            this.myPosition.y = Math.min(this.yBoundaries[1], Math.max(this.yBoundaries[0], newPosition))
            this.movePlayer(this.myPlayerID, this.myPosition.y);
            fireEvent("onPlayerMove", document, this.myPosition.y);
        })
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

    movePlayer(id, y) {
        let ui = this._ui.players[id];
        ui.style.top = y + "px";
    }

    removePlayer(id) {
        this._ui.players[id].style.display = "none";
    }
}
