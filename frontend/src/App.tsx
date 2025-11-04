import React, { useEffect } from 'react';
import './App.css';

function App() {

  useEffect(() => {
    fetch('http://localhost:9999/getBooks')
      .then(response => {
        console.log(response);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        console.log(response.json());
      })
      // .then(fetchedData => {
      //   setData(fetchedData);
      //   setLoading(false);
      // })
      .catch(error => {
        console.log(error);
        // setError(error);
        // setLoading(false);
      });
  }, []);

  return (
    <div className="App">
      Hello World
    </div>
  );
}

export default App;
