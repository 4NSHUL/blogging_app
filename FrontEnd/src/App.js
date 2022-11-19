import React from "react";
import {BrowserRouter as Router,
  Route,Routes,}
  from "react-router-dom";
import "./App.css";
  
import Posts from "./components/Posts/Posts";
import Navbar from "./components/navbar/navbar";
  
const App = () => {
  return (
    <Router>
      <div className="main-container">
      <Navbar />
      <Routes>
        <Route exact path="/" element={<Posts />} />
      </Routes>
      
    </div>

    </Router>
  );
};

export default App;