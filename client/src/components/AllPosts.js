import React, { Fragment, useState } from "react";
import Message from "./Message";
import Loading from "./Loading";
import axios from "axios";
import { API_URL } from "../config";
import Typist from "react-typist";

export const AllPostsPage = () => {
  const [message, setMessage] = useState("");
  const [url, setUrl] = useState("");
  const [captions, setCaption] = useState([]);
  const [status, setStatus] = useState(false);
  const [secstatus, setSecStatus] = useState(false);
  const [tristatus, setTriStatus] = useState(false);

  const onChange = e => {
    setUrl(e.target.value);
    setStatus(true);
    setTriStatus(false);
  };

  const onSubmit = async e => {
    e.preventDefault();
    setSecStatus(true);

    try {
      const res = await axios.get("http://google.com", {//API_URL + "/api/v1/getcaption", {
        params: { fileName: url }
      });

      if (res.status === 200) {
        // const capLines = res.data.split("\n");
        // setCaption(capLines);
        // setSecStatus(false);
        // setTriStatus(true);
        localStorage.setItem("imageUrl", this.state.url);
      }
    } catch (err) {
      if (err.response.status === 500) {
        console.log(err.response);
        setMessage("Internal Server Error!");
      } else {
        setMessage(err.response.data.message);
      }
    }
  };

  return (
    <Fragment>
      <h1 className="display-3 text-center mb-4">
        <i className="fas fa-music"></i> Cap That Pic
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
          onChange={onChange}
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
      {secstatus ? <Loading /> : null}
      {tristatus ? (
        <div className="row mt-5 mb-5">
          <div className="col-md-8 m-auto text-center">
            <h4>{captions[0]}</h4>
            <h4>{captions[1]}</h4>
            <h4>{captions[2]}</h4>
          </div>
        </div>
      ) : null}
      {status ? (
        <div
          className="mt-5 mb-5 text-center"
          style={{ paddingBottom: "100px" }}
        >
          <form onSubmit={onSubmit}>
            <input
              type="submit"
              value="Generate Caption"
              className="btn btn-secondary btn-lg"
            />
          </form>
        </div>
      ) : null}
    </Fragment>
  );
};

export default AllPostsPage;
