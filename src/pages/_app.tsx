import '../styles/globals.css'
import '../styles/index.scss'
import '@fortawesome/fontawesome-free/css/all.css'

import type { AppProps } from 'next/app'
import { Provider } from 'react-redux'
import store from '../redux/store'
import { useRouter } from 'next/router'
import { useAppDispatch, useAppSelector } from '../redux/hooks'
import { getUser, selectIsAuthenticated } from '../redux/currentUserSlice'
import { useEffect } from 'react'
import Constants from '../constants'

function SlashbaseApp({ Component, pageProps }: AppProps) {
  return <Provider store={store}>
    <SlashbaseAppComponent>
      <Component {...pageProps} />
    </SlashbaseAppComponent>
  </Provider>
}


const SlashbaseAppComponent = ({children}: any) => {
  const router = useRouter()
  const dispatch = useAppDispatch()
  const isAuthenticated: boolean|null = useAppSelector(selectIsAuthenticated)

  useEffect(() => {
      (async () => {
          const currentPath = Object.values(Constants.APP_PATHS).find(x => x.as === router.asPath)
          if (currentPath){
            const { payload } : any = await dispatch((getUser()))
            if((isAuthenticated === null && payload.isAuthenticated) || !currentPath.isAuth || isAuthenticated){
                return
            }
          }
          if(router.route != '/_error')
            router.replace(Constants.APP_PATHS.LOGIN.as)
      })()
  }, [dispatch, isAuthenticated])
  
  return children
}

export default SlashbaseApp
