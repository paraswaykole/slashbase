import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { DBConnection, DBDataModel, DBQuery } from '../data/models'
import eventService from '../events/eventService'

export interface DBConnectionState {
  dbConnection?: DBConnection
  dbDataModels: DBDataModel[]
  dbQueries: DBQuery[]
  isFetchingDBDataModels: boolean
  isDBDataModelsFetched: boolean
  isDBQueriesFetched: boolean
}

const initialState: DBConnectionState = {
  dbConnection: undefined,
  dbDataModels: [],
  dbQueries: [],
  isFetchingDBDataModels: false,
  isDBDataModelsFetched: false,
  isDBQueriesFetched: false,
}

export const getDBConnection = createAsyncThunk(
  'dbConnection/getDBConnection',
  async (payload: { dbConnId: string }, { rejectWithValue, getState }: any) => {
    const { dbConnection } = (getState() as any)['dbConnection'] as DBConnectionState
    if (dbConnection && dbConnection.id === payload.dbConnId) {
      return {
        dbConnection: dbConnection,
        new: false
      }
    }
    const result = await eventService.getSingleDBConnection(payload.dbConnId)
    if (result.success) {
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
  async (payload: { dbConnId: string }, { rejectWithValue, getState }: any) => {
    const result = await eventService.getDBDataModelsByConnectionId(payload.dbConnId)
    if (result.success) {
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

export const getDBQueries = createAsyncThunk(
  'dbConnection/getDBQueries',
  async (payload: { dbConnId: string }, { rejectWithValue, getState }: any) => {
    const result = await eventService.getDBQueriesInDBConn(payload.dbConnId)
    if (result.success) {
      const dbQueries = result.data
      return {
        dbQueries: dbQueries,
      }
    } else {
      return rejectWithValue(result.error)
    }
  },
  {
    condition: (_, { getState }: any) => {
      const { isDBQueriesFetched } = getState()['dbConnection'] as DBConnectionState
      return !isDBQueriesFetched
    }
  }
)

export const saveDBQuery = createAsyncThunk<{ dbQuery: DBQuery }, { dbConnId: string, queryId: string, name: string, query: string }>(
  'dbConnection/saveDBQuery',
  async (payload, { rejectWithValue }: any) => {
    const result = await eventService.saveDBQuery(payload.dbConnId, payload.name, payload.query, payload.queryId)
    if (result.success) {
      const dbQuery = result.data
      return {
        dbQuery: dbQuery,
      }
    } else {
      return rejectWithValue(result.error)
    }
  }
)

export const deleteDBQuery = createAsyncThunk<{ queryId: string }, { queryId: string }>(
  'dbConnection/deleteDBQuery',
  async (payload, { rejectWithValue }: any) => {
    const result = await eventService.deleteDBQuery(payload.queryId)
    if (result.success) {
      return { queryId: payload.queryId }
    } else {
      return rejectWithValue(result.error)
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
      .addCase(getDBConnection.fulfilled, (state, action: any) => {
        if (action.payload.new) {
          state.dbDataModels = []
          state.dbQueries = []
          state.isDBDataModelsFetched = false
          state.isDBQueriesFetched = false
        }
        state.dbConnection = action.payload.dbConnection
      })
      .addCase(getDBDataModels.fulfilled, (state, action: any) => {
        state.dbDataModels = action.payload.dataModels
        state.isDBDataModelsFetched = true
        state.isFetchingDBDataModels = false
      })
      .addCase(getDBDataModels.rejected, (state, action: any) => {
        state.isFetchingDBDataModels = false
      })
      .addCase(getDBDataModels.pending, (state, action: any) => {
        state.isFetchingDBDataModels = true
      })
      .addCase(getDBQueries.fulfilled, (state, action: any) => {
        state.dbQueries = action.payload.dbQueries
        state.isDBQueriesFetched = true
      })
      .addCase(saveDBQuery.fulfilled, (state, action: any) => {
        const idx = state.dbQueries.findIndex(x => x.id === action.payload.dbQuery.id)
        if (idx === -1) {
          state.dbQueries.push(action.payload.dbQuery)
        } else {
          state.dbQueries[idx] = action.payload.dbQuery
        }
      })
      .addCase(deleteDBQuery.fulfilled, (state, action: any) => {
        state.dbQueries = state.dbQueries.filter(x => x.id !== action.payload.queryId)
      })
  },
})

export const { reset } = dbConnectionSlice.actions

export const selectDBConnection = (state: AppState) => state.dbConnection.dbConnection
export const selectDBDataModels = (state: AppState) => state.dbConnection.dbDataModels
export const selectDBDQueries = (state: AppState) => state.dbConnection.dbQueries
export const selectIsFetchingDBDataModels = (state: AppState) => state.dbConnection.isFetchingDBDataModels

export default dbConnectionSlice.reducer