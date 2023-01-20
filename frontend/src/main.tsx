import './styles/globals.css'
import './styles/index.scss'
import '@fortawesome/fontawesome-free/css/all.css'
import React from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'
import { BrowserRouter } from "react-router-dom"
import { Provider } from 'react-redux'
import store from './redux/store'
import posthog from 'posthog-js'

// posthog.init(
//     String(process.env.REACT_APP_POSTHOG_KEY),
//     {
//         api_host: process.env.REACT_APP_POSTHOG_API_HOST,
//         capture_pageview: false
//     }
// )

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <React.StrictMode>
        <Provider store={store}>
            <BrowserRouter>
                <App />
            </BrowserRouter>
        </Provider>
    </React.StrictMode>
)
