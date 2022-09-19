import React from 'react';
import axios from 'axios';
import './CardsDeck.css';

const DESK_URL = "https://deckofcardsapi.com/api/deck/new/shuffle";

class CardsDesk extends React.Component {
    constructor(props) {
        super(props);
        this.deck_id = '';
        this.state = {
            remaining: 52,
            list: [],
        };
        this.onButtonClick = this.onButtonClick.bind(this);
    }
    async componentDidMount() {
        const response = await axios.get(DESK_URL);
        this.deck_id = response.data.deck_id;
        console.log(response);
    }
    async onButtonClick() {
        const response = await axios.get(`https://deckofcardsapi.com/api/deck/${this.deck_id}/draw`);
        this.setState({
            list: [...this.state.list, {url: response.data.cards[0].image, rotate: Math.floor(Math.random() * 360)}],
            remaining: this.state.remaining - 1
        });
    }
    render() {
        return (
            <div className="board">
                {
                    this.state.remaining === 0 ?
                    <div>Finish</div> :
                    <div className="card-deck">
                        <button onClick={this.onButtonClick} style={{margin: 50}}>Shuffle</button>
                        {
                            this.state.list.map(card => <img src={card.url} style={{transform: `rotate(${card.rotate}deg)`}} key={card.url} alt="card" />)
                        }
                    </div>
                }
            </div>
        );
    }
}

export default CardsDesk;