import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import apiService from '../network/apiService'

import type { AppState } from './store'

export interface APIState {
  version: string
}

const initialState: APIState = {
  version: ''
}

export const healthCheck = createAsyncThunk(
  'api/healthCheck',
  async () => {
    const response = await apiService.getHealthCheck()
    return {
      version: response.version,
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
      .addCase(healthCheck.fulfilled, (state, action: any) => {
        state.version = action.payload.version
      })
  },
})

export const { reset } = apiSlice.actions

export const selectAPIVersion = (state: AppState) => state.api.version

export default apiSlice.reducer