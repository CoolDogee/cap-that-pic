import React from "react";
import FileUpload from "./components/FileUpload";
import "./App.css";

const App = () => (
  <div className="container mt-4">
    <h4 className="display-4 text-center mb-4">
      <i className="fas fa-music"></i> Cap That Pic
    </h4>
    <FileUpload />
  </div>
);

export default App;
