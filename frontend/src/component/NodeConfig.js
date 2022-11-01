import React from "react";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import axios from "axios";

import "./NodeConfig.css";

export default function () {
  const baseUrl = "http://10.40.50.34:6868";

  const runHandler = async (e) => {
    await axios({
      method: "post",
      url: baseUrl + "/api/v1/node/command",
      headers: {},
      data: {
        node_id_list: [
          "ec5bf52b73d078fa0ac7d27281da2104a588440e8571c2c4f00aa69d5464e387",
        ],
        command: "RUN",
      },
    });

    alert("Pushed Node Run Command");
  };
  return (
    <div className="nodeBoard">
      <Card style={{ width: "18rem" }}>
        <Card.Body>
          <Card.Title>
            <div
              style={{
                display: "flex",
                justifyContent: "space-between",
              }}
            >
              <span>a0f90acdd7</span>
            </div>
          </Card.Title>
          <label>Test Category</label>
          <div className="d-flex form-group mt-1">
            <input
              type="password"
              className="form-control mt-1"
              placeholder="All"
            />
          </div>
          <label>Test Group</label>
          <div className="d-flex form-group mt-1">
            <input
              type="password"
              className="form-control mt-1"
              placeholder="All"
            />
          </div>
          <label>Test Cases</label>
          <div className="d-flex form-group mt-1">
            <input
              type="password"
              className="form-control mt-1"
              placeholder="From 0"
            />
            <p className="ms-1 me-1"> _ </p>
            <input
              type="password"
              className="form-control mt-1"
              placeholder="To n"
            />
          </div>
          <div className="run-area">
            <Button
              className="upload-btn"
              variant="outline-success"
              onClick={runHandler}
            >
              Run
            </Button>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
