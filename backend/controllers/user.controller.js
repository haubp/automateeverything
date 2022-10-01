const userModel = require("../models/user.model");
const fs = require('fs');
const path = require('path');
const crypto = require("crypto");

async function getUser(req, res) {
    const users = await userModel.getUser();
    res.status(200).json(users);
}

async function deleteUser(req, res) {
    // TODO : Check if delete user ok or not
    await userModel.deleteUser(req.body.user_id);
    // Create User folder if not exist
    if ( req.body.user_name && 
         fs.existsSync(path.join(__dirname, "..", "public", "user_" + req.body.user_name))) {
        fs.rmSync(  path.join(__dirname, "..", "public", "user_" + req.body.user_name), 
                    {recursive: true, force: true})
    }
    res.status(200).json({});
}

async function createUser(req, res) {
    if (req.body.user_name && req.body.password) {
        const sha256 = crypto.createHash("sha256");
        const userAuthen = sha256.update(req.body.user_name + user.password).digest("base64");
        // TODO : Check if create user ok or not
        await userModel.createUser({"authen": userAuthen});

        // Create User folder if not exist
        if ( !fs.existsSync(path.join(__dirname, "..", "public", "user_" + req.body.user_name))) {
            fs.mkdirSync(path.join(__dirname, "..", "public", "user_" + req.body.user_name));
        }

        res.status(200).json({"token": userAuthen});
    } else {
        res.status(404).json({});
    }
}

module.exports = {
    getUser,
    deleteUser,
    createUser
};