import React, { useState, useEffect } from "react";
import logo from "./logo.svg";
import "./App.css";
import { hello } from "./services/helloService";

function App() {
  const [message, setMessage] = useState("");

  const handleHello = async () => {
    try {
      const response = await hello();
      setMessage(`Successful ping with response: ${await response.text()}`);
    } catch (e) {
      setMessage("Failed to ping");
    }
  };

  useEffect(() => {
    handleHello();
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>{message}</p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
