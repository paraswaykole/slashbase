import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import { DBQuery } from '../data/models'
import eventService from '../events/eventService'


export interface DBQueryState {
  [tabId: string]: {
    dbQuery: DBQuery | undefined
  }
}

const initialState: DBQueryState = {}

const createInitialDBQueryState = (state: DBQueryState, tabId: string) => {
  if (state[tabId] === undefined) {
    state[tabId] = {
      dbQuery: undefined,
    }
  }
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
    setDBQuery: (state, { payload }: { payload: { data: DBQuery | undefined, tabId: string } }) => {
      createInitialDBQueryState(state, payload.tabId)
      state[payload.tabId].dbQuery = payload.data
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(getDBQuery.fulfilled, (state, action: any) => {
        createInitialDBQueryState(state, action.meta.arg.tabId)
        if (action.payload.success) {
          state[action.meta.arg.tabId].dbQuery = action.payload.data
        }
      })
  },
})

export const { reset, setDBQuery } = dbQuerySlice.actions

export const selectDBQuery = (state: AppState) => state.tabs.activeTabId ? state.dbQuery[String(state.tabs.activeTabId)]?.dbQuery : undefined


export default dbQuerySlice.reducer