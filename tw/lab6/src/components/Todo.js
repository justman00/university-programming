import React, { Component } from "react";
import styled from "styled-components";

const StyledListItem = styled.li`
  list-style-type: disc;
`;

class Todo extends Component {
  render() {
    return (
      <StyledListItem
        onClick={() => {
          this.props.toggleCompleted(this.props.id);
        }}
        style={{
          textDecoration: this.props.completed ? "line-through" : "none",
        }}
      >
        {this.props.task}
      </StyledListItem>
    );
  }
}
export default Todo;
