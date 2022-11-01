import React, { useEffect, useState } from "react";
import Button from "react-bootstrap/Button";
import "./LoadingButton.css";

function simulateNetworkRequest() {
  return new Promise((resolve) => setTimeout(resolve, 2000));
}

function LoadingButton(props) {
  const [isLoading, setLoading] = useState(false);

  useEffect(() => {
    if (isLoading) {
      simulateNetworkRequest().then(() => {
        setLoading(false);
      });
    }
  }, [isLoading]);

  const handleClick = () => setLoading(true);

  return (
    <Button
      className="loading-button"
      disabled={isLoading}
      onClick={!isLoading ? handleClick : null}
    >
      {isLoading ? props.text + "…" : props.text}
    </Button>
  );
}

export default LoadingButton;
