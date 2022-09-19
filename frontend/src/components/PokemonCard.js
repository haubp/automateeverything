import React from 'react';
import './PokemonCard.css';

class PokemonCard extends React.Component {
    render() {
        return (
            <div className="Card">
                <div className="name">{this.props.info.name}</div>
                <img className="cardImg" src={this.props.url + this.props.info.id + ".png"} alt={this.props.info.name}></img>
                <div className="type">Type: {this.props.info.type}</div>
                <div className="exp">Exp: {this.props.info.base_experience}</div>
            </div>
        );
    }
}

export default PokemonCard;
