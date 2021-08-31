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
import { getUser, selectIsAuthenticated } from '../redux/currentUserSlice'
import { useEffect } from 'react'
import Constants from '../constants'
import { getProjects } from '../redux/projectsSlice'
import { getAllDBConnections } from '../redux/allDBConnectionsSlice'
import { Toaster } from 'react-hot-toast'

function SlashbaseApp({ Component, pageProps }: AppProps) {
  return <Provider store={store}>
    <SlashbaseAppComponent>
      <Component {...pageProps} />
      <Toaster />
    </SlashbaseAppComponent>
  </Provider>
}


const SlashbaseAppComponent = ({children}: any) => {
  const router = useRouter()
  const dispatch = useAppDispatch()
  const isAuthenticated: boolean|null = useAppSelector(selectIsAuthenticated)

  useEffect(() => {
      (async () => {
          const currentPath = Object.values(Constants.APP_PATHS).find(x => x.path === router.route)
          if (currentPath){
            const { payload } : any = await dispatch((getUser()))
            if((isAuthenticated === null && payload.isAuthenticated) || !currentPath.isAuth || isAuthenticated){
                return
            }
          }
          if(router.route != '/_error')
            router.replace(Constants.APP_PATHS.LOGIN.path)
      })()
      // prefetch or preload data
      if (isAuthenticated){
        dispatch(getProjects())
        dispatch(getAllDBConnections())
      }
  }, [dispatch, isAuthenticated])

  // SPA redirectes (forcing index.html) (disabled-SSR)
  const queryKeys = Object.keys(router.query)
  let finalPath = router.route
  for(let i=0; i<queryKeys.length; i++){
    const qkey = queryKeys[i]
    finalPath = finalPath.replace(`[${qkey}]`, String(router.query[qkey]))
  }
  if(typeof window !== 'undefined' && window.location.pathname !== finalPath){
    router.replace(window.location.href.slice(window.location.origin.length))
    return null
  }
  
  return children
}

export default SlashbaseApp
