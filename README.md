# cap-that-pic
### Code Chrysalis X Mercari Greenfield Project by [Sahil](https://github.com/sahil505), [Aniket](https://github.com/aniket1743), [Liu](https://github.com/Rocuku) & [Shashank](https://github.com/shashankjakka)
---
*Cap That Pic* generates suitable captions for your images

## What is it?
Just think how many times you struggle coming up with good captions for you images, this application helps you to come up with artistic captions that best suits the attributes of the image.

## How do you use it?
> Cap That Pic application supports two image input method i:e upload from your local computer or enter an online (web) image URL.
1. User uploads the image from local computer by clicking on *Upload Image* or enter an online (web) image URL in the 'Enter Image URL' input box.
2. Then click on CapThatPic to generate the caption that best matches your image. Yayy <3
3. You can then share your image to Instagram or Facebook with the generated caption using the *Facebook Share* or *Instagram Share* button.

## Minimal Viable Product (MVP) [Using a User Story]
- I am John Doe
- I recently clicked some cool images on a hiking trip.
- I want to share the image on social media to impress my friends by using an artistic caption for an image.
- I need a product that can generate an artistic caption for my images.
- I upload the image on ‘Cap That Pic’ and yeah!!!! I can now easily get amazing captions for my images and I can also share on Instagram and Facebook ;)

## Technologies Used
- Cap That Pic uses *ReactJS* as the Frontend structural framework.
- Backend architecture is implemented using *Golang*
- *CircleCI* for continous integration and delivery with automated testing.
<!-- - Docker -->
- *Heroku* as a cloud platform for deployment and managing the application.
- *Azure (Computer Vision)* API to extrate tags from user's image.
- *MusixMatch* API to fetch lyrics corresponding to an artist.

## Challenges
1. What if tags extraced from image are less?
2. % matching of tags (extracted from azure API) with the lyrics (extracted from MusixMatch API)
3. If the random artist selected (in backend) does not have the extracted tags in it's lyrics, what to do?
4. Ways to upload multiple images at the same time and generate captions.
5. Connect our application with Instagram/Facebook/Twitter to share the image with caption directly without leaving our application.

## Task Assignment
1. <u> Sahil Khokhar </u>: Build the User Interface to upload the image from local machine
2. <u> Aniket Choudhary </u>: Make endpoints and write test cases to fetch data from MusixMatch API.
3. <u> Shashank Jakka </u>: Make endpoints and write test cases to fetch data from Azure API.
4. <u> Liu Songjie </u>: Implement the algorithm to generate caption by using the lyrics obtained from MusixMatch API.