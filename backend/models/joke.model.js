const knex = require("knex");

const database = knex({
    client: 'pg',
    connection: {
        host: '127.0.0.1',
        user: 'postgres',
        password: '',
        database: 'porfolio'
    }
});

const joke = {
    getJokes: async function() {
        let jokes = await database('jokes').select("*").catch(e => console.log(e));
        return jokes;
    },
    createJoke: async function(jokes) {
        jokes.forEach(async function(joke) {
            await database("jokes").insert({id: joke.id, joke: joke.joke, likes: joke.likes, dislikes: joke.dislikes});
        });
    },
    updateJoke: async function(jokeId, joke, likes, dislikes) {
        let joke = await database('jokes').select('*').where({id: jokeId}).catch(e => console.log(e));
        joke[0].likes += likes;
        joke[0].dislikes += dislikes;
        joke[0].joke = joke;
        await database("jokes").where({id: jokeId}).update({joke: joke[0].joke, likes: joke[0].likes, dislikes: joke[0].dislikes});
    }
};

module.exports = joke;