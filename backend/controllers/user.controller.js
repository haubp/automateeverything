const userModel = require("../models/user.model");
const fs = require("fs");
const path = require("path");
const crypto = require("crypto");

async function getUser(req, res) {
  if (!req.query.user_name || !req.query.password) {
    res.status(404).json({ status: "Invalid user_name or password" });
    return;
  }
  const sha256 = crypto.createHash("sha256");
  const user_id = sha256
    .update(req.query.user_name + req.query.password)
    .digest("hex");
  // TODO : Check if create user ok or not
  const users = await userModel.getUser(user_id);
  if (users.length == 0) {
    res.status(500).json({ status: "User not found" });
    return;
  }
  if (users.length != 1) {
    res.status(500).json({ status: "Corrupted User DB" });
    return;
  }
  res.status(200).json({ token: users[0].user_id });
}

async function deleteUser(req, res) {
  // TODO : Check if delete user ok or not
  await userModel.deleteUser(req.body.user_id);
  // Create User folder if not exist
  if (
    req.body.user_id &&
    fs.existsSync(
      path.join(__dirname, "..", "public", "user_" + req.body.user_id)
    )
  ) {
    fs.rmSync(
      path.join(__dirname, "..", "public", "user_" + req.body.user_id),
      { recursive: true, force: true }
    );
  }
  res.status(200).json({});
}

async function createUser(req, res) {
  if (req.body.user_name && req.body.password) {
    const sha256 = crypto.createHash("sha256");
    const user_id = sha256
      .update(req.body.user_name + req.body.password)
      .digest("hex");
    // TODO : Check if create user ok or not
    await userModel.createUser(user_id, req.body.user_name);

    // Create User folder if not exist
    if (
      !fs.existsSync(path.join(__dirname, "..", "public", "user_" + user_id))
    ) {
      fs.mkdirSync(path.join(__dirname, "..", "public", "user_" + user_id));
      fs.mkdirSync(
        path.join(__dirname, "..", "public", "user_" + user_id, "test_template")
      );
      fs.mkdirSync(
        path.join(__dirname, "..", "public", "user_" + user_id, "test_result")
      );
    }

    res.status(200).json({ token: user_id });
  } else {
    res.status(404).json({});
  }
}

module.exports = {
  getUser,
  deleteUser,
  createUser,
};
