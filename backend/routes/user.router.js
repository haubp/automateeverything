const express = require("express");
const { getUser, deleteUser, createUser } = require("../controllers/user.controller");

const userRouter = express.Router();

userRouter.get("/", getUser);
userRouter.delete("", deleteUser);
userRouter.post("/", createUser);

module.exports = userRouter;