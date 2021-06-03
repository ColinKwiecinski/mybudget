import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import * as serviceWorker from './serviceWorker';
import Plotly from 'plotly.js-dist';

ReactDOM.render(<App />, document.getElementById('root'));



// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();

<<<<<<< HEAD

=======
// const IP = "https://api.justinlim.me/transactions"
// // Send transaction to backend
// const thisForm = document.getElementById('transactionForm');
// thisForm.addEventListener('submit', async function (e) {
//     e.preventDefault();
//     const formData = new FormData(thisForm).entries()
//     const response = await fetch(IP, {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify(Object.fromEntries(formData))
//     });

//     const result = await response.json();
//     console.log(result)
// });

// thisForm.addEventListener('submit', () => {
//   console.log("working")
// })
>>>>>>> 313e2d95d01f886b1199db8ae4a876b690c53097

// // Display transaction visualization on frontend
// const plotButton = document.getElementById('chartsButton');
// plotButton.addEventListener(plotButton, async function (e) {
//   e.preventDefault();
//   const resp = await fetch(IP + "/1", { // TODO figure out specific user id

//   })
// })

<<<<<<< HEAD
// function barChart(result) {
//   var output = document.getElementById('plotoutput')
//   Plotly.newPlot(output, [{
//     x: ['test', 'test2'],
//     y: [10, 20],
//     type: 'bar'
//   }])
// }
=======
function barChart(result) {
  var output = document.getElementById('plotoutput')
  Plotly.newPlot(output, [{
    x: ['test', 'test2'],
    y: [10, 20],
    type: 'bar'
  }])
}
>>>>>>> 313e2d95d01f886b1199db8ae4a876b690c53097
