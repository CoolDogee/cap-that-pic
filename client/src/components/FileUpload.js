import React, { Fragment, useState } from "react";
import Message from "./Message";
import Progress from "./Progress";
import axios from "axios";

export const FileUpload = () => {
  const [file, setFile] = useState("");
  const [filename, setFilename] = useState("Choose Image");
  const [fileType, setFileType] = useState("");
  const [uploadedFile, setUploadedFile] = useState({});
  const [message, setMessage] = useState("");
  const [uploadPercentage, setUploadPercentage] = useState(0);
  const [caption, setCaption] = useState("");

  const onChange = e => {
    setFile(e.target.files[0]);
    setFilename(e.target.files[0].name);
  };

  const onSubmit = async e => {
    e.preventDefault();
    const formData = new FormData();
    formData.append("file", file);

    // TODO: Check of the uploaded file is a image. Aloow only .png or .jpeg

    try {
      const res = await axios.post("/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data"
        },
        onUploadProgress: progressEvent => {
          setUploadPercentage(
            parseInt(
              Math.round((progressEvent.loaded * 100) / progressEvent.total)
            )
          );
          // Clear Percentage
          setTimeout(() => setUploadPercentage(0), 10000);
        }
      });

      const { fileName, filePath } = res.data;
      if (res.status === 200) {
        console.log("Image successfully uploaded");
        setUploadedFile({ fileName, filePath });
        setMessage("File Successfully Uploaded!");
      }
    } catch (err) {
      if (err.response.status === 500) {
        setMessage("Internal Server Error!");
      } else {
        setMessage(err.response.data.message);
      }
    }
  };

  const onSubmit2 = async e => {
    e.preventDefault();
    const nameSplit = uploadedFile.fileName.split(".");
    const imageType = nameSplit[nameSplit.length - 1];

    try {
      const res = await axios.get("/api/v1/getcaption", {
        params: { fileName: "iamcooldogee" + "." + imageType }
      });

      const caption = res.data;
      if (res.status === 200) {
        console.log("Generated Caption Successfully");
        setCaption(caption);
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
      <form onSubmit={onSubmit}>
        <div className="custom-file mb-4">
          <input
            type="file"
            className="custom-file-input"
            id="customFile"
            onChange={onChange}
          />
          <label className="custom-file-label" htmlFor="customFile">
            {filename}
          </label>
        </div>

        <Progress percentage={uploadPercentage} />

        <input
          type="submit"
          value="Upload"
          className="btn btn-primary btn-block mt-4"
        />
      </form>
      {uploadedFile ? (
        <div className="row mt-5">
          <div className="col-md-6 m-auto">
            <h3 className="text-center">{uploadedFile.fileName}</h3>
            <img style={{ width: "100%" }} src={uploadedFile.filePath} alt="" />
          </div>
        </div>
      ) : null}
      {caption ? (
        <div className="row mt-5">
          <div className="col-md-6 m-auto">
            <h3 className="text-center">{caption}</h3>
          </div>
        </div>
      ) : null}
      {uploadedFile ? (
        <form onSubmit={onSubmit2}>
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
