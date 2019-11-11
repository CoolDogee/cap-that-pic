# cap-that-pic

### Code Chrysalis X Mercari Greenfield Project by [Sahil](https://github.com/sahil505), [Aniket](https://github.com/aniket1743), [Liu](https://github.com/Rocuku) & [Shashank](https://github.com/shashankjakka)

---

_Cap That Pic_ generates suitable captions for your images. It uses the technology of Microsoft Azure's Computer Vision and MusixMatch API to generate intelligent and artistic captions that best describe a users image. Give it a try ;)

## What is it?

Just think how many times you struggle coming up with good captions for you images, this application helps you to come up with artistic captions that best suits the attributes of the image.

## How do you use it?

> Cap That Pic application supports two image input method i:e upload from your local computer or enter an online (web) image URL.

1. User uploads the image from local computer by clicking on _Upload Image_ or enter an online (web) image URL in the 'Enter Image URL' input box.
2. Then click on CapThatPic to generate the caption that best matches your image. Yayy <3
3. You can then share your image to Instagram or Facebook with the generated caption using the _Facebook Share_ or _Instagram Share_ button.

## Unique Selling Proposition (USP) for the Product

After extracting the tags from Azure (Computer Vision) API and lyrics from Musixmatch then we execute our custom build algorithm to find the best matched caption for the picture which is the key feature or USP of our product.

## Minimal Viable Product (MVP) [Using a User Story]

- I am John Doe
- I recently clicked some cool images on a hiking trip.
- I want to share the image on social media to impress my friends by using an artistic caption for an image.
- I need a product that can generate an artistic caption for my images.
- I upload the image on ‘Cap That Pic’ and yeah!!!! I can now easily get amazing captions for my images and I can also share on Instagram and Facebook ;)

## Essential Features

- Structured User Interface for user to upload image or enter a URL.
- Backend architecture (routes and handling requests) to connect with Azure & MusixMatch API.
- Custom build algorithm to generate a best matching caption based on the tags and lyrics fetched from the above APIs.
- Use captcha to verify users

## Technologies Used

- Cap That Pic uses _ReactJS_ as the Frontend structural framework.
- Backend architecture is implemented using _Golang_
- _CircleCI_ for continuous integration and delivery with automated testing.
- _Heroku_ as a cloud platform for deployment and managing the application.
- _Azure (Computer Vision)_ API to extract tags from user's image.
- _MusixMatch_ API to fetch lyrics corresponding to an artist.

## Challenges

1. What if tags extraced from image are less?
2. % matching of tags (extracted from azure API) with the lyrics (extracted from MusixMatch API)
3. If the random artist selected (in backend) does not have the extracted tags in it's lyrics, what to do?
4. Ways to upload multiple images at the same time and generate captions.
5. Connect our application with Instagram/Facebook/Twitter to share the image with caption directly without leaving our application.
6. Integrate ReactJs with Backend written in GoLang.

## Tasks & Assignment

- Sahil
  - Build the User Interface structure so that a user can upload the image from local machine or provide an image URL.
  - Try to connect the application to Instagram/Facebook/Twitter.
- Aniket
  - Make endpoints and write test cases to fetch data from MusixMatch API.
- Liu
  - Implement the algorithm to generate caption by using the lyrics obtained from MusixMatch API.
- Shashank
  - Make endpoints and write test cases to fetch data from Azure API.

## Day-to-Day Goals

- Day 0
  - [x] Come up with solo project ideas.
  - [x] Finalize one project idea for the team.
  - [x] Decide features that we need for MVP using a user story.
  - [x] Brainstorm risks and challenges that we might face while building our product.
  - [x] Come up with the USP for our product.
  - [x] Set team assignments for the team project.
- Day 1
  - [x] Make a short presentation for stakeholder meeting.
  - [x] Stand-up after lunch : Catch up on progress, reset todos and milestones
  - [x] Create a configuration file for CircleCI to run automated tests.
  - [x] Create a pipeline to integrate frontend with the backend.
  - [x] Find a way to deploy our app on heroku.
  - [ ] Create endpoints to fetch tags from the Azure API.
  - [x] Create endpoints to fetch song lyrics from the Musixmatch API.
  - [x] Implement the basic algorithm to generate caption using tags and song lyrics.
- Day 2
  - [ ] Stand-up after lunch
  - [ ] Add feature to enter image URL
