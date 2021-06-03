import React from 'react';
import PropTypes from 'prop-types';
import './Styles/MainPageContent.css';

const TransactionForm = ({ setField, submitForm, values, fields }) => {
    return <>
        <form id="trans" onSubmit={submitForm}>
            {fields.map(d => {
                const { key, name } = d;
                return <div key={key}>
                    <span>{name}: </span>
                    <input
                        value={values[key]}
                        name={key}
                        onChange={setField}
                        type="text"
                    />
                </div>
            })}
            <div key="type">
                <span>Transaction Type: </span>
                <select id="type" name="type" value={values["type"]} onChange={setField}>
                    <option value="rent">Rent</option>
                    <option value="gas">Gas</option>
                    <option value="shopping">Shopping</option>
                    <option value="etc">ETC</option>
                </select>
            </div>
            
            <input type="submit" value="Submit" />
        </form>
    </>
}

TransactionForm.propTypes = {
    setField: PropTypes.func.isRequired,
    submitForm: PropTypes.func.isRequired,
    values: PropTypes.shape({
        uid: PropTypes.string.isRequired,
        name: PropTypes.string,
        memo: PropTypes.string,
        date: PropTypes.string.isRequired,
        amount: PropTypes.string.isRequired,
        type: PropTypes.string.isRequired
    }),
    fields: PropTypes.arrayOf(PropTypes.shape({
        key: PropTypes.string,
        name: PropTypes.string
    }))
}

export default TransactionForm;