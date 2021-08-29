import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { DBConnection, DBDataModel } from '../data/models'
import apiService from '../network/apiService'

export interface DBConnectionState {
    dbConnection?: DBConnection,
    dbDataModels: DBDataModel[],
    isDBDataModelsFetched: boolean,
}

const initialState: DBConnectionState = {
  dbConnection: undefined,
  dbDataModels: [],
  isDBDataModelsFetched: false,
}

export const getDBConnection = createAsyncThunk(
  'dbConnection/getDBConnection',
  async (payload: {dbConnId: string}, { rejectWithValue, getState }) => {
    const { dbConnection } = (getState() as any)['dbConnection'] as DBConnectionState
    if (dbConnection && dbConnection.id === payload.dbConnId){
      return {
        dbConnection: dbConnection,
        new: false 
      }
    }
    const result = await apiService.getSingleDBConnection(payload.dbConnId)
    if(result.success){
      const dbConnection = result.data
      return {
        dbConnection: dbConnection,
        new: true
      }
    } else {
      return rejectWithValue(result.error)
    }
  },
)

export const getDBDataModels = createAsyncThunk(
  'dbConnection/getDBDataModels',
  async (payload: {dbConnId: string}, { rejectWithValue }) => {
    const result = await apiService.getDBDataModelsByConnectionId(payload.dbConnId)
    if(result.success){
      const dataModels = result.data
      return {
        dataModels: dataModels,
      }
    } else {
      return rejectWithValue(result.error)
    }
  },
  {
    condition: (_, { getState }: any) => {
      const { isDBDataModelsFetched } = getState()['dbConnection'] as DBConnectionState
      return !isDBDataModelsFetched
    }
  }
)

export const dbConnectionSlice = createSlice({
  name: 'dbConnection',
  initialState,
  reducers: {
    reset: (state) => initialState
  },
  extraReducers: (builder) => {
    builder
      .addCase(getDBConnection.fulfilled, (state,  action) => {
        if (action.payload.new){
          state.dbDataModels = []
          state.isDBDataModelsFetched = false
        }
        state.dbConnection = action.payload.dbConnection
      })
      .addCase(getDBDataModels.fulfilled, (state,  action) => {
        state.dbDataModels = action.payload.dataModels
        state.isDBDataModelsFetched = true
      })
  },
})

export const { reset } = dbConnectionSlice.actions

export const selectDBConnection = (state: AppState) => state.dbConnection.dbConnection
export const selectDBDataModels = (state: AppState) => state.dbConnection.dbDataModels

export default dbConnectionSlice.reducer