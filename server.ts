import express from "express";

const app = express();

console.log("dir", __dirname);
app.use((req, res) => {
  res.sendFile("./src/index.html", { root: __dirname });
});
app.listen(3000, () => {
  console.log("Server is running on port 3000");
});
