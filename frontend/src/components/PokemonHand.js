import React from 'react';
import PokemonCard from './PokemonCard';
import PokemonsInfo from './model/pokemons';
import './PokemonHand.css';

class PokemonHand extends React.Component {
    render() {
        if (this.props.randomIndexArr.length > 0)
        {
            return (
                <div className="cards">
                {
                    this.props.randomIndexArr.map((index) => <PokemonCard key={PokemonsInfo.list[index].id} url={PokemonsInfo.url} info={PokemonsInfo.list[index]} />)
                }
                </div>
            );
        }
        else
        {
            return (
                <div>Empty Hand</div>
            );
        }
    }
}

export default PokemonHand;
