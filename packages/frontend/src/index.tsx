import { render } from 'solid-js/web';
import App from './App';
import './daisyui.css'
import 'uno.css';
import {Router} from 'solid-app-router';

render(() => (<Router>
    <App />
    </Router>
    ), document.querySelector('div#root')!);
