import React, { Fragment, useState } from "react";
import Message from "./Message";
import Progress from "./Progress";
import axios from "axios";

export const FileUpload = () => {
  // const [file, setFile] = useState("");
  // const [filename, setFilename] = useState("Choose Image");
  // const [uploadedFile, setUploadedFile] = useState({});
  // const [message, setMessage] = useState("");
  // const [uploadPercentage, setUploadPercentage] = useState(0);
  const [url, setUrl] = useState("");

  const onChange = e => {
    // setFile(e.target.files[0]);
    // setFilename(e.target.files[0].name);
    // console.log(e.target.value);
    setUrl(e.target.value);
  };

  // const onSubmit = async e => {
  //   e.preventDefault();
  //   const formData = new FormData();
  //   formData.append("file", file);

  //   // TODO: Check of the uploaded file is a image. Aloow only .png or .jpeg

  //   try {
  //     const res = await axios.post("/upload", formData, {
  //       headers: {
  //         "Content-Type": "multipart/form-data"
  //       },
  //       onUploadProgress: progressEvent => {
  //         setUploadPercentage(
  //           parseInt(
  //             Math.round((progressEvent.loaded * 100) / progressEvent.total)
  //           )
  //         );
  //         // Clear Percentage
  //         setTimeout(() => setUploadPercentage(0), 10000);
  //       }
  //     });

  //     const { fileName, filePath } = res.data;
  //     if (res.status === 200) {
  //       console.log("Image successfully uploaded");
  //       setUploadedFile({ fileName, filePath });
  //       console.log(uploadedFile.filePath);
  //       setMessage("File Successfully Uploaded!");
  //     }
  //   } catch (err) {
  //     if (err.response.status === 500) {
  //       setMessage("Internal Server Error!");
  //     } else {
  //       setMessage(err.response.data.message);
  //     }
  //   }
  // };
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
    </Fragment>
  );
};

export default FileUpload;
