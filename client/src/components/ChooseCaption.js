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
import ImageFilter from "react-image-filter";
// https://muffinman.io/react-image-filter/

function displayHashtags(tags) {
  if (!tags || !tags.length) {
    return '';
  }
  var tagshash = tags.map(t => '#' + t.split(" ").join("_"));

  return tagshash.join(', ');
}

const NONE = [
  1, 0, 0, 0, 0,
  0, 1, 0, 0, 0,
  0, 0, 1, 0, 0,
  0, 0, 0, 1, 0,
];

export const CaptionsPage = () => {
  const [message, setMessage] = useState("");
  const [chosenCaption, setCaption] = useState({});
  const [hashtags, setHashtags] = useState([]);
  // TODO
  const [values, setValues] = useState([...NONE]);
  const [filter, setFilter] = useState(values);
  const [applyFilter, setApplyFilter] = useState(true);
  const [colorOne, setColorOne] = useState(null);
  const [colorTwo, setColorTwo] = useState(null);
  const [generatedCaptions, setGeneratedCaptions] = useState([]);
  const [loading, setLoadingstatus] = useState(true);
  const [initialFetch, setInitialFetch] = useState(false);

  const onChangeCaption = e => {
    setCaption(e);
    setMessage("Caption chosen!");
    console.log(e);
    delayMessage("");
  };

  const getFilterString = (fil, colOne, colTwo) => {
    if(colOne === [250, 50, 50] && colTwo === [20, 20, 100]) {
      return "Duotone (red / blue)";
    } else if(colOne === [50, 250, 50] && colTwo === [250, 20, 220]) {
      return "Duotone (green / purple)";
    } else if(colOne === [40, 250, 250] && colTwo === [250, 150, 30]) {
      return "Duotone (light blue/orange)";
    } else if (colOne === [40, 70, 200] && colTwo ===  [220, 30, 70]) {
      return "Duotone (blue / red)";
    }
    return typeof fil === 'string' ? fil : 'none';
  }

  const onSubmit = async e => {
    e.preventDefault();
    console.log(chosenCaption);

    axios.post(API_URL + '/api/v1/caption', chosenCaption)
      .then(function (response) {
        console.log("caption made");
        console.log(response.data.info);
        var captionID = response.data.info;
        // caption created successfully

        console.log(getFilterString(filter, colorOne, colorTwo));
        axios.post(API_URL + '/api/v1/post', {
          ImgURL: localStorage.getItem('imageUrl'),
          CaptionID: captionID,
          Filter: getFilterString(filter, colorOne, colorTwo),
          Tags: chosenCaption.Tags.concat(hashtags)
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
    setLoadingstatus(true);
    axios.get(API_URL + '/api/v1/getcaption', {
      params: { fileName: localStorage.getItem('imageUrl'), length: 1 }
    })
      .then(function (response) {
        console.log('Captions found');
        console.log(response.data);
        setGeneratedCaptions(response.data);

      })
      .catch(function (error) {
        setMessage("No captions could be generated!");
        console.log(error.response);
      });
    setLoadingstatus(false);
  };

  if (!initialFetch) {
    getCaptions();
  }

  // reflect change in filter
  const ButtonFilter = (e, fil, val, colOne, colTwo) => {
    e.preventDefault();
    setApplyFilter(true); setFilter(fil);
    if (val) {
      setValues(val);
    }
    if (colOne) {
      setColorOne(colOne);
    }
    if (colTwo) {
      setColorTwo(colTwo);
    }
  }

  return (
    <Fragment>
      <h1 className="display-3 text-center mb-4">
        <img src={Logo} style={{ height: "1.5em" }} /> Cap That Pic
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
            <ImageFilter image={localStorage.getItem("imageUrl")}
              key="keykey" filter={applyFilter ? filter : NONE}
              colorOne={colorOne} colorTwo={colorTwo}
              onChange={(m) => { setValues(m) }}
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
          <Col lg="1"></Col>
          <Col lg="7" md="12">
            {/* Add filters here */}
            <Row >
              <Col lg="12">
                Choose filters
                <Row style={{ marginTop: "0.3em" }}>
                  <Col md="2"><Button variant="secondary"
                  onClick={(e) => { ButtonFilter(e, NONE, NONE, null, null)}}>None
                  </Button></Col>

                  <Col md="2"><Button variant="secondary"
                    onClick={(e) => { ButtonFilter(e, 'invert', null, null, null)}}> Invert
                  </Button></Col>

                  <Col md="2"><Button variant="secondary"
                    onClick={(e) => { ButtonFilter(e, 'grayscale', null, null, null)}}> Grayscale
                  </Button></Col>
                
                  <Col md="2"><Button variant="secondary"
                    onClick={(e) => { ButtonFilter(e, 'sepia', null, null, null)}}> Sepia
                  </Button></Col>
                  
                  <Col md="4"><Button variant="secondary"
                    onClick={(e) => { ButtonFilter(e, 'duotone', null, [250, 50, 50], [20, 20, 100])}}> Duotone (red / blue)
                  </Button></Col>
                </Row>
                <Row style={{marginTop: "1em"}}>
                  <Col md="4"><Button variant="secondary"
                    onClick={(e) => { ButtonFilter(e, 'duotone', null, [50, 250, 50], [250, 20, 220])}}> Duotone (green / purple)
                  </Button></Col>
                  
                  <Col md="4"><Button variant="secondary"
                    onClick={(e) => { ButtonFilter(e, 'duotone', null, [40, 250, 250], [250, 150, 30])}}> Duotone (light blue/orange)
                  </Button></Col>

                  <Col md="4"><Button variant="secondary"
                    onClick={(e) => { ButtonFilter(e, 'duotone', null, [40, 70, 200], [220, 30, 70])}}> Duotone (blue / red)
                  </Button></Col>

                </Row>
              </Col>

              <Col lg="12" style={{ marginTop: "2em" }}>Displaying Top 10 captions for your image</Col>
              {Object.keys(generatedCaptions).length || loading ?
                generatedCaptions.map((caption) =>
                  <Col lg="12" style={{ marginTop: "1.5em" }}>
                    <Card>
                      <Card.Body>
                        <Button variant="primary" style={{ float: "right" }} onClick={(e) => { onChangeCaption(caption) }}><FaCheck /></Button>

                        <Card.Title style={{ fontStyle: "italic" }}>
                          <FaQuoteLeft style={{ height: "0.5em", marginTop: "-0.7em" }} />
                          {caption.Text}
                          <FaQuoteRight style={{ height: "0.5em", marginTop: "-0.7em" }} />
                        </Card.Title>
                        <Card.Text className="text-muted">{displayHashtags(caption.Tags)}</Card.Text>
                      </Card.Body>
                    </Card>
                  </Col>
                ) : <Loading />}
            </Row>

          </Col>
        </Row>
      </Container>
      <Footer />
    </Fragment>
  );
};

export default CaptionsPage;
