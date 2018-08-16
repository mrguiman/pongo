"use strict";

class Game {
    constructor() {
        this.Ball = {}

        document.addEventListener("gameinit", (e) => {
            this.update(e.data);
            fireEvent("paintboard", document, { Width: e.data.Width, Height: e.data.Height });
        })
        document.addEventListener("gameupdate", (e) => {
            this.update(e.data);
            fireEvent("repaint", document, e.data);
        });
    }

    update(data) {
        this.Ball = data.Ball;
        this.Board = { Width: data.Width, Height: data.height }
    }

}
