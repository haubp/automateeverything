import React from "react";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import "./NodeReport.css";

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
          <div
            className="mt-5"
            style={{ display: "flex", justifyContent: "space-between" }}
          >
            <label>
              Pass
              <span
                className="badge badge-primary text-warap"
                style={{ width: "2rem", color: "black" }}
              >
                1
              </span>
            </label>
            <Button className="result-btn" variant="outline-success">
              Detail
            </Button>
          </div>
          <div
            className="mt-4"
            style={{ display: "flex", justifyContent: "space-between" }}
          >
            <label>
              Fail
              <span
                className="badge badge-primary text-warap"
                style={{ width: "2rem", color: "black" }}
              >
                0
              </span>
            </label>
            <Button className="result-btn" variant="outline-warning">
              Detail
            </Button>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
