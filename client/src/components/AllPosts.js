import React, { Fragment, useState } from "react";
import Message from "./Message";
import Loading from "./Loading";
import axios from "axios";
import { API_URL } from "../config";
import Footer from "./Footer";
import { Container, Row, Col, Card } from "react-bootstrap";
import Logo from "../images/capthatpic.png";

function displayHashtags(tags) {
  if (!tags || !tags.length) {
    return '';
  }
  var tagshash = tags.map(t => '#' + t.split(" ").join("_"));

  return tagshash.join(', ');
}

export const AllPostsPage = () => {
  const [message, setMessage] = useState("");
  const [posts, setPosts] = useState([]);
  const [loading, setLoadingstatus] = useState(false);
  const [initialFetch, setInitialFetch] = useState(false);

  const getPosts = () => {
    setInitialFetch(true);
    setLoadingstatus(true);
    axios.get(API_URL + '/api/v1/posts')
      .then(function (response) {
        console.log('Received posts');
        console.log(response.data.info);
        setPosts(response.data.info);
      })
      .catch(function (error) {
        console.log("Could not fetch posts!");
        setMessage("Could not fetch posts!");
        console.log(error.response);
      });
    setLoadingstatus(false);
  };

  if (!initialFetch) {
    getPosts();
  }

  return (
    <Fragment>
      <h1 className="display-3 text-center mb-4">
        <img src={Logo} style={{ height: "1.5em" }} /> Cap That Pic
    </h1>
      <h4 className="text-center mb-4">
        <a href="/">Caption your own photo!</a>
        {/* <Typist>Beautiful Things Don't Ask For Attention. But They Deserve A Perfect Caption</Typist> */}
      </h4>
      <hr></hr>
      {message ? <Message msg={message} /> : null}
      <Container style={{ marginBottom: "5em" }}>
        <Row>
          <Col md="12">{loading ? <Loading /> : null}</Col>
          {posts.map((post) =>
            <Col md="4" style={{ marginBottom: "2em" }}>
              <a href={"/i/" + post.ID}>
                <Card>
                  <Card.Img variant="top" src={post.ImgURL || "holder.js/100px180"} style={{ height: "400px", width: "auto" }} />
                  <Card.Body>
                    <Card.Title>{post.Caption.Text[0] || "caption"}</Card.Title>
                    <Card.Text className="text-muted">{displayHashtags(post.Tags)}</Card.Text>
                  </Card.Body>
                </Card>
              </a>
            </Col>
          )}
        </Row>
      </Container>
      <Footer />
    </Fragment>
  );
};

export default AllPostsPage;
