const express = require("express");
const path = require("path");
const nodeRouter = require("./routes/node.router");
const userRouter = require("./routes/user.router");

const app = express();
const PORT = 6868;

/*
    Logger
*/
app.use((req, res, next) => {
    const start = Date.now();
    next();
    const delta = Date.now() - start;
    console.log(`${req.method} ${req.baseUrl} ${req.path} ${delta}ms`);
});

/*
    Serve public file
*/
app.use('/public', express.static(path.join(__dirname, "public")));

/*
    JSON middleware
*/
app.use(express.json());

/*
    Node Router
*/
app.use("/api/v1/node", nodeRouter);

/* 
    User Router
*/
app.use("/api/v1/user", userRouter);

app.listen(PORT, () => {
    console.log(`Listening on ${PORT} ...`);
});