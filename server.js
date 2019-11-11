const express = require("express");
const fileUpload = require("express-fileupload");
const fs = require("fs");
const path = require("path");

const app = express();

app.use(fileUpload());

// Image Upload Endpoint
app.post("/upload", (req, res) => {
  if (req.files === null) {
    return res.status(400).json({ message: "No image uploaded" });
  }

  const file = req.files.file;
  const nameSplit = file.name.split(".");
  const imageType = nameSplit[nameSplit.length - 1];
  const fileName = "iamcooldogee";

  const directory = path.join(__dirname, "/client/public/uploads");

  fs.readdir(directory, (err, files) => {
    if (err) throw err;
    if (files.length === 1) {
      for (const file of files) {
        fs.unlink(path.join(directory, file), err => {
          if (err) throw err;
        });
      }
    }
  });

  file.mv(
    `${__dirname}/client/public/uploads/${fileName + "." + imageType}`,
    err => {
      if (err) {
        console.error(err);
        return res.status(500).send(err);
      }

      res.status(200).json({
        fileName: file.name,
        filePath: `/uploads/${fileName + "." + imageType}`
      });
    }
  );
});

app.listen(5000, () => console.log("Server started on port 5000"));
