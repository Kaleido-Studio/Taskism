import { Route, Routes } from 'solid-app-router';
import { Component } from 'solid-js';
import { LandingGuest } from './pages/LandingGuest';

const App: Component = () => (
  <Routes>
    <Route path="/" component={LandingGuest} />
  </Routes>
);

export default App;
