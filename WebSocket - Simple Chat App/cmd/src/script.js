const ws = new WebSocket("ws://localhost:8081/ws");

ws.onopen = function () {
    console.log("WebSocket Error");
};

ws.onerror = function (error) {
    console.error("", error);
};

ws.onmessage = function (event) {
    console.log("Message received: ", event.data);
    const msg = JSON.parse(event.data);
    
    // Add the message to all chat windows
    document.querySelectorAll(".messages").forEach(chatBox => {
        const item = document.createElement("li");
        item.textContent = msg.username + ": " + msg.message;
        chatBox.appendChild(item);
    });
};

ws.onclose = function () {
    console.log("WebSocket connection closed");
};

// Captures the submission of any chat form
document.querySelectorAll(".chat-form").forEach(form => {
    form.onsubmit = function (event) {
        event.preventDefault();
        const username = form.querySelector(".username").value;
        const message = form.querySelector(".message").value;
        if (ws.readyState === WebSocket.OPEN) {
            ws.send(JSON.stringify({ username, message }));
            console.log("Message sent:", username, message);
        } else {
            console.warn("WebSocket is not ready");
        }
        form.querySelector(".message").value = '';
    };
});
