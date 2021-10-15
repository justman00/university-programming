import React, { Component } from 'react';
import TodoList from './components/TodoList';
import TodoForm from './components/TodoForm';
import MetaTags from 'react-meta-tags';
import styled from 'styled-components';
import { Container, Row } from 'reactstrap';
import './App.css';

const StyledDiv = styled.div`
  border-top-left-radius: 42px;
  background-color: white;
  padding: 2em;
`;
const StyledH2 = styled.h2`
  font-family: Arial, Helvetica, sans-serif;
`;
const StyledH6 = styled.h6`
  color: gray;
`;

const LOCALSTORAGE_KEY = 'todos';

class App extends Component {
  constructor() {
    super();
    const localStorageTodos = window.localStorage.getItem(LOCALSTORAGE_KEY);
    this.state = {
      todos: localStorageTodos ? JSON.parse(localStorageTodos) : [],
    };
  }

  onWriteToLocalStorage = (todos) => {
    const jsonTodos = JSON.stringify(todos);
    window.localStorage.setItem(LOCALSTORAGE_KEY, jsonTodos);
  };

  onUpdateData = (todos) => {
    this.setState({ todos });
    this.onWriteToLocalStorage(todos);
  };

  addTodo = (val) => {
    const newTodos = [
      ...this.state.todos,
      {
        task: val,
        id: Date.now(),
        completed: false,
      },
    ];

    this.onUpdateData(newTodos);
  };

  toggleCompleted = (id) => {
    const updatedTodos = this.state.todos.map((todo) => {
      if (todo.id === id) {
        return {
          ...todo,
          completed: !todo.completed,
        };
      }
      return todo;
    });
    this.onUpdateData(updatedTodos);
  };
  removeCompleted = () => {
    const filteredTodos = this.state.todos.filter((todo) => {
      if (todo.completed) {
        return false;
      }
      return true;
    });
    this.onUpdateData(filteredTodos);
  };

  render() {
    return (
      <Container className="mt-5">
        <MetaTags>
          <meta
            name="viewport"
            content="width=device-width, initial-scale=1.0"
          ></meta>
        </MetaTags>
        <Row className="d-flex justify-content-center">
          <Container className="col-md-6">
            <StyledDiv>
              <StyledH2>ToDo List:</StyledH2>
              <StyledH6>
                You have to complete {this.state.todos.length}
                {this.state.todos.length === 1 ? ' task' : ' tasks'}
              </StyledH6>
              <TodoList
                todos={this.state.todos}
                toggleCompleted={this.toggleCompleted}
              />
              <TodoForm
                addTodo={this.addTodo}
                removeCompleted={this.removeCompleted}
              />
            </StyledDiv>
          </Container>
        </Row>
      </Container>
    );
  }
}

export default App;
