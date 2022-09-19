import React from 'react';
import './Box.css';

class Box extends React.Component {
    constructor(props) {
        super(props);
        this.onDelete = this.onDelete.bind(this);
    }
    onDelete() {
        this.props.delete(this.props.index);
    }
    render() {
        return (
            <div>
                <div className="box" style={{
                    display: 'inline-block',
                    width: this.props.width + 'px',
                    height: this.props.height + 'px',
                    backgroundColor: this.props.color,
                }}></div>
                <button onClick={this.onDelete}>delete</button>
            </div>
            
        );
    }
}

export default Box;