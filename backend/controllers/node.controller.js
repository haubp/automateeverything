const nodeModel = require("../models/node.model");

async function getNode(req, res) {
    const nodes = await nodeModel.getNode();
    res.status(200).json(nodes);
}

async function deleteNode(req, res) {
    // TODO : Check if delete node ok or not
    await nodeModel.deleteNode(req.body.node_id);
    res.status(200).json({});
}

async function createNode(req, res) {
    // TODO : Check if create node ok or not
    await nodeModel.createNode(req.body);
    res.status(200).json({});
}

module.exports = {
    getNode,
    deleteNode,
    createNode
};