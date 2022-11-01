import React from "react";
import "./test.css";

function Test() {
  return (
    <body>
      <header>
        <h1>Franklin Running Club</h1>
      </header>
      <div className="container">
        <main className="main">
          <h2>Come join us!</h2>
          <p>
            The Franklin Running club meets at 6:00pm every Thursday at the town
            square. Runs are three to five miles, at your own pace.
          </p>
        </main>
        <aside className="sidebar">
          <a href="/twitter" className="button-link">
            follow us on Twitter
          </a>
          <a href="/facebook" className="button-link">
            follow us on Facebook
          </a>
        </aside>
      </div>
    </body>
  );
}

export default Test;
