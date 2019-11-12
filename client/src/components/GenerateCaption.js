// import React, { Fragment, useState } from "react";
// import Message from "./Message";
// import axios from "axios";

// export const GenerateCaption = () => {
//   const [filename, setFilename] = useState("");
//   const [caption, setCaption] = useState("");
//   const [message, setMessage] = useState("");

//   const onChange = e => {
//     setFile(e.target.files[0]);
//     setFilename(e.target.files[0].name);
//   };

//   const onSubmit = async e => {
//     e.preventDefault();
//     // const formData = new FormData();
//     // formData.append("file", file);

//     try {
//       const res = await axios.get("/api/v1/getcaption", formData, {
//         headers: {
//           "Content-Type": "multipart/form-data"
//         }
//       });

//       const caption = res.data;
//       if (res.status === 200) {
//         console.log("Generated Caption Successfully");
//         setCaption(caption);
//       }
//     } catch (err) {
//       if (err.response.status === 500) {
//         setMessage("Internal Server Error!");
//       } else {
//         setMessage(err.response.data.message);
//       }
//     }
//   };
//   return (
//     <Fragment>
//       {message ? <Message msg={message} /> : null}
//       <form onSubmit={onSubmit}>
//         <input
//           type="submit"
//           value="Generate Caption"
//           className="btn btn-primary btn-block mt-4"
//         />
//       </form>
//       {caption ? (
//         <div className="row mt-5">
//           <div className="col-md-6 m-auto">
//             <h3 className="text-center">{caption}</h3>
//           </div>
//         </div>
//       ) : null}
//     </Fragment>
//   );
// };

// export default GenerateCaption;
