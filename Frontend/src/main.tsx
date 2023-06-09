// @ts-ignore
import React from "react";
import ReactDOM from 'react-dom/client';
import App from "./app.tsx";
import {BrowserRouter} from "react-router-dom";
import './assets/css/style.css';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <BrowserRouter>
    <App />
  </BrowserRouter>,
)