import React from "react";
import Card from "react-bootstrap/Card";
import Dropdown from "react-bootstrap/Dropdown";

import "./NodeConfig.css";

export default function () {
  return (
    <>
      <div className="nodeBoard">
        <Card style={{ width: "18rem" }}>
          <Card.Body>
            <Card.Title>Node ID</Card.Title>
            <div className="form-group mt-3">
              <label>Category</label>
              <Dropdown>
                <Dropdown.Toggle variant="success" id="dropdown-basic">
                  Dropdown Button
                </Dropdown.Toggle>

                <Dropdown.Menu>
                  <Dropdown.Item href="#/action-1">Action</Dropdown.Item>
                  <Dropdown.Item href="#/action-2">
                    Another action
                  </Dropdown.Item>
                  <Dropdown.Item href="#/action-3">
                    Something else
                  </Dropdown.Item>
                </Dropdown.Menu>
              </Dropdown>
            </div>
            <div className="form-group mt-3">
              <label>Group</label>
              <Dropdown>
                <Dropdown.Toggle variant="success" id="dropdown-basic">
                  Dropdown Button
                </Dropdown.Toggle>

                <Dropdown.Menu>
                  <Dropdown.Item href="#/action-1">Action</Dropdown.Item>
                  <Dropdown.Item href="#/action-2">
                    Another action
                  </Dropdown.Item>
                  <Dropdown.Item href="#/action-3">
                    Something else
                  </Dropdown.Item>
                </Dropdown.Menu>
              </Dropdown>
            </div>
            <div className="form-group mt-3">
              <label>Test Cases</label>
              <input
                type="password"
                className="form-control mt-1"
                placeholder="Enter test cases"
              />
            </div>
          </Card.Body>
        </Card>
      </div>
      <button type="submit" className="btn btn-primary">
        Run
      </button>
    </>
  );
}
