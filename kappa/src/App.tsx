import React from 'react';
import './App.css';
import { Header } from './components/Header';
import { TodoList } from "./components/TodoList";

function App() {
  return (
    <div className="App">
      <Header />
      <TodoList todos={[
        {title: "Do dishes", isCompleted: true, description: "Hello, sample description"},
        {title: "Mow the lawn", isCompleted: false}

        ]} />
    </div>
  );
}

export default App;
