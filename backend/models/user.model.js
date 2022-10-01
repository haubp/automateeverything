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

// Get current list of user from database
let currentTotalUser = 0;
(async function() {
    let totalUser = await database('users').count("user_id as CNT").catch(e => console.log(e));
    currentTotalUser = totalUser[0].CNT;
    console.log(`Current Total User is ${currentTotalUser}`);
})()

const userModel = {
    getUser: async function() {
        let users = await database('users').select("*").catch(e => console.log(e));
        return users;
    },
    createUser: async function(user) {
        console.log("Create new user");
        await database("users").insert({   
                                            user_id: currentTotalUser,
                                            authen:  user.authen
                                       } )
                               .catch(e => console.log(e));
        // TODO: Check if create user ok or not before increase user counter
        currentTotalUser++;
    },
    deleteUser: async function(user_id) {
        await database("users").del().where({user_id: user_id}).catch(e => console.log(e));
    }
};

module.exports = userModel;