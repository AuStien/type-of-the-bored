import React, { useState, useEffect } from "react";
import tilfeldigeord from "tilfeldigeord";

function App() {
  const [input, setInput] = useState("");
  const [word, setWord] = useState("");

  useEffect(() => {
    newWord()
  }, []);

  useEffect(() => {
    if (input !== "" && input === word) {
      newWord()
    }
  }, [input, word]);

  const newWord = () => {
    setInput("");
    setWord(tilfeldigeord.getTilfeldigOrd());
  }

  const handle = (l, i) => {
    if (i >= input.length) {
      return (
        <span style={{ color: "white" }} key={i}>
          {l}
        </span>
      );
    } else if (l === input[i]) {
      return (
        <span style={{ color: "green" }} key={i}>
          {l}
        </span>
      );
    } else {
      return (
        <span style={{ color: "red" }} key={i}>
          {l}
        </span>
      );
    }
  };

  return (
    <div
      className="App"
      style={{
        position: "fixed",
        top: "35%",
        left: "40%",
        width: "100%",
        display: "flex",
        justifyContent: "center",
        flexDirection: "column",
      }}
    >
      <div style={{ width: "400px" }}>
        <p style={{ textAlign: "left" }}>{word?.split("").map(handle)}</p>
        <input
          value={input}
          onChange={(e) => setInput(e.target.value)}
          style={{ fontSize: "xx-large" }}
        />
      </div>
    </div>
  );
}

export default App;
