import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import eventService from '../events/eventService'

import type { AppState } from './store'

export interface APIState {
  isConnecting: boolean
  isConnected: boolean
  version: string
}

const initialState: APIState = {
  isConnecting: false,
  isConnected: false,
  version: ''
}

export const connectLocal = createAsyncThunk(
  'api/connectLocal',
  async () => {
    const response = await eventService.getHealthCheck()
    return {
      isConnected: response.success,
      version: response.version,
    }
  },
  {
    condition: (_, { getState }: any) => {
      const { isConnecting } = getState()['api'] as APIState
      return !isConnecting
    }
  }
)


export const apiSlice = createSlice({
  name: 'api',
  initialState,
  reducers: {
    reset: () => initialState,
  },
  extraReducers: (builder) => {
    builder
      .addCase(connectLocal.rejected, (state, action: any) => {
        state.isConnecting = false
      })
      .addCase(connectLocal.pending, (state, action: any) => {
        state.isConnecting = true
      })
      .addCase(connectLocal.fulfilled, (state, action: any) => {
        state.isConnected = action.payload.isConnected
        state.version = action.payload.version
      })
  },
})

export const { reset } = apiSlice.actions

export const selectIsConnected = (state: AppState) => state.api.isConnected

export const selectAPIVersion = (state: AppState) => state.api.version

export default apiSlice.reducer