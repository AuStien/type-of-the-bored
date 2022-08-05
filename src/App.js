import React, { useState, useEffect } from "react";
import tilfeldigeord from "tilfeldigeord";

function App() {
  const [input, setInput] = useState("");
  const [word, setWord] = useState("");
  const [startTime, setStartTime] = useState(Date.now())
  const [time, setTime] = useState(0)
  const [finishedWords, setFinishedWords] = useState(0)
  const [wpm, setWpm] = useState(0)

  useEffect(() => {
    newWord()
  }, []);

  useEffect(() => {
    if (input !== "" && input === word) {
      newWord()
      setFinishedWords(finishedWords + 2)
    }
  }, [input, word]);

  useEffect(() => {
    setWpm(Math.round(finishedWords/(time/60)))
  }, [finishedWords, time])

  setInterval(() => {
    setTime(Math.round((Date.now() - startTime) / 1000))
  }, 1000)

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
        <p>WPM: {wpm}</p>
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
