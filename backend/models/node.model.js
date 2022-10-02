const database = require("./database.config");

const nodeModel = {
    getNodesFromUserId: async function(user_id) {
        let nodes = await database('nodes')
                            .select("*")
                            .where({user_id: user_id})
                            .catch(e => console.log(e));
        return nodes;
    },
    getUserIdFromNodeId: async function(node_id) {
        const result_user_id = await database('nodes')
                            .select("user_id")
                            .where({node_id})
                            .catch(e => console.log(e));
        if (result_user_id.length != 1) {
            return "";
        }
        return result_user_id[0].user_id;
    },
    createNodeForUserId: async function(user_id, node_id) {
        await database("nodes")
                .insert({
                            node_id: node_id,
                            user_id: user_id
                        })
                .catch(e => console.log(e));
    },
    deleteNodeFromNodeId: async function(node_id) {
        await database("nodes")
                .del()
                .where({node_id: node_id})
                .catch(e => console.log(e));
    }
};

module.exports = nodeModel;