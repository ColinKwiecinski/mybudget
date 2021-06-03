import React, { useState, useEffect, Component } from 'react';
import PageTypes from '../../../../Constants/PageTypes/PageTypes';
import './Styles/MainPageContent.css';
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import TransactionForm from '../MainPageContent/TransactionForm';
import PropTypes from 'prop-types';
import Errors from '../../../Errors/Errors';


class MainPageContent extends Component {
    static propTypes = {
        setPage: PropTypes.func,
        setAuthToken: PropTypes.func
    }

    constructor(props) {
        super(props);

        this.state = {
            uid: this.props.user.id,
            name: "",
            memo: "",
            date: "",
            amount: "",
            type: ""
        };

        this.fields = [
            {
                name: "Name",
                key: "name"
            },
            {
                name: "Memo",
                key: "memo"
            },
            {
                name: "Transaction Date",
                key: "date"
            },
            {
                name: "Amount",
                key: "amount"
            }];
    }

    /**
     * @description setField will set the field for the provided argument
     */
    setField = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    /**
     * @description setError sets the error message
     */
    setError = (error) => {
        this.setState({ error })
    }

    /**
     * @description submitForm handles the form submission
     */
    submitForm = async (e) => {
        console.log(this.state["type"])
        e.preventDefault();
        const { uid,
            name,
            memo,
            date,
            amount,
            type } = this.state;
        const sendData = {
            uid,
            name,
            memo,
            date,
            amount,
            type
        };
        const response = await fetch(api.base + "/transactions", {
            method: 'POST',
            body: JSON.stringify(sendData),
            headers: new Headers({
                'Content-Type': 'application/json'
            })
        });
        if (response.status >= 300) {
            const error = await response.text();
            this.setError(error);
            return;
        }
        const authToken = this.props.authToken
        localStorage.setItem("Authorization", authToken);
        this.setError("");
        this.props.setAuthToken(authToken);
        const user = await response.json();
        this.props.setUser(user);
    }


    render() {
        const values = this.state;
        return <>
            <div id="welcome">Welcome to MyBudget, {this.props.user.name}</div>
            <TransactionForm
                setField={this.setField}
                submitForm={this.submitForm}
                values={values}
                fields={this.fields} />
            <button onClick={(e) => this.props.setPage(e, PageTypes.analysis)}>View Analysis</button>
            <button onClick={(e) => this.props.setPage(e, PageTypes.allTransactions)}>All Transaction</button>
        </>
        }
    }


export default MainPageContent;