import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import { DBQueryLog } from '../data/models'
import eventService from '../events/eventService'


export interface DBHistoryState {
  dbQueryLogs: DBQueryLog[],
  dbQueryLogsNext: number | undefined
}

const initialState: DBHistoryState = {
  dbQueryLogs: [],
  dbQueryLogsNext: undefined
}

export const getDBQueryLogs = createAsyncThunk(
  'dbHistory/getDBQueryLogs',
  async (payload: any, { getState }: any) => {
    const { dbQueryLogsNext } = getState()['dbHistory'] as DBHistoryState
    const { dbConnId } = payload
    const result = await eventService.getDBHistory(dbConnId, dbQueryLogsNext)
    return result
  }
)

export const dbHistorySlice = createSlice({
  name: 'dbHistory',
  initialState,
  reducers: {
    reset: () => initialState,
  },
  extraReducers: (builder) => {
    builder
      .addCase(getDBQueryLogs.fulfilled, (state, action: any) => {
        if (action.payload.success) {
          state.dbQueryLogs = state.dbQueryLogs.concat(...action.payload.data.list)
          state.dbQueryLogsNext = action.payload.data.next
        }
      })
  },
})

export const { reset } = dbHistorySlice.actions

export const selectDBQueryLogs = (state: AppState) => state.dbHistory.dbQueryLogs

export const selectDBQueryLogsNext = (state: AppState) => state.dbHistory.dbQueryLogsNext


export default dbHistorySlice.reducer