// import React from "react";
import ReactDOM from "react-dom";
import React, { useState } from "react";
import ReactTooltip from "react-tooltip";

import "./App.css";
import data from "./data.json";
// import Navbar from "./navbar.js"

import MapChart from "./MapChart";

function App() {
  const [content, setContent] = useState("");
  return (
    <div>
      {/* <Navbar/> */}
      <div>
        {
          data.map(data_row => {
              <h4>
                {data_row.test_key}
              </h4>
          })}
      </div>
      <MapChart setTooltipContent={setContent}/>
      <ReactTooltip className={"tooltip"}>{content}</ReactTooltip>
    </div>
  );
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);
