import React from "react";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import axios from "axios";
import "./NodeReport.css";

export default function () {
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
                style={{ width: "3rem", color: "black" }}
              >
                1
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
                0
              </span>
            </label>
          </div>
          <div className="report-download">
            <Button className="result-btn" variant="outline-warning">
              Download
            </Button>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
