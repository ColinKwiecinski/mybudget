import React, { Component } from 'react'
import './Styles/MainPageContent.css';
class FilteredTransaction extends Component {
    render() {
        let total = 0.0
        for (let i = 0; i < this.props.data.length; i++) {
            total += Number(this.props.data[i].amount);
        }
        
        return (
            <div>
                <hr class="solid"></hr>
                <h2>{this.props.filter + " Transactions"}</h2>
                <h3>{this.props.filter} Transaction Count: {this.props.data.length}</h3>
                <p><b>Total Amount Spent:</b> ${total}</p>
            </div>
            
        );
    }
}

export default FilteredTransaction