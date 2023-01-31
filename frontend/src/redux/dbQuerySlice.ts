import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import { DBQuery } from '../data/models'
import eventService from '../events/eventService'


export interface DBQueryState {
  dbQuery: DBQuery | undefined
}

const initialState: DBQueryState = {
  dbQuery: undefined,
}

export const getDBQuery = createAsyncThunk(
  'dbQuery/getDBQuery',
  async (payload: any, { }: any) => {
    const { queryId } = payload
    const result = await eventService.getSingleDBQuery(queryId)
    return result
  }
)

export const runQuery = createAsyncThunk(
  'dbQuery/runQuery',
  async (payload: any, { }: any) => {
    const { dbConnectionId, query } = payload
    const result = await eventService.runQuery(dbConnectionId, query)
    return result
  }
)

export const dbQuerySlice = createSlice({
  name: 'dbQuery',
  initialState,
  reducers: {
    reset: () => initialState,
    setDBQuery: (state, { payload }: { payload: DBQuery | undefined }) => {
      state.dbQuery = payload
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(getDBQuery.fulfilled, (state, action: any) => {
        if (action.payload.success) {
          state.dbQuery = action.payload.data
        }
      })
  },
})

export const { reset, setDBQuery } = dbQuerySlice.actions

export const selectDBQuery = (state: AppState) => state.dbQuery.dbQuery


export default dbQuerySlice.reducer