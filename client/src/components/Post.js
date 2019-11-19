import React, { Fragment, useState } from "react";
import Message from "./Message";
import Loading from "./Loading";
import axios from "axios";
import { API_URL } from "../config";
import Footer from "./Footer";
import { Container, Row, Col } from "react-bootstrap";
import { FaQuoteLeft, FaQuoteRight } from "react-icons/fa";
import {
  FacebookShareButton, FacebookIcon,
  TwitterShareButton, TwitterIcon,
  WhatsappShareButton, WhatsappIcon,
  EmailShareButton, EmailIcon,
  RedditShareButton, RedditIcon,
} from 'react-share';
import Logo from "../images/capthatpic.png";

function displayHashtags(tags) {
  if (!tags || !tags.length) {
    return '';
  }
  var tagshash = tags.map(t => '#' + t.replace(" ", "_"));

  return tagshash.join(', ');
}

export const PostPage = () => {
  const [message, setMessage] = useState("");
  const [caption, setCaption] = useState({});
  const [post, setPost] = useState({});
  const [loading, setLoadingstatus] = useState(false);
  const [initialFetch, setInitialFetch] = useState(false);
  // unset imageurl
  localStorage.removeItem('imageUrl');

  const getPost = () => {
    setInitialFetch(true);
    setLoadingstatus(true);
    const postID_href = window.location.href.split("/").slice(-1)[0];
    axios.get(API_URL + '/api/v1/post/' + postID_href)
      .then(function (response) {
        console.log('Post found successfully');
        console.log(response.data.info);
        setPost(response.data.info);

        // try to retrieve caption now
        axios.get(API_URL + '/api/v1/caption/' + response.data.info.CaptionID)
          .then(function (resp) {
            console.log('Caption retrieved');
            console.log(resp.data);
            setCaption(resp.data.info);
            setLoadingstatus(false);

          })
          .catch(function (err) {
            setMessage("Caption could not be retrieved :/ Please try reloading")
            setLoadingstatus(false);
            console.log(err.response);
          });

      })
      .catch(function (error) {
        setMessage("No such post found! Are you sure you entered the correct url ?")
        setLoadingstatus(false);
        console.log(error.response);
      });
  };

  if (!initialFetch) {
    getPost();
  }

  return (
    <Fragment>
      <h1 className="display-3 text-center mb-4">
        <img src={Logo} style={{ height: "1.5em" }} /> Cap That Pic
    </h1>
      <hr></hr>
      {message ? <Message msg={message} /> : null}
      <Container style={{ marginBottom: "5em" }}>
        <Row>
          <Col md="12">{loading ? <Loading /> : null}</Col>
          {/* The image, in its full glory */}
          <Col lg="4">
            <img
              style={{ maxwidth: "100%" }}
              src={post.ImgURL || "https://bitsofco.de/content/images/2018/12/broken-1.png"}
              className="img-fluid" alt=""
            />
          </Col>
          <Col lg="1"></Col>
          <Col lg="6" md="12">
            <Row style={{ marginTop: '3em' }}></Row>
            <Row style={{ fontStyle: "italic", fontSize: "3em", align: "center" }} className="text-muted">
              <FaQuoteLeft style={{ height: "0.5em" }} />
              {caption.Text || "caption"}
              <FaQuoteRight style={{ height: "0.5em" }} />
            </Row>

            <Row style={{ marginTop: '3em' }}>
              Suggested Hashtags: &nbsp;&nbsp;<span style={{ color: 'blue' }}>{displayHashtags(post.Tags) || "#tag"}</span>
            </Row>
            {Object.keys(post).length ?
              <Row style={{ marginTop: '3em' }}>
                Share on social media: &nbsp; &nbsp;
              <br />
                <FacebookShareButton url={window.location}
                  quote={caption.Text ? caption.Text[0] : ''} hashtag={displayHashtags(post.Tags)}>
                  <FacebookIcon
                    size={32}
                    round />
                </FacebookShareButton>
                <TwitterShareButton
                  url={window.location}
                  title={caption.Text ? caption.Text[0] : ''}
                  style={{ marginLeft: "1em" }}>
                  <TwitterIcon size={32} round />
                </TwitterShareButton>
                <WhatsappShareButton
                  url={window.location}
                  title={caption.Text ? caption.Text[0] : ''}
                  separator=":: " style={{ marginLeft: "1em" }}>
                  <WhatsappIcon size={32} round />
                </WhatsappShareButton>

                <EmailShareButton
                  url={window.location}
                  subject={'Check out this picture on CapThatPic!'}
                  body={caption.Text ? caption.Text[0] + ' ' + displayHashtags(post.Tags) : displayHashtags(post.Tags)} style={{ marginLeft: "1em" }}>
                  <EmailIcon size={32} round />
                </EmailShareButton>

                <RedditShareButton
                  url={window.location}
                  title={caption.Text ? '"' + caption.Text[0] + '" by CapThatPic' : 'Check this picture on CapThatPic'}
                  windowWidth={660} windowHeight={460} style={{ marginLeft: "1em" }}>
                  <RedditIcon size={32} round />
                </RedditShareButton>
              </Row>
              : null}
            <Row style={{ marginTop: "3em" }}>
              <Col md="3"></Col>
              <Col md="6">
                <span style={{ textAlign: "center" }}><a href="/posts">See more posts !</a></span>
              </Col>
            </Row>
          </Col>
        </Row>
      </Container>
      <Footer />
    </Fragment>
  );
};

export default PostPage;
