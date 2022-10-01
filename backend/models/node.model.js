const knex = require("knex");

const database = knex({
    client: 'pg',
    connection: {
        host: '127.0.0.1',
        user: 'postgres',
        password: 'a',
        database: 'autmation'
    }
});

const nodeModel = {
    getNode: async function() {
        let nodes = await database('nodes').select("*").catch(e => console.log(e));
        return nodes;
    },
    createNode: async function(node) {
        await database("nodes").insert({node_id: node.node_id, user_id: node.user_id}).catch(e => console.log(e));
    },
    deleteNode: async function(node_id) {
        await database("nodes").del().where({node_id: node_id}).catch(e => console.log(e));
    }
};

module.exports = nodeModel;