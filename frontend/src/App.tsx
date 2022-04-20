import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Routes, Route, Link } from "react-router-dom";
import Login from './Login';

function App() {
  return (
    <div className="App">
       <div className="App">
      <h1>Welcome to React Router!</h1>
      <Routes>
        <Route path="/Login" element={<Login />} />
        {/* <Route path="about" element={<About />} /> */}
      </Routes>
    </div>
    </div>

  );
}

export default App;
