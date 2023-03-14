const express = require("express");
const { getNode, deleteNode, createNode, addNodeCommand, getNodeCommand, reportNodeResult, uploadAgentLog, uploadAutoLog } = require("../controllers/node.controller");

const nodeRouter = express.Router();

nodeRouter.get("/", getNode);
nodeRouter.delete("/", deleteNode);
nodeRouter.post("/", createNode);

nodeRouter.post("/command", addNodeCommand);
nodeRouter.get("/command", getNodeCommand);

nodeRouter.post("/result", reportNodeResult);

nodeRouter.post("/agent/log", uploadAgentLog);
nodeRouter.post("/auto/log", uploadAutoLog)

module.exports = nodeRouter;