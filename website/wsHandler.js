const socket = new WebSocket("ws://localhost:8080/main");
const chatList = document.getDocumentBy

console.log('connecting to server')

// Event handler for when the WebSocket connection is opened
socket.onopen = function(event) {
    console.log("WebSocket connection opened.");
};

// Event handler for when a message is received from the server
socket.onmessage = function(event) {
    console.log("Received message from server: " + event.data);
};

// Event handler for errors
socket.onerror = function(event) {
    console.error("WebSocket error: " + event);
};

// Event handler for when the connection is closed
socket.onclose = function(event) {
    if (event.wasClean) {
        console.log("WebSocket connection closed cleanly, code=" + event.code + ", reason=" + event.reason);
    } else {
        console.error("WebSocket connection abruptly closed.");
    }
};

// Send a message when the button is clicked
document.getElementById("send_message").addEventListener("click", function() {
    const message = "Hello, server!";
    socket.send(message);
});



