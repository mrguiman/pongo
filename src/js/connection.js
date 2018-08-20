'use strict';

class Connection {
    constructor() {
        let ws = new WebSocket("ws://localhost:8080/websocket");
        let event = document.createEvent("customevent");

        ws.onopen = () => {
            // Give window responsability for pinging the server to keep the connection alive
            window.setInterval(() => { ws.send("PING") }, 1000)
            ws.send("READY");
        };

        ws.onmessage = (evt) => {
            let data = JSON.parse(evt.data)

            switch(data.Type) {
                case "INIT":
                    fireEvent("onInit", document, data);
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
