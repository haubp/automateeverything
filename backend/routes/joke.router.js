const express = require("express");
const { getJoke, updateJoke, createJoke } = require("../controllers/joke.controller");

const jokeRouter = express.Router();

jokeRouter.get("/jokes", getJoke);
jokeRouter.put("/jokes/:jokeId", updateJoke);
jokeRouter.post("/jokes", createJoke);

module.exports = jokeRouter;