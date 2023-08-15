import React from 'react';
import ReactDOM from 'react-dom/client';
import { Toaster } from 'react-hot-toast';

import App from './App.tsx';

import { TOAST_DEFAULT_OPTIONS } from '@/config/helpers/toast.helper.ts';
import '@/styles/index.css';
import ThemeProvider from '@/theme/ThemeProvider.tsx';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider>
      <App />
      <Toaster {...TOAST_DEFAULT_OPTIONS} />
    </ThemeProvider>
  </React.StrictMode>,
);
