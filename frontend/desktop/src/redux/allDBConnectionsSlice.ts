import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { DBConnection } from '../data/models'
import { AddDBConnPayload } from '../events/payloads'
import eventService from '../events/eventService'

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
  async (payload: { force?: boolean }) => {
    const result = await eventService.getAllDBConnections()
    const dbConnections = result.success ? result.data : []
    return {
      force: payload?.force ?? false,
      dbConnections: dbConnections,
    }
  },
  {
    condition: (payload: { force?: boolean }, { getState }: any) => {
      if (payload?.force === true) {
        return true
      }
      const { dbConnections, isFetching } = getState()['allDBConnections'] as AllDBConnectionsState
      const isFetched = dbConnections.length > 0
      if (isFetched || isFetching) {
        return false
      }
      return true
    }
  }
)

export const addNewDBConn = createAsyncThunk(
  'allDBConnections/addNewDBConn',
  async (payload: AddDBConnPayload, { rejectWithValue, getState }: any) => {
    const response = await eventService.addNewDBConn(payload)
    if (response.success) {
      const dbConn = response.success ? response.data : null
      return {
        dbConn: dbConn
      }
    } else {
      return rejectWithValue(response.error)
    }
  }
)


export const allDBConnectionSlice = createSlice({
  name: 'allDBConnections',
  initialState,
  reducers: {
    reset: (state) => initialState
  },
  extraReducers: (builder) => {
    builder
      .addCase(getAllDBConnections.pending, (state) => {
        state.isFetching = true
      })
      .addCase(getAllDBConnections.fulfilled, (state, action) => {
        state.isFetching = false
        if (action.payload.force) {
          state.dbConnections = action.payload.dbConnections
        } else {
          state.dbConnections = state.dbConnections.concat(action.payload.dbConnections)
        }
      })
      .addCase(addNewDBConn.fulfilled, (state, action: any) => {
        if (action.payload.dbConn)
          state.dbConnections.push(action.payload.dbConn)
      })
  },
})

export const { reset } = allDBConnectionSlice.actions

export const selectAllDBConnections = (state: AppState) => state.allDBConnections.dbConnections

export default allDBConnectionSlice.reducer