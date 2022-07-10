import { Component } from 'solid-js';
import favicon from './assets/favicon.svg';
import './App.css';

const App: Component = () => (
  <div>
    <div class="logo">
      <a href="https://www.solidjs.com" target="_blank" rel="noreferrer">
        <img src={favicon} />
      </a>
    </div>
    <h1>Hello Vite + Solid</h1>
  </div>
);

export default App;
