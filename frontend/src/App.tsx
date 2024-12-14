import { Button } from '@carbon/react';
import { ChangeEvent, useState } from 'react';
import { Ping } from '../wailsjs/go/main/App';
import logo from './assets/images/logo-universal.png';
import './main.scss';

function App() {
  const [resultText, setResultText] = useState('Please enter your name below 👇');
  const [name, setName] = useState('');
  const updateName = (e: ChangeEvent<HTMLInputElement>) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Ping().then(updateResultText);
  }

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo" />
      <div id="result" className="result">
        {resultText}
      </div>
      <div id="input" className="input-box">
        <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
        <button className="btn" onClick={greet}>
          Greet
        </button>
        <Button>Greet</Button>
      </div>
    </div>
  );
}

export default App;