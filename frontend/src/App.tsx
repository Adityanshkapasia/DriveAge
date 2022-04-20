import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Routes, Route, Link } from "react-router-dom";
import Login from './Login';
import Register from './Register';
import Viewallblog from './Viewallblog';
import Createblogpost from './Createblogpost';
import Viewpost from './Viewpost';
import Viewallcars from './Viewallcars';

function App() {
  return (
    <div className="App">
       <div className="App">
      <h1>Welcome to React Router!</h1>
      <Routes>
        <Route path="/Login" element={<Login />} />
        <Route path="/Register" element={<Register />} />
        <Route path="/Blog" element={<Viewallblog />} />
        <Route path="/CreatePost" element={<Createblogpost />} />
        <Route path="/ViewPost/:id" element={<Viewpost />} />
        <Route path="/cars" element={<Viewallcars />} />
      </Routes>
    </div>
    </div>

  );
}

export default App;
