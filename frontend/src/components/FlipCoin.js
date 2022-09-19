import React from 'react';
import './FlipCoin.css';
import styled from 'styled-components';
import coin_head from './img/coin_head.jpeg';
import coin_tail from './img/coin_tail.jpeg';

const Button = styled.button`
  background-color: black;
  color: white;
  font-size: 20px;
  padding: 10px 60px;
  border-radius: 5px;
  margin: 10px 0px;
  cursor: pointer;
  width: 60;
  height: 30;
`;

class FlipCoin extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            heads: 0,
            tails: 0,
            total: 0,
            isHead: true
        };
        this.flip = this.flip.bind(this);
    }
    flip() {
        let stt = {...this.state};
        if (Math.floor(Math.random() * 2) === 0) {
            stt.heads += 1;
            stt.isHead = true;
        } else {
            stt.tails += 1;
            stt.isHead = false;
        }
        stt.total += 1;
        this.setState(stt);
    }
    render() {
        return (
            <div className="flipBox">
                <img src={this.state.isHead ? coin_head : coin_tail} alt="coin"></img>
                <p>{`Total ${this.state.total}: ${this.state.heads} heads, ${this.state.tails} tails`}</p>
                <Button onClick={this.flip}>Flip</Button>
            </div>
        );
    }
}

export default FlipCoin;