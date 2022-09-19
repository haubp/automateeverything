import React from 'react';
import './Lottery.css';

class Lottery extends React.Component {
    render() {
        return (
            <div className="table">
            {
                this.props.lotteryNums.map((num) => <div className="lottery">{num}</div>)
            }
            </div>
        );
    }
}

export default Lottery;