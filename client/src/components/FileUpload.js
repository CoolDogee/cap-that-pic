import React, { Fragment, useState } from "react";
import Message from "./Message";
import axios from "axios";
// import Typist from "react-typist";

export const FileUpload = () => {
  const [message, setMessage] = useState("");
  const [url, setUrl] = useState("");
  const [caption, setCaption] = useState([]);
  const [status, setStatus] = useState(false);

  const onChange = e => {
    setUrl(e.target.value);
    setStatus(true);
  };

  const onSubmit = async e => {
    e.preventDefault();

    try {
      const res = await axios.get("/api/v1/getcaption", {
        params: { fileName: url }
      });

      if (res.status === 200) {
        const capLines = res.data.split("\n");
        setCaption(capLines);
      }
    } catch (err) {
      if (err.response.status === 500) {
        setMessage("Internal Server Error!");
      } else {
        setMessage(err.response.data.message);
      }
    }
  };

  return (
    <Fragment>
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
      {setUrl ? (
        <div className="row mt-5 mb-5">
          <div className="col-md-8 m-auto text-center">
            <h4>{caption[0]}</h4>
            <h4>{caption[1]}</h4>
            <h4>{caption[2]}</h4>
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

export default FileUpload;
