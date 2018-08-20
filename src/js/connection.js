'use strict';

class Connection {
    constructor() {
        let ws = new WebSocket("ws://localhost:8080/websocket");
        let event = document.createEvent("customevent");

        ws.onopen = () => {
            // Give window responsability for pinging the server to keep the connection alive
            window.setInterval(() => { ws.send("ping") }, 1000)
            ws.send("ready");
        };

        ws.onmessage = (evt) => {
            let data = JSON.parse(evt.data)

            switch(data.Type) {
                case "SET":
                    fireEvent("preparegame", document, data.Game);
                    ws.send("start")
                    break;
                case "UPDATE":
                    fireEvent("updategame", document, data.Game);
                    break;
            }
        };

        ws.onclose = () => {
            console.log("Connection closed");
        };

        return ws
    }
}
