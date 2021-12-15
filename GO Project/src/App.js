import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from "./components/ChatHistory/ChatHistory.jsx";
import ChatInput from "./components/ChatInput/ChatInput.jsx";



//Creates chat history constructor


class App extends Component {
 state = { chatHistory: []}
  

  //call back function that will run automatically alerting the log wiht a new message
  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }

  send(event) {
    if(event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  
clearHistory(){
  //window.location.reload();
this.setState({ chatHistory: [] });

} 

  //Render function that calls to the screen all of our components: Header, ChatHistory and our button
  render() {
    return (
      <div className="App">
        <Header />
        
        <iframe src="https://giphy.com/embed/rtRflhLVzbNWU" width="200" height="300" frameBorder="0" class="giphy-embed" allowFullScreen></iframe><p><a href="https://giphy.com/stickers/duck-enjoy-ease-rtRflhLVzbNWU"></a></p>
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
    
        <button className="clearButton" type="submit" onClick={() =>this.clearHistory()} > Clear Chat History </button>
      </div>
    );
  }
}
export default App;
