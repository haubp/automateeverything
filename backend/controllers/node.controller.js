const nodeModel = require("../models/node.model");
const nodeCommandsModel = require("../models/node_commands.model");
const crypto = require("crypto");
const fs = require('fs');
const path = require('path');

async function getNode(req, res) {
    if (!req.body.user_id) {
        res.status(200).json({status: "Node ID not found"});
        return;
    }
    const nodes = await nodeModel.getNodesFromUserId(req.body.user_id);
    res.status(200).json(nodes.map(node => {return {node_id: node.node_id}}));
}

async function deleteNode(req, res) {
    // TODO : Check if delete node ok or not
    if (!req.body.node_id) {
        res.status(404).json({status: "Node ID not found"});
        return;
    }
    await nodeModel.deleteNodeFromNodeId(req.body.node_id);
    res.status(200).json({});
}

async function createNode(req, res) {
    if (!req.body.user_id) {
        res.status(404).json({status: "User ID not found"});
        return;
    }
    const sha256 = crypto.createHash("sha256");
    const node_id = sha256.update(req.body.user_id + Date.now()).digest("hex");
    // TODO : Check if create node ok or not
    await nodeModel.createNodeForUserId(req.body.user_id, node_id);
    res.status(200).json({node_id});
}

async function addNodeCommand(req, res) {
    if (!req.body.node_id_list || !req.body.command) {
        res.status(404).json({status: "Invalid node_id_list or command"});
        return;
    }
    // TODO: Check db connection
    await nodeCommandsModel.addNodeCommand(req.body.node_id_list, req.body.command);
    res.status(200).json({});
}

async function getNodeCommand(req, res) {
    if (!req.body.node_id) {
        res.status(404).json({status: "Invalid node_id"});
        return;
    }
    const node_command = await nodeCommandsModel.getNodeCommand(req.body.node_id);
    res.status(200).json(node_command.command);
}

async function reportNodeResult(req, res) {
    // Query user_id from node_id
    if (!req.query.node_id) {
        res.status(404).json({status: "Invalid node_id"});
        return;
    }
    const user_id = await nodeModel.getUserIdFromNodeId(req.query.node_id);
    console.log(user_id);
    // Save body json to data folder
    if (fs.existsSync(path.join(__dirname, "..", "public", "user_" + user_id, "test_result"))) {
        console.log(`Save result from node ${req.query.node_id} into ${user_id} storage`);
        const nodeData = JSON.stringify(req.body);
        fs.writeFileSync(path.join(__dirname, "..", "public", "user_" + user_id, "test_result", req.query.node_id + ".json"), nodeData);
    } else {
        console.log(`${user_id} storage not found`);
    }
    res.status(200).json({});
}

module.exports = {
    getNode,
    deleteNode,
    createNode,
    addNodeCommand,
    getNodeCommand,
    reportNodeResult
};