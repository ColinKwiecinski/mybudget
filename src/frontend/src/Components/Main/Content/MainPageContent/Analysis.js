import React, { Component } from "react";
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import rd3 from 'react-d3-library'
import { Button, Row, Col, Container } from "react-bootstrap";
import FilteredTransaction from "./FilteredTransaction";
import './Styles/MainPageContent.css';

class Analysis extends Component {
    constructor(props) {
        super(props);
     
        this.state = {
            transactions: [],
            filteredTransaction: [],
            isLoading: false,
            error: null,
            filter: ""
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
        let { transactions, filteredTransaction, isLoading, error, filter } = this.state;
        if (error) {
            return <p>{error.message}</p>;
        }
        if (isLoading) {
            return <p>Loading ...</p>;
        }
        if (this.state.filteredTransaction == []) {
            this.state.filteredTransaction = transactions;
        }
        
        let onResetArray = () => {
            this.setState({ filteredTransaction: this.state.transactions, filter: "" })
        }

        let filterRent = () => {
            const rentTrans = transactions.filter( (trans) => trans.type.includes("rent"));
            this.setState({ filteredTransaction: rentTrans, filter: "Rent" });
        }
        let filterGas= () => {
            const gasTrans = transactions.filter( (trans) => trans.type.includes("gas"));
            this.setState({ filteredTransaction: gasTrans, filter: "Gas" });
        }
        let filterShopping = () => {
            const shoppingTrans = transactions.filter( (trans) => trans.type.includes("shopping"));
            this.setState({ filteredTransaction: shoppingTrans, filter: "Shopping" });
            console.log(filteredTransaction)
        }
        let filterEtc = () => {
            const etcTrans = transactions.filter( (trans) => trans.type.includes("etc"));
            this.setState({ filteredTransaction: etcTrans, filter: "Etc" });
        }
        return (
            <div id="container">
                <div id="vertical-center">
                
                
                    <Button event-key="reset-autos" onClick={onResetArray}>All Transactions</Button>
            
                    <Button event-key="ford-autos" onClick={filterRent}>Filter Rent</Button>
            
                    <Button event-key="chevy-autos" onClick={filterGas}>Filter Gas</Button>
                
                    <Button event-key="jeep-autos" onClick={filterShopping}>Filter Shopping</Button>
                
                    <Button event-key="jeep-autos" onClick={filterEtc}>Filter Etc</Button>
              
            
                </div>
                
            
            <FilteredTransaction data={this.state.filteredTransaction} filter={filter}/>
            </div>
            
          );
      }
     
     
}

export default Analysis;