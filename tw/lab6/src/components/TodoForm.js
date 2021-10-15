import React, { Component } from "react";
import { Form, FormGroup, Input, Button } from "reactstrap";

class TodoForm extends Component {
  constructor() {
    super();
    this.state = { inputVal: "" };
  }

  handleSubmit = (e) => {
    e.preventDefault();
    if (this.state.inputVal === "") return;

    this.props.addTodo(this.state.inputVal);
    this.setState({ inputVal: "" });
  };
  render() {
    return (
      <Form onSubmit={this.handleSubmit}>
        <FormGroup className="mb-3">
          <Input
            type="text"
            className="form-control"
            placeholder="enter task"
            value={this.state.inputVal}
            onChange={(e) => this.setState({ inputVal: e.target.value })}
          />{" "}
          <FormGroup>
            <Button type="submit" className="col-sm-6 " color="primary">
              Add ToDo
            </Button>
            <Button
              type="button"
              className="col-sm-6 "
              color="danger"
              onClick={() => this.props.removeCompleted()}
            >
              Delete ToDo
            </Button>
          </FormGroup>
        </FormGroup>
      </Form>
    );
  }
}

export default TodoForm;
