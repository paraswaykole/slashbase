import './styles/globals.css'
import './styles/index.scss'
import '@fortawesome/fontawesome-free/css/all.css'
import React from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'
import { BrowserRouter } from "react-router-dom"
import { Provider } from 'react-redux'
import store from './redux/store'
import product from './lib/product'

product.posthogInit()

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
