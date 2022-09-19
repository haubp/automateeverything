import React from 'react';
import TodoItem from './TodoItem';
import uuid from 'react-uuid';
import './TodoList.css';

class TodoList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            todos: [],
            todo: '',
        };
        this.deleteItem = this.deleteItem.bind(this);
        this.editItem = this.editItem.bind(this);
        this.onChange = this.onChange.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
        this.onDone = this.onDone.bind(this);
    }
    deleteItem(index) {
        let newTodos = [...this.state.todos];
        newTodos.splice(index, 1);
        this.setState({
            todos: newTodos
        });
    }
    editItem(index, text) {
        let newTodos = [...this.state.todos];
        newTodos[index].text = text;
        newTodos[index].isDone = false;
        this.setState({
            todos: newTodos
        });
    }
    onChange(e) {
        this.setState({
            todo: e.target.value
        });
    }
    onSubmit(e) {
        e.preventDefault();
        this.state.todos.push({
            text: this.state.todo,
            isDone: false
        });
        this.setState({
            todo: ''
        });
    }
    onDone(index) {
        let newTodos = [...this.state.todos];
        newTodos[index].isDone = !newTodos[index].isDone;
        this.setState({
            todos: newTodos
        });
    }
    render() {
        return (
            <div className="todoList">
                <form onSubmit={this.onSubmit}>
                    <label>
                        Todo 
                        <input name="todo" value={this.state.todo} onChange={this.onChange}></input>
                        <button>Add</button>
                    </label>
                </form>
                {
                    this.state.todos.map((todo, i) => <TodoItem key={uuid()} done={this.onDone} isDone={todo.isDone} edit={this.editItem} delete={this.deleteItem} text={todo.text} index={i} />)
                }

            </div>
        );
    }
}

export default TodoList;