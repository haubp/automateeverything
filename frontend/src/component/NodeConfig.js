import React, { useState } from "react";
import Card from "react-bootstrap/Card";

import "./NodeConfig.css";

export default function () {
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
        </Card.Body>
      </Card>
    </div>
  );
}
