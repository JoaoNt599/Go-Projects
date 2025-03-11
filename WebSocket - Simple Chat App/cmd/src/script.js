// const ws = new WebSocket("ws://localhost:8081/ws");

// ws.onopen = function () {
//     console.log("Connected to WebSocket");
// };

// ws.onerror = function (error) {
//     console.error("WebSocket error:", error);
// };

// ws.onmessage = function (event) {
//     console.log("Message received:", event.data);
//     const msg = JSON.parse(event.data);

//     // Get formatted time
//     const now = new Date();
//     const timeString = now.toLocaleTimeString("en-US", { hour: "2-digit", minute: "2-digit"});

//     // Append message to all chat windows
//     document.querySelectorAll(".messages").forEach(chatBox => {
//         const item = document.createElement("li");
//         item.innerHTML = `<strong>${msg.username}</strong>: ${msg.message} <span class="timestamp">${timeString}</span>`;
//         chatBox.appendChild(item);
//     });
// };

// ws.onclose = function () {
//     console.log("WebSocket connection closed");
// };

// // Handle form submission for all chat windows
// document.querySelectorAll(".chat-form").forEach(form => {
//     form.onsubmit = function (event) {
//         event.preventDefault();
//         const username = form.querySelector(".username").value;
//         const message = form.querySelector(".message").value;

//         if (ws.readyState === WebSocket.OPEN) {
//             ws.send(JSON.stringify({ username, message }));
//             console.log("📤 Message sent:", username, message);
//         } else {
//             console.warn("WebSocket is not ready");
//         }

//         form.querySelector(".message").value = '';
//     };
// });

const ws = new WebSocket("ws://localhost:8081/ws");

ws.onopen = function () {
    console.log("✅ Connected to WebSocket");
};

ws.onmessage = function (event) {
    console.log("📩 Message received:", event.data);
    const msg = JSON.parse(event.data);

    document.querySelectorAll(".messages").forEach(chatBox => {
        const item = document.createElement("li");
        item.innerHTML = `<strong>${msg.username}</strong>: ${msg.message} <span class="timestamp">${msg.time}</span>`;
        chatBox.appendChild(item);
    });
};

document.querySelectorAll(".chat-form").forEach(form => {
    form.onsubmit = function (event) {
        event.preventDefault();
        const username = form.querySelector(".username").value;
        const message = form.querySelector(".message").value;
        const now = new Date().toLocaleTimeString("en-US");

        if (ws.readyState === WebSocket.OPEN) {
            ws.send(JSON.stringify({ username, message, time: now }));
        } else {
            console.warn("⚠️ WebSocket is not ready");
        }

        form.querySelector(".message").value = '';
    };
});

