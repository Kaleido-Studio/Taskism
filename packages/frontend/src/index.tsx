import { render } from 'solid-js/web';
import App from './App';
import {Router} from 'solid-app-router';
import 'virtual:windi.css'

render(() => (<Router>
    <App />
    </Router>
    ), document.querySelector('div#root')!);
