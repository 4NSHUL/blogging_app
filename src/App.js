import React from "react";
import "./App.css";
  
import Posts from "./components/Posts/Posts";
  
const App = () => {
  return (
    <div className="main-container">
      <h1 className="main-heading" >
        Blog2
      </h1>
      <img className="bulb" src="https://img.freepik.com/premium-photo/woman-works-office-blue-background-concept-workspace-working-computer-freelance-banner_164357-1144.jpg?w=2000" alt="!!!" />
      <div className="post-container"><Posts /></div>
    </div>
    
  );
};

export default App;