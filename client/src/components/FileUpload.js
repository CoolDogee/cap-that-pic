import React, { Fragment, useState } from "react";
import Message from "./Message";
import Progress from "./Progress";
import axios from "axios";

export const FileUpload = () => {
  const [message, setMessage] = useState("");
  const [url, setUrl] = useState("");
  const [caption, setCaption] = useState("");

  const onChange = e => {
    console.log(e.target.value);
    setUrl(e.target.value);
  };

  const onSubmit = async e => {
    e.preventDefault();

    try {
      const res = await axios.get("/api/v1/getcaption", {
        params: { fileName: url }
      });

      const caption = res.data;
      if (res.status === 200) {
        console.log("Generated Caption Successfully");
        setCaption(res.data);
      }
    } catch (err) {
      if (err.response.status === 500) {
        setMessage("Internal Server Error!");
      } else {
        setMessage(err.response.data.message);
      }
    }
  };
  //dwdwdw
  return (
    <Fragment>
      <div className="custom-file mb-4">
        <label>Enter image URL</label>
        <input
          type="text"
          class="form-control"
          id="exampleFormControlInput1"
          placeholder="Enter Image URL"
          onChange={onChange}
        ></input>
      </div>
      {setUrl ? (
        <div className="row mt-5">
          <div className="col-md-6 m-auto">
            <img style={{ width: "100%" }} src={url} alt="" />
          </div>
        </div>
      ) : null}
      {setUrl ? <h4>{caption}</h4> : null}
      {setUrl ? (
        <form onSubmit={onSubmit}>
          <input
            type="submit"
            value="Generate Caption"
            className="btn btn-primary btn-block mt-4"
          />
        </form>
      ) : null}
    </Fragment>
  );
};

export default FileUpload;
