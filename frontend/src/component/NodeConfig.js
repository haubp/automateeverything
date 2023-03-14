import React, { useState } from "react";
import Card from "react-bootstrap/Card";
import Dropdown from "react-bootstrap/Dropdown";
import { CheckCircle } from "react-bootstrap-icons";

import "./NodeConfig.css";

export default function () {
  const onSelectCategory = function (e) {
    setSelectedCategory(e);
  };
  const onSelectGroup = function (e) {
    setSelectedGroup(e);
  };
  let [selectedCategory, setSelectedCategory] = useState("Category");
  let [selectedGroup, setSelectedGroup] = useState("Group");
  return (
    <>
      <div className="nodeBoard">
        <Card style={{ width: "18rem" }}>
          <Card.Body>
            <Card.Title>
              <div style={{ display: "flex", justifyContent: "space-between" }}>
                <span>Node ID</span>
                <CheckCircle style={{ alignSelf: "end" }} />
              </div>
            </Card.Title>
            <Dropdown className="mb-2" onSelect={onSelectCategory}>
              <Dropdown.Toggle
                className="w-75"
                variant="success"
                id="dropdown-basic"
              >
                {selectedCategory}
              </Dropdown.Toggle>
              <Dropdown.Menu>
                <Dropdown.Item href="#/action-1">Action</Dropdown.Item>
                <Dropdown.Item href="#/action-2">Another action</Dropdown.Item>
                <Dropdown.Item href="#/action-3">Something else</Dropdown.Item>
              </Dropdown.Menu>
            </Dropdown>
            <Dropdown className="mb-3" onSelect={onSelectGroup}>
              <Dropdown.Toggle
                className="w-75"
                variant="success"
                id="dropdown-basic"
              >
                {selectedGroup}
              </Dropdown.Toggle>
              <Dropdown.Menu>
                <Dropdown.Item href="#/action-1">Action</Dropdown.Item>
                <Dropdown.Item href="#/action-2">Another action</Dropdown.Item>
                <Dropdown.Item href="#/action-3">Something else</Dropdown.Item>
              </Dropdown.Menu>
            </Dropdown>
            <label>Test Cases</label>
            <div className="d-flex form-group mt-1">
              <input
                type="password"
                className="form-control mt-1"
                placeholder="From"
              />
              <p className="ms-1 me-1"> _ </p>
              <input
                type="password"
                className="form-control mt-1"
                placeholder="To"
              />
            </div>
          </Card.Body>
        </Card>
      </div>
    </>
  );
}
