import '../styles/globals.css'
import '../styles/index.scss'
import '@fortawesome/fontawesome-free/css/all.css'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/duotone-light.css'

import type { AppProps } from 'next/app'
import { Provider } from 'react-redux'
import store from '../redux/store'
import { useRouter } from 'next/router'
import { useAppDispatch, useAppSelector } from '../redux/hooks'
import { useEffect } from 'react'
import Constants from '../constants'
import { getProjects } from '../redux/projectsSlice'
import { getAllDBConnections } from '../redux/allDBConnectionsSlice'
import { Toaster } from 'react-hot-toast'
import { getConfig } from '../redux/configSlice'

function SlashbaseApp({ Component, pageProps }: AppProps) {
  return <Provider store={store}>
    <SlashbaseAppComponent>
      <Component {...pageProps} />
      <Toaster />
    </SlashbaseAppComponent>
  </Provider>
}


const SlashbaseAppComponent = ({ children }: any) => {
  const router = useRouter()
  const dispatch = useAppDispatch()

  useEffect(() => {
    // prefetch or preload data
    dispatch(getProjects())
    dispatch(getAllDBConnections({}))
    dispatch(getConfig())
  }, [dispatch])

  // SPA redirectes (forcing index.html) (disabled-SSR)
  const queryKeys = Object.keys(router.query)
  let finalPath = router.route
  for (let i = 0; i < queryKeys.length; i++) {
    const qkey = queryKeys[i]
    finalPath = finalPath.replace(`[${qkey}]`, String(router.query[qkey]))
  }
  if (router.route != '/_error' && typeof window !== 'undefined' && window.location.pathname !== finalPath) {
    router.replace(window.location.href.slice(window.location.origin.length))
    return null
  }

  return children
}

export default SlashbaseApp
