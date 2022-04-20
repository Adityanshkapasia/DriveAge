import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Routes, Route, Link } from "react-router-dom";
import Login from './Login';
import Register from './Register';

function App() {
  return (
    <div className="App">
       <div className="App">
      <h1>Welcome to React Router!</h1>
      <Routes>
        <Route path="/Login" element={<Login />} />
        <Route path="/Register" element={<Register />} />
      </Routes>
    </div>
    </div>

  );
}

export default App;
