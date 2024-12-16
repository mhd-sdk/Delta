import '@carbon/charts-react/styles.css';
import React from 'react';
import { createRoot } from 'react-dom/client';
import App from './App';
import './App.css';
import './main.scss';
import '/node_modules/react-grid-layout/css/styles.css';
import '/node_modules/react-resizable/css/styles.css';

const container = document.getElementById('root');

const root = createRoot(container!);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
