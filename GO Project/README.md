#  CPE 490 Final Project: Topic 2 ~ A Simple Go Chat Application
A project by: Akash Rana, Brianna Garland, & Dominic Zecchino 



## Team Profiles
| Name | GitHub | 
| ------ | ------ |
| Brianna Garland | [https://github.com/BriannaPGarland] |
| Akash Rana | [https://github.com/akashrana290]|
| Dominic Zecchino| [https://github.com/dzecchino]|


## Introduction
Go is a statistically typed compiled language built with the intention of providing users with a simple yet effective language to build software. This project focuses on building a basic messaging application through use of Go and React, which is a JavaScript library. Broken down into 6 distinct sections, the process included an initial setup for the application, achieving  simple communication, frontend development, multiple client handling, frontend upgrades, and dockerizing the backend. Despite some complications along the way, our group was successful in locally hosting a working application.



> The overarching goal of this project 
> was to create a simple chat application
> using the programming language "Go" 
> allowing at least 3 users to be able to  
> communicate with each other in a group
> chat of sorts. 


## How Our Program Works

- Initial setup began with the creation of a backend directory and a few terminal commands to initialize go modules within our backend directory. The generated modules include the packages and versions required for a successful build of the application as well as validation of the versions. A test function was used to ensure proper functioning up to this point. Frontend development then started with a new directory. Node Package Manager and Node Package Execute, abbreviated as npm and npx respectively, were used to instantiate this project as a basic ReactJS application, adding several files and subdirectories to the frontend directory. Finalizing the initial setup meant testing that our application was able to be hosted using the terminal command npm start.

- Simple communication within the app was completed with help from WebSocket. This is a computer communications protocol that will, for our purposes, listen for messages and write them back via the same WebSocket. This is advantageous because it eliminates the necessity for another TCP handshake for each time that the server is polled for updates. An available GitHub repository was used for installation of WebSocket packages in our backend directory. A basic server with room for enhancements was then created for test purposes, printing “Simple Server” to prove that no unwanted issues arose. The previously mentioned advantage of a WebSocket endpoint was achieved through upgrading (using websocket.Upgrader{ }) from the standard http endpoint, which does require a handshake for every poll. Regarding backend WebSocket functionality, the final step was to write a function to read incoming messages and print them back out. This however, is not the last incorporation of code for WebSocket. Frontend WebSocket connection code was needed to interact with the server. Seen as an output in Figure _____, this meant that 2 functions were needed for the purpose of listening for new events and sending the messages back out. Concluding the simple connection phase of this project, a button was initialized in our App.js file.

## Our Process

Insert from the report 

## Final Results
After several hours of work with the team’s chat application, they were able to get a functioning chat application using go and react on the web browser. This application allowed three separate browsers to send and view messages from the other users on the other browsers. The team was also able to add several unique aspects to the chat application to make it even more their own. The team added a clear chat history button that allows the individual user to clear their chat history they see on the screen while still staying in communication with the other browser windows. The team was also able to add and create their own user interface with a lot of time and dedication to the project to make it Stevens themed. You can view the final result example below, but you can also see our source code and a video demonstration of the application running at the following link: [https://github.com/akashrana290/cpe490_go_project]. 

*Final Application Screenshots*: 
This screenshot shows three separate browsers communicating with each other and essentially publishing and subscribing in order to see and maintain messages sent. You can also see when a new user is added and or disconnected.

## Conclusions
Overall, this project really turned out to be one of the best projects the team has worked on this semester. It was very rewarding to be able to take what they spent the entire semester learning and actually get to apply some of it to create something cool! The team was able to learn a lot about the Go Programming language and even React which was really a double gift as both are used so widely in the field today. It is amazing to see what we were able to create just for none of the team ever using Go before and only one of them being familiar with react. Even though the team did use a tutorial to assist them in the core set up, they were still able, through what they learned in the semester, fully understand what was being asked of them and it allowed them to problem solve when they ran into errors along the way. The knowledge they gained allowed them to add a few of their own features and designs and also to very indepthly comment on their source code to show an in-depth understanding of the material. There were several other sources out there that just gave the result of essentially a fully functional texting app like slack but without any explanation. The team wanted to keep it somewhat simple so they could gain the most from it and make contributions of their own. In the future, this application could be opened up to so many things. The team say a lot of resources including messaging channels or users to the application. This provides the team a good opportunity to possibly continue their work and gain a deeper understanding of such an useful programming language and skill set! 
 
 
 
## Resources
https://golang.org/doc/tutorial/getting-started  (Hello World Tutorial) 
First Tutrial: https://getstream.io/blog/building-a-chat-server-with-go/ 
Second tutorial:https://tutorialedge.net/projects/chat-system-in-go-and-react/ 

