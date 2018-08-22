'use strict';

class Connection {
    constructor() {
        let ws = new WebSocket("ws://localhost:8080/websocket");
        let event = document.createEvent("customevent");

        ws.onopen = () => {
            // Give window responsability for pinging the server to keep the connection alive
            window.setInterval(() => { ws.send(JSON.stringify({Type:"PING"})) }, 1000)
            ws.send(JSON.stringify({Type: "READY"}));
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

        ws.onclose = (evt) => {
            console.log("Connection closed: " + evt.reason);
        };

        document.addEventListener("onPlayerMove", (e) => {
            ws.send(JSON.stringify({Type: "UPDATEPOSITION", Y: e.data}));
        });
        return ws
    }
}
