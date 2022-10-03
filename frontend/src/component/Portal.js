import React from 'react';
import Col from 'react-bootstrap/Col';
import Nav from 'react-bootstrap/Nav';
import Row from 'react-bootstrap/Row';
import Tab from 'react-bootstrap/Tab';
import NodeConfig from './NodeConfig';
import NodeReport from './NodeReport';
import './Portal.css';

function Portal() {
  return (
    <Tab.Container id="left-tabs-example" defaultActiveKey="config">
      <Row>
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
        <Col sm={9}>
          <Tab.Content>
            <Tab.Pane eventKey="config">
              <NodeConfig />
            </Tab.Pane>
            <Tab.Pane eventKey="report">
              <NodeReport />
            </Tab.Pane>
          </Tab.Content>
        </Col>
      </Row>
    </Tab.Container>
  );
}

export default Portal;