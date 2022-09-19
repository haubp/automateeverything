import React from 'react';
import './ColorGame.css';

class Cell extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            backgroundColor: this.getRandomColor()
        };
        this.changeColor = this.changeColor.bind(this);
    }
    getRandomColor() {
        return `#${Math.floor(Math.random() * 10)}${Math.floor(Math.random() * 10)}${Math.floor(Math.random() * 10)}`;
    }
    changeColor() {
        this.setState({
            backgroundColor: this.getRandomColor()
        });
    }
    render() {
        return (
            <div className="cell" onClick={this.changeColor} style={{backgroundColor: this.state.backgroundColor}}></div>
        );
    }
}

class ColorGame extends React.Component {
    render() {
        return (
            <div className="gameArea">
                <div>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                </div>
                <div>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                </div>
                <div>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                </div>
                <div>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                </div>
                <div>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                    <Cell></Cell>
                </div>
            </div>
        );
    }
}

export default ColorGame;