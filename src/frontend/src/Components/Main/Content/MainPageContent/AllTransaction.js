import React, { Component } from "react";
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import './Styles/MainPageContent.css';


class AllTransaction extends Component {
    constructor(props) {
        super(props);
     
        this.state = {
            transactions: [],
            isLoading: false,
            error: null,
          };
      }
     
      componentDidMount() {
        this.setState({ isLoading: true });
     
        fetch("https://api.justinlim.me/transactions/" + this.props.user.id)
        .then(response => {
            if (response.ok) {
              return response.json();
            } else {
              throw new Error('Something went wrong ...');
            }
          })
          .then(data => this.setState({ transactions: data, isLoading: false }))
          .catch(error => this.setState({ error, isLoading: false }));
      }

      render() {
        const { transactions, isLoading, error } = this.state;
        if (error) {
            return <p>{error.message}</p>;
        }
        if (isLoading) {
            return <p>Loading ...</p>;
        }
        return (
            <div>
                <h2 id="alltrans">All Transactions</h2>
                <h3>Total Transactions: {transactions.length}</h3>
                {transactions.map(tran => 
                <div>
                    <h4>Transaction Name: {tran.name}</h4>
                    <ul id="alltransactions">
                        <li><b>Transaction Memo:</b> {tran.memo}</li>
                        <li><b>Transaction Date:</b> {tran.date}</li>
                        <li><b>Transaction Amount:</b> {tran.amount}</li>
                        <li><b>Transaction Type:</b> {tran.type}</li>
                    </ul>
                </div>
                )}
            </div>
          ); 
      }
     
     
}

export default AllTransaction;