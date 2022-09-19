const express = require("express");
const path = require("path");
const jokeRouter = require("./routes/joke.router");

const app = express();
const PORT = 3434;

app.use((req, res, next) => {
    const start = Date.now();
    next();
    const delta = Date.now() - start;
    console.log(`${req.method} ${req.baseUrl} ${req.path} ${delta}ms`);
});
app.use('/site', express.static(path.join(__dirname, "public")));
app.use(express.json());

app.use("/api/v1", jokeRouter);

app.listen(PORT, () => {
    console.log(`Listening on ${PORT} ...`);
});