const joke = require("../models/joke.model");

async function getJoke(req, res) {
    const jokes = await joke.getJokes();
    res.status(200).json(jokes);
}

async function updateJoke(req, res) {
    const jokeId = req.params.jokeId;
    await joke.updateJoke(jokeId, req.body.joke, req.body.likes, req.body.dislikes);
    res.status(200).json({});
}

async function createJoke(req, res) {
    await joke.createJoke(req.body);
    res.status(200).json({});
}

module.exports = {
    getJoke,
    updateJoke,
    createJoke
};