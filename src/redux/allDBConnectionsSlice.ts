import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { DBConnection, Project } from '../data/models'
import apiService from '../network/apiService'

export interface AllDBConnectionsState {
    dbConnections: Array<DBConnection>
    isFetching: boolean
}

const initialState: AllDBConnectionsState = {
  dbConnections: [],
  isFetching: false,
}

export const getAllDBConnections = createAsyncThunk(
  'allDBConnections/getAllDBConnections',
  async () => {
    const result = await apiService.getAllDBConnections()
    const dbConnections = result.success ? result.data : []
    return {
      dbConnections: dbConnections,
    }
  },
  {
    condition: (_, { getState }: any) => {
      const { dbConnections, isFetching} = getState()['allDBConnections'] as AllDBConnectionsState
      const isFetched = dbConnections.length > 0
      if (isFetched || isFetching) {
        return false
      }
      return true
    }
  }
)

export const allDBConnectionSlice = createSlice({
  name: 'allDBConnections',
  initialState,
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(getAllDBConnections.pending, (state) => {
        state.isFetching = true
      })
      .addCase(getAllDBConnections.fulfilled, (state,  action) => {
        state.isFetching = false
        state.dbConnections = state.dbConnections.concat(action.payload.dbConnections)
      })
  },
})

export const selectAllDBConnections = (state: AppState) => state.allDBConnections.dbConnections

export default allDBConnectionSlice.reducer