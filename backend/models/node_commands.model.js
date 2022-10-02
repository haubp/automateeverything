const database = require("./database.config");

const nodeCommandsModel = {
    addNodeCommand: async function(node_id_list, command) {
        node_id_list.forEach(async (node_id) => {
            await database("node_commands")
                    .insert({
                            node_id: node_id,
                            command: command
                        })
                    .catch(e => console.log(e));
        });
    },
    getNodeCommand: async function(node_id) {
        const node_commands = await database("node_commands")
                                        .select("*")
                                        .where({node_id})
                                        .catch(e => console.log(e));
        this.deleteNodeCommand(node_id);
        return node_commands[0];
    },
    deleteNodeCommand: async function(node_id) {
        console.log(`delete command for node ${node_id}`);
        await database("node_commands")
                .del()
                .where({node_id})
                .catch(e => console.log(e));
    }
};

module.exports = nodeCommandsModel;