const express = require("express");
const { getUser, deleteUser, createUser, updateTemplate } = require("../controllers/user.controller");

const userRouter = express.Router();

userRouter.get("/", getUser);
userRouter.delete("", deleteUser);
userRouter.post("/", createUser);

userRouter.post("/template", updateTemplate);

module.exports = userRouter;