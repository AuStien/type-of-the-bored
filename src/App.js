import React, { useState, useEffect } from "react";

function App() {
  const [input, setInput] = useState("");
  const [word, setWord] = useState("");
  const [desc, setDesc] = useState("");
  const [nextWord, setNextWord] = useState("");

  useEffect(() => {
    (async () => {
      const w = await getWord();
      setInput("");
      setDesc(w.definition);
      setWord(w.word);
    })();
  }, []);

  useEffect(() => {
    getWord().then((w) => setNextWord(w));
  }, [word]);

  useEffect(() => {
    if (input !== "" && input === word) {
      setInput("");
      setDesc(nextWord.definition);
      setWord(nextWord.word);
    }
  }, [input]);

  const getWord = async () => {
    return new Promise((resolve, reject) => {
      fetch("https://random-words-api.vercel.app/word")
        .then((res) => res.json())
        .then((data) => {
          resolve(data[0]);
        })
        .catch((e) => reject(e));
    });
  };

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
        <p style={{ fontSize: "17px", color: "GrayText" }}>{desc}</p>
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
