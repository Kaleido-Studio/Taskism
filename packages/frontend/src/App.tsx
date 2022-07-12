import { Component } from 'solid-js';
import toast, { Toaster } from 'solid-toast';
import { IoCheckmarkCircle } from 'solid-icons/io';
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
    <button
      onClick={(e) => {
        e.preventDefault();
        toast.success('这是一条消息', {
          duration: Infinity,
          icon: <IoCheckmarkCircle size={24} />,
        });
      }}
    >
      点我
    </button>
    <Toaster position="top-center" />
  </div>
);

export default App;
