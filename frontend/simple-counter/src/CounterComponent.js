import React, { useState } from "react";

function CounterComponent(props) {
  const [counter, setCounter] = useState(0);
  return (
    <div>
      <h1>Simple Counter Application</h1>
      <p>Counter: {counter}</p>
      <button onClick={() => setCounter(counter + 1)}>Increment</button>
      <br></br>
      <button onClick={() => setCounter(counter - 1)}>Decrement</button>
      <br></br>
      <button onClick={() => setCounter(0)}>Reset</button>
    </div>
  );
}
export default CounterComponent;
