import React from 'react';
import PokemonHand from './PokemonHand';
import './PokemonDex.css';
import styled from 'styled-components';
import PokemonsInfo from './model/pokemons';
import Dice from './Dice';
import Lottery from './Lottery';
import FlipCoin from './FlipCoin';
import ColorGame from './ColorGame';

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

class PokemonDex extends React.Component {
    constructor(props) {
        super(props);
        this.cardsOrder = [];
        this.state = {
            diceOneNumber: 1,
            diceTwoNumber: 2,
            rolling: false,
            lotteryNums: [1,2,3,4]
        };

        this.onStartButtonClick = this.onStartButtonClick.bind(this);
        this.roll = this.roll.bind(this);
        this.generate = this.generate.bind(this);
    }
    onStartButtonClick() {
        if (this.cardsOrder.length === 0)
        {
            this.cardsOrder = [0,1,2,3,4,5,6,7];
        }
        else 
        {
            this.rotate(this.cardsOrder, 0, 6);
            this.rotate(this.cardsOrder, 1, 7);
            this.rotate(this.cardsOrder, 2, 5);
            this.rotate(this.cardsOrder, 4, 7);
        }
        this.setState({});
    }
    rotate(arr, startIndex, endIndex) {
        let tempArr = JSON.parse(JSON.stringify(arr.slice(startIndex, endIndex)));
        arr[startIndex] = arr[endIndex];
        for(let i = 0; i < tempArr.length; i++) 
        {
            arr[startIndex + 1 + i] = tempArr[i];
        }
    }
    makeWinner(arr) {
        if (arr.length === 0) return "";

        let firstHandScore = 0;
        let secondHandScore = 0;
        for(let i = 0; i < 4; i++) {
            firstHandScore += PokemonsInfo.list[arr[i]].base_experience;
        }
        for(let i = 4; i < 8; i++) {
            secondHandScore += PokemonsInfo.list[arr[i]].base_experience;
        }
        return firstHandScore > secondHandScore ? 'The winner is First Hand' : 'The winner is Second Hand';
    } 
    roll() {
        this.setState({
            diceOneNumber: 1 + Math.floor(6 * Math.random()),
            diceTwoNumber: 1 + Math.floor(6 * Math.random()),
            rolling: true
        });
        setTimeout(() => {
            this.setState({
                rolling: false
            });
        }, 1000);
    }
    generate() {
        let newLotteryNums = [];
        for (let i = 0; i < 4; i++) {
            newLotteryNums.push(Math.floor(Math.random() * 10));
        }
        this.setState({
            lotteryNums: newLotteryNums
        });
    }
    render() {
        return (
            <div className="dex">
                <div className="rowTable">
                    <div className="tableGame">
                        <ColorGame></ColorGame>
                    </div>
                    <div className="tableGame">
                        <FlipCoin />
                    </div>
                </div>
                <div className="tableGame">
                    <Lottery lotteryNums={this.state.lotteryNums} />
                    <Button onClick={this.generate}>Generate</Button>
                </div>
                <div className="tableGame">
                    <div className="rollTable" >
                        <Dice number={this.state.diceOneNumber} rolling={this.state.rolling}></Dice>
                        <Dice number={this.state.diceTwoNumber} rolling={this.state.rolling}></Dice>
                    </div>
                    <Button onClick={this.roll} disabled={this.state.rolling}>
                        {this.state.rolling ? 'Rolling...' : 'Roll'}
                    </Button>
                </div>
                <div className="tableGame">
                    <div className="header">
                        <p>Pokemon Game</p>
                    </div>
                    <div className="gameResult">{this.makeWinner(this.cardsOrder)}</div>
                    <PokemonHand randomIndexArr={this.cardsOrder.slice(0, 4)} />
                    <PokemonHand randomIndexArr={this.cardsOrder.slice(4, 8)} />
                    <Button onClick={this.onStartButtonClick}>Play</Button>
                </div>  
            </div>
        );
    }
}

export default PokemonDex;
