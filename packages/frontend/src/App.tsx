import { Component } from 'solid-js';
import toast, { Toaster } from 'solid-toast';
import './App.css';
import favicon from './assets/favicon.svg';

const App: Component = () => (
  <div>
    <div class="logo">
      <a href="https://www.solidjs.com" target="_blank" rel="noreferrer">
        <img src={favicon} />
      </a>
    </div>
    <h1>Hello Vite + Solid</h1>
    <button
      onClick={(e) => {
        e.preventDefault();
        toast.success('这是一条消息', {
          duration: Infinity,
        });
      }}
    >
      点我
    </button>
    <Toaster position="top-center" />
  </div>
);

export default App;
