import React, { useState } from "react";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import axios from "axios";
import Spinner from "react-bootstrap/Spinner";
import { ArrowClockwise } from "react-bootstrap-icons";
import "./NodeReport.css";

export default function () {
  let [isRunning, setRunning] = useState(true);
  let [passed, setPassed] = useState(0);
  let [failed, setFailed] = useState(0);
  const baseUrl = "http://10.40.50.34:6868";
  const downloadHandler = (e) => {
    axios({
      url:
        baseUrl +
        "/public/user_d82494f05d6917ba02f7aaa29689ccb444bb73f20380876cb05d1f37537b7892/test_result.json", //your url
      method: "GET",
      responseType: "blob", // important
    }).then((response) => {
      // create file link in browser's memory
      const href = URL.createObjectURL(response.data);

      // create "a" HTML element with href to file & click
      const link = document.createElement("a");
      link.href = href;
      link.setAttribute("download", "test_result.json"); //or any other extension
      document.body.appendChild(link);
      link.click();

      // clean up "a" element & remove ObjectURL
      document.body.removeChild(link);
      URL.revokeObjectURL(href);
    });
  };

  const reloadHandler = async () => {
    axios({
      url:
        baseUrl +
        "/public/user_d82494f05d6917ba02f7aaa29689ccb444bb73f20380876cb05d1f37537b7892/test_result.json", //your url
      method: "GET",
    }).then(res => res.data).then((data) => {
      setRunning(false);
      console.log(data);
      let pass = 0;
      let fail = 0;
      for (let groupIndex in data.test_category_groups) {
        let group = data.test_category_groups[groupIndex];
        for (let testCaseIndex in group.test_group_tests) {
          let testCase = group.test_group_tests[testCaseIndex]
          if (testCase.result == "Success") {
            pass += 1;
          } else {
            fail += 1;
          }
        }
      }

      setPassed(pass);
      setFailed(fail);
    });
  }

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
              <ArrowClockwise onClick={reloadHandler}/>
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
                style={{ width: "3rem", color: "black" }}
              >
                { passed }
              </span>
            </label>
          </div>
          <div
            className="mt-4"
            style={{ display: "flex", justifyContent: "space-between" }}
          >
            <label>
              Fail
              <span
                className="badge badge-primary text-warap"
                style={{ width: "3rem", color: "black" }}
              >
                { failed }
              </span>
            </label>
          </div>
          <div className="report-download">
            {isRunning ? (
              <Spinner animation="border" role="status" size="sm">
                <span className="visually-hidden">Loading...</span>
              </Spinner>
            ) : (
              <Button className="result-btn" variant={failed != 0 ? "outline-warning" : "outline-success"} onClick={downloadHandler}>
                Download
              </Button>
            )}
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
