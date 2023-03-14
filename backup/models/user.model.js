const database = require("./database.config");

const userModel = {
    getUser: async function(user_id) {
        let user = await database('users')
                            .select("*")
                            .where({user_id: user_id})
                            .catch(e => console.log(e));
        return user;
    },
    createUser: async function(user_id, user_name) {
        console.log("Create new user");
        await database("users")
                .insert({   
                            user_id: user_id,
                            user_name:  user_name
                        })
                .catch(e => console.log(e));
    },
    deleteUser: async function(user_id) {
        await database("users")
                .del()
                .where({user_id: user_id})
                .catch(e => console.log(e));
    }
};

module.exports = userModel;