import WebSocket, { WebSocketServer } from "ws";

const wss = new WebSocketServer({
  port: 8080,
});

const rooms = new Map<string, WebSocket[]>();

wss.on("connection", (ws) => {
  console.log("New client connected");
  ws.send("Welcome new client");

  ws.on("message", (data) => {
    console.log("Received: %s", data);
    handleMessage(ws, `${data}`);
    // wss.clients.forEach((client) => {
    //   console.log("received and distributing data", `${data}`);
    //   client.send(`${data}`);
    // });
  });

  ws.on("close", () => {
    console.log("web socket closed successfully.....");
  });
  ws.on("error", (error) => {
    console.error("WebSocket error:", error);
  });
});

wss.on("error", (error) => {
  console.error("WebSocket server error:", error);
});

console.log("log the message result", rooms);

function handleMessage(ws: WebSocket, message: string) {
  ws.send(JSON.stringify(message));

  try {
    const data = JSON.parse(message) as {
      type: "message" | "join" | "leave" | "list" | "create" | "delete";
      room: string;
      text: string;
      username?: string;
    };
    const type = data.type;
    if (data.type === "join") {
      const room = data.room;

      if (!rooms.has(room)) {
        rooms.set(room, []);
      }
      rooms.get(room)?.push(ws);
    } else if (type === "message") {
      console.log("at first send message cases.......");
      const room = data.room;
      const clients = rooms.get(room) ?? [];
      clients.forEach((client) => {
        // if (client !== ws && client.readyState === WebSocket.OPEN) { //* if you want to send message to all clients except sender
        if (client.readyState === WebSocket.OPEN) {
          client.send(
            JSON.stringify({
              type: "message",
              room: data.room,
              text: data.text,
            })
          );
        }
      });
    }
  } catch (err) {
    console.log("error in parsing the message", err);
  }
}
