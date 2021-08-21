import '../styles/globals.css'
import '../styles/index.scss'
import '@fortawesome/fontawesome-free/css/all.css'

import type { AppProps } from 'next/app'
import { Provider } from 'react-redux'
import store from '../redux/store'

function MyApp({ Component, pageProps }: AppProps) {
  return <Provider store={store}>
    <Component {...pageProps} />
  </Provider>
}
export default MyApp
