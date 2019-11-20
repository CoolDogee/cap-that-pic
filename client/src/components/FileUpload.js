import React, { Fragment, useState } from "react";
import Message from "./Message";
import Loading from "./Loading";
import axios from "axios";
import { API_URL } from "../config";
import Typist from "react-typist";
import Footer from "./Footer";
import Logo from "../images/capthatpic.png";
import { Row, Col } from "react-bootstrap";
import Slider from "rc-slider";
import 'rc-slider/assets/index.css';

export const FileUpload = () => {
  const [message, setMessage] = useState("");
  const [url, setUrl] = useState("");
  const [status, setStatus] = useState(false);
  const [loadingstatus, setLoadingstatus] = useState(false);

  const onChangeUrl = e => {
    setUrl(e.target.value);
    setStatus(true);
  };

  const onChangeLength = e => {
    const valToLength = {1: 1, 50: 2, 100: 3};
    localStorage.setItem("length", valToLength[e.value]);
  };

  const onSubmit = async e => {
    e.preventDefault();
    setLoadingstatus(true);

    try {
      const res = await axios.get(API_URL + "/api/v1/validateImageURL", {
        params: { fileName: url }
      });

      setLoadingstatus(false);
      if (res.status === 200) {
        localStorage.setItem("imageUrl", url);
        console.log(res.data);
        setMessage("Generating captions! Hold tight ...")
        setTimeout(() => {
          window.location = "/choose-caption";
        }, 2500);
      }
    } catch (err) {
      console.log(err.response);
      if (err.response.status === 500) {
        setMessage("Internal Server Error!");
      } else {
        setMessage(err.response.data);
      }
    }
  };

  return (
    <Fragment>
      <h1 className="display-3 text-center mb-4">
        <img src={Logo} style={{ height: "1.5em" }} /> Cap That Pic
    </h1>
      <h4 className="text-center mb-4">
        <Typist> Every Picture Deserves The Perfect Caption</Typist>
      </h4>
      <hr></hr>
      {message ? <Message msg={message} /> : null}
      <div className="custom-file mb-4 mt-5">
        <input
          type="text"
          className="form-control"
          id="exampleFormControlInput1"
          placeholder="Please Enter an Image URL"
          onChange={onChangeUrl}
        ></input>
      </div>
      {setUrl ? (
        <div className="row mt-5">
          <div className="col-md-6 m-auto">
            <img
              style={{ width: "100%" }}
              src={url}
              className="img-fluid"
              alt=""
            />
          </div>
        </div>
      ) : null}
      {loadingstatus ? <Loading /> : null}
      {status ? (
        <div
          className="mt-5 mb-5 text-center"
          style={{ paddingBottom: "100px" }}
        >
          <Row style={{marginBottom: "1em"}}>
            <Col md="12">Choose caption length</Col>
            <Col md="3"></Col>
            <Col md="6">
              <Slider min={1} defaultValue={1} marks={{ 1: "Low", 50: "Medium", 100:  "High"}}
              step={null} handle={onChangeLength}/>
            </Col>
          </Row>
          <br />
          <form onSubmit={onSubmit}>
            <input
              type="submit"
              value="Generate Caption"
              className="btn btn-secondary btn-lg"
            />
          </form>
        </div>
      ) : null}
      <Row style={{ marginTop: "2em", marginBottom: "5em" }}>
        <Col md="4"></Col>
        <Col md="3">
          <span style={{ textAlign: "center" }}><a href="/posts">See sample posts</a></span>
        </Col>
      </Row>
      <Footer />
    </Fragment>
  );
};

export default FileUpload;
