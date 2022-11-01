import React, { useState, useRef } from "react";
import Col from "react-bootstrap/Col";
import Nav from "react-bootstrap/Nav";
import Row from "react-bootstrap/Row";
import Tab from "react-bootstrap/Tab";
import Button from "react-bootstrap/Button";
import NodeConfig from "./NodeConfig";
import NodeReport from "./NodeReport";
import { Icon1Circle, Check } from "react-bootstrap-icons";
import Spinner from "react-bootstrap/Spinner";
import axios from "axios";
import "./Portal.css";

function Portal() {
  let [isRunning, setRunning] = useState(false);
  const [selectedFile, setSelectedFile] = useState("");

  const changeHandler = (event) => {
    setSelectedFile(event.target.files[0]);
  };

  const baseUrl = "http://10.40.50.34:6868";

  const onUploadClick = (event) => {
    const reader = new FileReader();
    reader.onload = async (e) => {
      const templateJSON = JSON.parse(e.target.result);
      await axios({
        method: "post",
        url:
          baseUrl +
          "/api/v1/user/template?user_id=d82494f05d6917ba02f7aaa29689ccb444bb73f20380876cb05d1f37537b7892",
        headers: {
        },
        data: templateJSON,
      });
      alert("Test Template is being uploaded!");
      setRunning(true);
    };
    reader.readAsText(selectedFile);
  };

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
          <Nav variant="pills" className="flex-column">
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
        <Col sm={2}>
          <div className="border p-3 upload-card">
            <input
              type="file"
              id="file"
              onChange={changeHandler}
              style={{ top: "50%" }}
            />
            <Button
              className="upload-btn"
              variant="outline-success"
              onClick={onUploadClick}
            >
              Upload
            </Button>
          </div>
        </Col>
      </Row>
    </Tab.Container>
  );
}

export default Portal;
