const express = require("express");
const { getNode, deleteNode, createNode, addNodeCommand, getNodeCommand, reportNodeResult } = require("../controllers/node.controller");

const nodeRouter = express.Router();

nodeRouter.get("/", getNode);
nodeRouter.delete("/", deleteNode);
nodeRouter.post("/", createNode);

nodeRouter.post("/command", addNodeCommand);
nodeRouter.get("/command", getNodeCommand);

nodeRouter.post("/result", reportNodeResult);

module.exports = nodeRouter;