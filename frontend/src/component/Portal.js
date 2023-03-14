import React from "react";
import Col from "react-bootstrap/Col";
import Nav from "react-bootstrap/Nav";
import Row from "react-bootstrap/Row";
import Tab from "react-bootstrap/Tab";
import Button from "react-bootstrap/Button";
import NodeConfig from "./NodeConfig";
import NodeReport from "./NodeReport";
import "./Portal.css";

function Portal() {
  return (
    <Tab.Container id="left-tabs-example" defaultActiveKey="config">
      <Row sm={1}>
        <p>User ID</p>
      </Row>
      <Row sm={10} className="left-pannel-controller">
        <Col sm={3}>
          <Nav variant="pills" className="flex-column">
            <Nav.Item>
              <Nav.Link eventKey="config">Config</Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link eventKey="report">Report</Nav.Link>
            </Nav.Item>
          </Nav>
        </Col>
        <Col sm={7}>
          <Tab.Content>
            <Tab.Pane eventKey="config">
              <NodeConfig />
            </Tab.Pane>
            <Tab.Pane eventKey="report">
              <NodeReport />
            </Tab.Pane>
          </Tab.Content>
        </Col>
        <Col sm={1}>
          <Button variant="outline-success">Run</Button>{" "}
        </Col>
      </Row>
    </Tab.Container>
  );
}

export default Portal;
