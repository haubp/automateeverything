import React from 'react';
import './TodoItem.css';

class TodoItem extends React.Component {
    constructor(props) {
        super(props);
        this.onDelete = this.onDelete.bind(this);
        this.state = {
            showForm: false,
            text: this.props.text
        };
        this.allowEdit = this.allowEdit.bind(this);
        this.onEnter = this.onEnter.bind(this);
        this.onChange = this.onChange.bind(this);
        this.onDone = this.onDone.bind(this);
    }
    onDelete() {
        this.props.delete(this.props.index);
    }
    allowEdit() {
        this.setState({
            showForm: true
        });
    }
    onEnter(e) {
        e.preventDefault();
        this.props.edit(this.props.index, this.state.text);
        this.setState({
            showForm: false
        });
    }
    onChange(e) {
        this.setState({
            text: e.target.value
        });
    }
    onDone(e) {
        this.props.done(this.props.index);
    }
    render() {
        return (
            <div>
                {
                    this.props.isDone ? "(Done)" + this.props.text : this.props.text
                }
                <button onClick={this.onDone}>Done</button>
                <button onClick={this.onDelete}>x</button>
                <button onClick={this.allowEdit}>Edit</button>
                {
                    this.state.showForm ? 
                    <form onSubmit={this.onEnter}>
                        <label>
                            Edit
                            <input name="text" value={this.state.text} onChange={this.onChange}></input>
                            <button>Enter</button>
                        </label>
                    </form> : <div></div>
                }
            </div>
        );
    }
}

export default TodoItem;