import React from "react";
import FileUpload from "./components/FileUpload";
import Footer from "./components/Footer";
import Typist from "react-typist";
import "./App.css";

const App = () => (
  <div className="container mt-4">
    <h1 className="display-3 text-center mb-4">
      <i className="fas fa-music"></i> Cap That Pic
    </h1>
    <h4 className="text-center mb-4">
      <Typist> Every Picture Deserves The Perfect Caption</Typist>
    </h4>
    <hr></hr>
    <FileUpload />
    <Footer />
  </div>
);

export default App;
