const express = require("express");
const { getNode, deleteNode, createNode } = require("../controllers/node.controller");

const nodeRouter = express.Router();

nodeRouter.get("/", getNode);
nodeRouter.delete("/", deleteNode);
nodeRouter.post("/", createNode);

module.exports = nodeRouter;