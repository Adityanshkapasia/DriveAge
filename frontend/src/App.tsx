import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Routes, Route, Link, useNavigate } from "react-router-dom";
import Login from './Login';
import Register from './Register';
import Viewallblog from './Viewallblog';
import Createblogpost from './Createblogpost';
import Viewpost from './Viewpost';
import Viewallcars from './Viewallcars';
import Comparecars from './Comparecars';
import News from './News';

function App() {
  const navigate=useNavigate()
  const onLogout=()=>{
    fetch("http://localhost:8080/logout", {
        credentials : 'include' ,
    }) .then (res=>{
      navigate("/Login")
        })
        
    
      }
  return (
    <div className="App">
       <div className="App">
      <h1>Welcome to React Router!</h1>
      <nav
        style={{
          borderBottom: "solid 1px",
          paddingBottom: "1rem",
        }}
      >
        <Link to="/Blog">View All Posts</Link> 
        <Link to="/CreatePost">Create Post</Link>
        <Link to="/cars">View All Cars</Link> 
        <Link to="/compare">Compare Cars</Link> 
        <Link to="/news">News</Link> 
        <a onClick={e => onLogout()}>logout</a>
      </nav>
      <Routes>
        <Route path="/Login" element={<Login />} />
        <Route path="/Register" element={<Register />} />
        <Route path="/Blog" element={<Viewallblog />} />
        <Route path="/CreatePost" element={<Createblogpost />} />
        <Route path="/ViewPost/:id" element={<Viewpost />} />
        <Route path="/cars" element={<Viewallcars />} />
        <Route path="/compare" element={<Comparecars />} />
        <Route path="/news" element={<News />} />
      </Routes>
    </div>
    </div>

  );
}

export default App;
