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

## Demo Time!

[Try out our Application](https://cap-that-pic.herokuapp.com/).

## Want to Contribute?

Please create a [new issue](https://github.com/CoolDogee/cap-that-pic/issues) and make a [Pull Request](https://github.com/CoolDogee/cap-that-pic/pulls) corresponding to that issue :)

## Unique Selling Proposition (USP) for the Product

After extracting the tags from Azure (Computer Vision) API and lyrics from Musixmatch then we execute our custom build algorithm to find the best matched caption for the picture which is the key feature or USP of our product.

## Minimal Viable Product (MVP) [Using a User Story]

- I am John Doe
- I recently clicked some cool images on a hiking trip.
- I want to share the image on social media to impress my friends by using an artistic caption for an image.
- I need a product that can generate an artistic caption for my images.
- I upload the image on ‘Cap That Pic’ and yeah!!!! I can now easily get amazing captions for my images and I now share these on Facebook or Instagram ;)

## Essential Features

- Structured User Interface for user to upload image or enter a URL.
- Backend architecture (routes and handling requests) to connect with Azure & MusixMatch API.
- Custom build algorithm to generate a best matching caption based on the tags and lyrics fetched from the above APIs.
- Use captcha to verify users

## Technologies Used

- Cap That Pic uses _ReactJS_ as the Frontend structural framework.
- Backend architecture is implemented using _Golang_
- _Azure (Computer Vision)_ SDK to extract tags from user's image.
- _MongoDB_ contains a dataset of song lyrics corresponding to several artists.
- _CircleCI_ for continuous integration and delivery with automated testing.
- _Docker Containers_ to package and deploy the application as one package.
- _Heroku_ as a cloud platform for deployment and managing the application.

## Engineering Challenges

1. What if tags extraced from image are less?
2. % matching of tags (extracted from azure API) with the lyrics (extracted from MusixMatch API)
3. If the random artist selected (in backend) does not have the extracted tags in it's lyrics, what to do?
4. Ways to upload multiple images at the same time and generate captions.
5. Connect our application with Instagram/Facebook/Twitter to share the image with caption directly without leaving our application.
6. Integrate ReactJs with Backend written in GoLang.

## Future Goals

1. Upload images from local computer.
2. Login/Logout using Facebook/Twitter.
3. Implementation of reCAPTCHA.
4. A URL that generates caption on its own (so that friends can share it among themselves).
5. Improve the UI/UX of the application.
