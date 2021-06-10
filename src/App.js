import React, { useState, useEffect } from 'react';

function App() {
  const [input, setInput] = useState("")
  const [word, setWord] = useState("")

  useEffect(() => {
    getWord()
  }, [])

  useEffect(() => {
    if (input !== "" && input === word) {
      getWord()
    }
  }, [input])
  
  const getWord = () => {
    fetch("https://random-words-api.vercel.app/word")
    .then(res => res.json())
    .then(data => {
      setWord(data[0].word)
      setInput("")
    })
  }

  const handle = (l, i) => {
    if (i >= input.length) {
      return <span style={{color: 'white'}} >{l}</span>
    }else if (l === input[i]) {
      return <span style={{color: 'green'}} >{l}</span>
    }else {
      return <span style={{color: 'red'}} >{l}</span>
    }
  }

  return (
    <div className="App" style={{position: "fixed", top: "35%", left: "40%", width: "100%", display: "flex", justifyContent: "center", flexDirection: "column", }}>
      <div style={{width: "200px"}}>
        <p style={{textAlign: "left"}}>{word.split("").map(handle)}</p>
        <input value={input} onChange={e => setInput(e.target.value)} style={{fontSize: "xx-large"}} />    
      </div>
    </div>
  );
}

export default App;
