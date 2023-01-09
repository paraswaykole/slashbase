import './styles/globals.css'
import './styles/index.scss'
import '@fortawesome/fontawesome-free/css/all.css'
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import { BrowserRouter } from "react-router-dom"
import { Provider } from 'react-redux'
import store from './redux/store'
import posthog from 'posthog-js'

posthog.init(
  String(process.env.REACT_APP_POSTHOG_KEY),
  {
    api_host: process.env.REACT_APP_POSTHOG_API_HOST,
    capture_pageview: false
  }
)

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
)
root.render(
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Provider>
  </React.StrictMode>
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
// reportWebVitals();
