import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { DBConnection, DBDataModel } from '../data/models'
import apiService from '../network/apiService'

export interface DBConnectionState {
    dbConnection?: DBConnection,
    dbDataModels: DBDataModel[]
}

const initialState: DBConnectionState = {
  dbConnection: undefined,
  dbDataModels: []
}

export const getDBConnection = createAsyncThunk(
  'dbConnection/getDBConnection',
  async (payload: {dbConnId: string}, { rejectWithValue }) => {
    const result = await apiService.getSingleDBConnection(payload.dbConnId)
    if(result.success){
      const dbConnection = result.data
      return {
        dbConnection: dbConnection,
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
)

export const projectsSlice = createSlice({
  name: 'dbConnection',
  initialState,
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(getDBConnection.fulfilled, (state,  action) => {
        state.dbConnection = action.payload.dbConnection
      })
      .addCase(getDBDataModels.fulfilled, (state,  action) => {
        state.dbDataModels = action.payload.dataModels
      })
  },
})

export const selectDBConnection = (state: AppState) => state.dbConnection.dbConnection
export const selectDBDataModels = (state: AppState) => state.dbConnection.dbDataModels

export default projectsSlice.reducer