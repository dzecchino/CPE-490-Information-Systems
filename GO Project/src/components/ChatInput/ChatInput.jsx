import React, { Component } from "react";
import "./ChatInput.scss";

//Creates chat input box
class ChatInput extends Component {
  render() {
    return (
      <div className="ChatInput">
        <input type ="text" onKeyDown={this.props.send} />
      </div>
    );
  }
}

export default ChatInput;