import React, { Fragment, useState } from "react";
import Message from "./Message";
import Loading from "./Loading";
import axios from "axios";
import { API_URL } from "../config";
import Typist from "react-typist";
import Footer from "./Footer";
import { Container, Row, Col, Card, Button } from "react-bootstrap";
import { FaQuoteLeft, FaQuoteRight, FaCheck } from "react-icons/fa";
import Logo from "../images/capthatpic.png";

function displayHashtags(tags) {
  var tagshash = tags.map(t => '#' + t.replace(" ","_"));

  return tagshash.join(', ');
}
export const CaptionsPage = () => {
  const [message, setMessage] = useState("");
  const [chosenCaption, setCaption] = useState({});
  const [hashtags, setHashtags] = useState([]);
  // TODO
  const [filter, setFilter] = useState("");
  const [generatedCaptions, setGeneratedCaptions] = useState([]);
  const [loading, setLoadingstatus] = useState(false);
  const [initialFetch, setInitialFetch] = useState(false);

  const onChangeCaption = e => {
    setCaption(e);
    setMessage("Caption chosen!");
    console.log(e);
    delayMessage("");
  };

  // temporary data
  const tempData = [
    { "Text": "Beautiful", "Src": "Kayabacho", "Tags": ["story"] },
    { "Text": "Lively", "Src": "link.com", "Tags": ["life", "lessons"] },
    { "Text": "Oustanding", "Src": "home.net", "Tags": ["grammy", "nana", "solo"] },
    { "Text": "Beautiful", "Src": "Kayabacho", "Tags": ["story"] },
    { "Text": "Lively", "Src": "link.com", "Tags": ["life", "lessons"] },
    { "Text": "Oustanding", "Src": "home.net", "Tags": ["grammy", "nana", "solo"] },
    { "Text": "Beautiful", "Src": "Kayabacho", "Tags": ["story"] },
    { "Text": "Lively", "Src": "link.com", "Tags": ["life", "lessons"] },
    { "Text": "Oustanding", "Src": "home.net", "Tags": ["grammy", "nana", "solo"] },
    { "Text": "Beautiful", "Src": "Kayabacho", "Tags": ["story"] },
    { "Text": "Lively", "Src": "link.com", "Tags": ["life", "lessons"] },
    { "Text": "Oustanding", "Src": "home.net", "Tags": ["grammy", "nana", "solo"] },
  ];

  const tempTags = ["cool", "chill"];

  const onSubmit = async e => {
    e.preventDefault();

    var postCaption = chosenCaption;
    postCaption.UserGenerated = true;

    axios.post(API_URL + '/api/v1/caption', {
      Text: [chosenCaption.Text],
      Src : chosenCaption.Src || "",
      Mood: chosenCaption.Mood || [],
      Tags: chosenCaption.Tags || [],
      Type: chosenCaption.Type || "",
      UserGenerated: true 
    })
      .then(function (response) {
        console.log("caption made");
        console.log(response.data.info);
        var captionID = response.data.info;
        // caption created successfully

        axios.post(API_URL + '/api/v1/post', {
          ImgURL    : localStorage.getItem('imageUrl'),
          CaptionID : captionID,
          Filter    : "",
          Tags      : chosenCaption.Tags.concat(hashtags)
        })
          .then(function (resp) {
            // successful post creation
            console.log('Post created');
            console.log(resp.data);
            setMessage('Post created successfully! Redirecting ...');
            setTimeout(() => {
              window.location = "/i/" + resp.data.info;
            }, 2500);
          })
          .catch(function (err) {
            // caption created, but not post
            console.log('Caption created up but error in making post');
            console.log(err.response);
            setMessage('Caption created but error in making post!');
            // console.log(err);
          });
      })
      .catch(function (error) {
        console.log("Could not create caption");
        console.log(error.response);
        setMessage('Could not create caption!');
        // could not create caption
        // message = error.response.data;
        // document.getElementById("alertmsg").innerHTML = message;
        // console.log(error.response.data);
      });
  };

  // wild chance that there's no image
  if (localStorage.getItem("imageUrl") === null) {
    setMessage("Image urt not set! Returning to home page...");
    setTimeout(() => {
      window.location = "/";
    }, 2500);
  }

  // hide message after some time
  const delayMessage = msg => {
    setTimeout(() => {
      setMessage(msg);
    }, 2500);
  }

  // fetch generated captions
  // const getCaptions = async() => {
  const getCaptions = () => {
    setInitialFetch(true);
    try {
      setLoadingstatus(true);
      // const res = await axios.get(API_URL + "/api/v1/getcaption", {
      //   params: { fileName: localStorage.getItem("imageUrl") }
      // });
      var res = {
        status: 200,
        data: { captions: tempData, tags: tempTags },
      };
      setLoadingstatus(false);
      if (res.status === 200) {
        setGeneratedCaptions(res.data.captions);
        setHashtags(res.data.tags);
        console.log(res.data);
        setMessage("Choose a caption that you like!");
        // disappear message after a while
        delayMessage("");
      }
    } catch (err) {
      // console.log(err.response);
      // if (err.response.status === 500) {
      //   setMessage("Internal Server Error!");
      // } else {
      //   setMessage(err.response.data);
      // }
      delayMessage("Server seems to be down. Please try after some time.");
    }
  };

  if (!initialFetch) {
    getCaptions();
  }

  return (
    <Fragment>
      <h1 className="display-3 text-center mb-4">
        <img src={Logo} style={{height: "1.5em"}}/> Cap That Pic
    </h1>
      <h4 className="text-center mb-4">
        <Typist>Beautiful Things Don't Ask For Attention. But They Deserve A Perfect Caption</Typist>
      </h4>
      <hr></hr>
      {message ? <Message msg={message} /> : null}
      <Container style={{ marginBottom: "5em" }}>
        <Row>
          {/* The image, in its full glory */}
          <Col lg="4" md="12">
            <img
              style={{ width: "100%" }}
              src={localStorage.getItem("imageUrl")}
              className="img-fluid"
              alt=""
            />
            <Card>
              <Card.Body>
                <Card.Title>
                  Caption
                </Card.Title>
                <Card.Text>{chosenCaption.Text || ""}</Card.Text>
              </Card.Body>
            </Card>
            {/* Add filter condition here */}
            {Object.keys(chosenCaption).length ?
              <form onSubmit={onSubmit} align="center">
                <input
                  type="submit"
                  value="Create Post!"
                  className="btn btn-primary btn-lg"
                  style={{ marginTop: "1em" }}
                />
              </form>
              : null}
          </Col>

          {/* Filter options, and caption menu */}
          <Col lg="8" md="12">
            {loading ? <Loading /> : null}
            {/* Add filters here */}
            <Row >
              <Col lg="12"><span style={{ fontFamily: "sans-serif", fontWeight: "500" }}>
                Displaying Top 10 captions for your image
                  </span></Col>
              {generatedCaptions.map((caption) =>
                <Col lg="12" style={{ marginTop: "1.5em" }}>
                  <Card>
                    <Card.Body>
                      <Button variant="primary" style={{ float: "right" }} onClick={(e) => {onChangeCaption(caption)}}><FaCheck /></Button>

                      <Card.Title style={{ fontStyle: "italic" }}>
                        <FaQuoteLeft style={{ height: "0.5em", marginTop:"-0.7em"}} />
                        {caption.Text}
                        <FaQuoteRight style={{ height: "0.5em", marginTop: "-0.7em" }} />
                      </Card.Title>
                      <Card.Text className="text-muted">{displayHashtags(caption.Tags)}</Card.Text>
                    </Card.Body>
                  </Card>
                </Col>
              )}
            </Row>

          </Col>
        </Row>
      </Container>
      <Footer />
    </Fragment>
  );
};

export default CaptionsPage;
