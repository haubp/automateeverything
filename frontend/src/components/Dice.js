import React from 'react';
import './Dice.css';

function decideDotClass(dotPosition, diceNumber) {
    switch(dotPosition) {
        case 1: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot notShowDot";
                                case 3:
                                    return "dot notShowDot";
                                case 4:
                                    return "dot";
                                case 5:
                                    return "dot";
                                case 6:
                                    return "dot";
                                default:
                                    return "";
                            }
                        }
        case 2: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot notShowDot";
                                case 3:
                                    return "dot notShowDot";
                                case 4:
                                    return "dot notShowDot";
                                case 5:
                                    return "dot notShowDot";
                                case 6:
                                    return "dot notShowDot";
                                default:
                                        return "";
                            }
                        }
        case 3: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot";
                                case 3:
                                    return "dot";
                                case 4:
                                    return "dot";
                                case 5:
                                    return "dot";
                                case 6:
                                    return "dot";
                                default:
                                        return "";
                            }
                        }
        case 4: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot notShowDot";
                                case 3:
                                    return "dot notShowDot";
                                case 4:
                                    return "dot notShowDot";
                                case 5:
                                    return "dot notShowDot";
                                case 6:
                                    return "dot";
                                default:
                                    return "";
                            }
                        }
        case 5: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot";
                                case 2:
                                    return "dot notShowDot";
                                case 3:
                                    return "dot";
                                case 4:
                                    return "dot notShowDot";
                                case 5:
                                    return "dot";
                                case 6:
                                    return "dot notShowDot";
                                default:
                                    return "";
                            }
                        }
        case 6: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot notShowDot";
                                case 3:
                                    return "dot notShowDot";
                                case 4:
                                    return "dot notShowDot";
                                case 5:
                                    return "dot notShowDot";
                                case 6:
                                    return "dot";
                                default:
                                    return "";
                            }
                        }
        case 7: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot";
                                case 3:
                                    return "dot";
                                case 4:
                                    return "dot";
                                case 5:
                                    return "dot";
                                case 6:
                                    return "dot";
                                default:
                                        return "";
                            }
                        }
        case 8: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot notShowDot";
                                case 3:
                                    return "dot notShowDot";
                                case 4:
                                    return "dot notShowDot";
                                case 5:
                                    return "dot notShowDot";
                                case 6:
                                    return "dot notShowDot";
                                default:
                                    return "";
                            }
                        }
        case 9: return function() {
                            switch(diceNumber) {
                                case 1:
                                    return "dot notShowDot";
                                case 2:
                                    return "dot notShowDot";
                                case 3:
                                    return "dot notShowDot";
                                case 4:
                                    return "dot";
                                case 5:
                                    return "dot";
                                case 6:
                                    return "dot";
                                default:
                                        return "";
                            }
                        }
        default: return function() {
            return "";
        }
    }
}

const Dice = (props) => {
    return (
        <div className={props.rolling ? 'dice shaking' : 'dice'}>
            <div className="row">
                <div className={decideDotClass(1, props.number)()}></div>
                <div className={decideDotClass(2, props.number)()}></div>
                <div className={decideDotClass(3, props.number)()}></div>
            </div>
            <div className="row">
                <div className={decideDotClass(4, props.number)()}></div>
                <div className={decideDotClass(5, props.number)()}></div>
                <div className={decideDotClass(6, props.number)()}></div>
            </div>
            <div className="row">
                <div className={decideDotClass(7, props.number)()}></div>
                <div className={decideDotClass(8, props.number)()}></div>
                <div className={decideDotClass(9, props.number)()}></div>
            </div>
        </div>
    )
};

export default Dice;