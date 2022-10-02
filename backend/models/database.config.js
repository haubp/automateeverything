const knex = require("knex");

const db_config = {
    host: '127.0.0.1',
    user: 'postgres',
    password: 'a',
    database: 'autmation'
}

const database = knex({
    client: 'pg',
    connection: {
        host: db_config.host,
        user: db_config.user,
        password: db_config.password,
        database: db_config.database
    }
});

module.exports = database;