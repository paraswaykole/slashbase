import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit'
import Constants from '../constants'

import currentUserReducer from './currentUserSlice'
import projectsReducer from './projectsSlice'

export function makeStore() {
  return configureStore({
    reducer: {
      currentUser: currentUserReducer,
      projects: projectsReducer,
    },
    devTools: !Constants.IS_LIVE
  })
}

const store = makeStore()

export type AppState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  AppState,
  unknown,
  Action<string>
>

export default store