import React, { useState } from "react";
import BackendConfig from "../BackendConfig";
import axios from "axios";

export default function (props) {
  let [authMode, setAuthMode] = useState("signin");
  let [userName, setUserName] = useState("");
  let [password, setPassword] = useState("");

  const changeAuthMode = () => {
    setAuthMode(authMode === "signin" ? "signup" : "signin");
    setUserName("");
    setPassword("");
  };

  const onSignInClick = async (e) => {
    e.preventDefault();

    if (!userName || !password) {
      return;
    }

    const config = {
      headers: {
        "Content-Type": "application/json",
      },
    };
    const response = await axios.get(
      `http://${BackendConfig.url}:${BackendConfig.port}${BackendConfig.api.user}?user_name=${userName}&password=${password}`,
      config
    );

    const token = response.data ? response.data : '';
  };

  const onSignUpClick = async (e) => {
    e.preventDefault();

    if (!password || !userName) {
      return;
    }

    const config = {
      headers: {
        "Content-Type": "application/json",
      },
    };
    const body = {
      user_name: userName,
      password: password,
    };
    console.log(body);
    const response = await axios.post(
      `http://${BackendConfig.url}:${BackendConfig.port}${BackendConfig.api.user}`,
      body,
      config
    );

    console.log(response.data);
  };

  if (authMode === "signin") {
    return (
      <div className="Auth-form-container">
        <form className="Auth-form" onSubmit={onSignInClick}>
          <div className="Auth-form-content">
            <h3 className="Auth-form-title">Sign In</h3>
            <div className="text-center">
              Not registered yet?{" "}
              <span className="link-primary" onClick={changeAuthMode}>
                Sign Up
              </span>
            </div>
            <div className="form-group mt-3">
              <label>User Name</label>
              <input
                type="text"
                className="form-control mt-1"
                placeholder="admin"
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
              />
            </div>
            <div className="form-group mt-3">
              <label>Password</label>
              <input
                type="password"
                className="form-control mt-1"
                placeholder="******"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <div className="d-grid gap-2 mt-3">
              <button type="submit" className="btn btn-primary">
                Sign In
              </button>
            </div>
          </div>
        </form>
      </div>
    );
  }

  return (
    <div className="Auth-form-container">
      <form className="Auth-form" onSubmit={onSignUpClick}>
        <div className="Auth-form-content">
          <h3 className="Auth-form-title">Sign Up</h3>
          <div className="text-center">
            Already registered?{" "}
            <span className="link-primary" onClick={changeAuthMode}>
              Sign In
            </span>
          </div>
          <div className="form-group mt-3">
            <label>User Name</label>
            <input
              type="text"
              className="form-control mt-1"
              placeholder="admin"
              value={userName}
              onChange={(e) => setUserName(e.target.value)}
            />
          </div>
          <div className="form-group mt-3">
            <label>Password</label>
            <input
              type="password"
              className="form-control mt-1"
              placeholder="*******"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="d-grid gap-2 mt-3">
            <button type="submit" className="btn btn-primary">
              Sign Up
            </button>
          </div>
        </div>
      </form>
    </div>
  );
}
