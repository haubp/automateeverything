import React from 'react';
import Box from './Box';
import uuid from 'react-uuid';
import './BoxList.css';

class BoxList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            width: 100,
            height: 100,
            color: 'blue',
            boxList: []
        };
        this.onChange = this.onChange.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
        this.onDelete = this.onDelete.bind(this);
    }
    onChange(e) {
        this.setState({
            [e.target.name]: e.target.value
        });
    }
    onSubmit(e) {
        e.preventDefault();
        this.state.boxList.push({
            width: this.state.width,
            height: this.state.height,
            color: this.state.color
        });
        this.setState({
            width: 100,
            height: 100,
            color: 'blue',
        });
    }
    onDelete(index) {
        let newBoxList = [...this.state.boxList];
        newBoxList.splice(index, 1);
        this.setState({
            boxList: newBoxList
        });
    }
    render() {
        return (
            <div>
                <form className="addBoxContainer" onSubmit={this.onSubmit}>
                    <label>
                        width 
                        <input value={this.state.width} name="width" onChange={this.onChange}></input>
                    </label>
                    <label>
                        height 
                        <input value={this.state.height} name="height" onChange={this.onChange}></input>
                    </label>
                    <label>
                        color 
                        <input value={this.state.color} name="color" onChange={this.onChange}></input>
                    </label>
                    <button>Add</button>
                </form>
                {
                    this.state.boxList.map((box, index) => <Box delete={this.onDelete} index={index} key={uuid()} width={box.width} height={box.height} color={box.color} />)
                }
            </div>
        );
    }
}

export default BoxList;