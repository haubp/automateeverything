import React, { useState } from "react";
import Col from "react-bootstrap/Col";
import Nav from "react-bootstrap/Nav";
import Row from "react-bootstrap/Row";
import Tab from "react-bootstrap/Tab";
import Button from "react-bootstrap/Button";
import NodeConfig from "./NodeConfig";
import NodeReport from "./NodeReport";
import { Icon1Circle, Check } from "react-bootstrap-icons";
import Spinner from "react-bootstrap/Spinner";
import LoadingButton from "./LoadingButton";
import "./Portal.css";

function Portal() {
  let [isRunning, setRunning] = useState(false);
  let toggleRun = function () {};
  return (
    <Tab.Container id="left-tabs-example" defaultActiveKey="config">
      <Row sm={1}>
        <div className="banner">
          <p>Hello</p>
          <a className="user-setting" href="/logout">
            ID
          </a>
        </div>
      </Row>
      <Row sm={10} className="left-pannel-controller">
        <Col sm={3} className="rounded controller">
          <Nav variant="pills" className="flex-column pt-1">
            <Nav.Item>
              <Nav.Link
                eventKey="config"
                className="border node-config-control"
              >
                Node Config <Icon1Circle />
              </Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link eventKey="report" className="border reporter-control">
                Reporter
                {isRunning ? (
                  <Spinner animation="border" role="status" size="sm">
                    <span className="visually-hidden">Loading...</span>
                  </Spinner>
                ) : (
                  <div></div>
                )}
                <Spinner animation="border" role="status" size="sm">
                  <span className="visually-hidden">Loading...</span>
                </Spinner>
              </Nav.Link>
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
          <LoadingButton text="Run" className="run-button" />
          <LoadingButton text="Cancel" className="cancel-button" />
        </Col>
      </Row>
    </Tab.Container>
  );
}

export default Portal;
