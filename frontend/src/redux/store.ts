import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit'
import Constants from '../constants'

import projectsReducer from './projectsSlice'
import dbConnectionReducer from './dbConnectionSlice'
import allDBConnectionsReducer from './allDBConnectionsSlice'
import configReducer from './configSlice'

export function makeStore() {
  return configureStore({
    reducer: {
      projects: projectsReducer,
      dbConnection: dbConnectionReducer,
      allDBConnections: allDBConnectionsReducer,
      config: configReducer,
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