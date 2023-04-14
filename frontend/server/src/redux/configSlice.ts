import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import type { AppState } from './store'
import storage from '../data/storage'

export interface ConfigState {
  isShowingSidebar: boolean | null,
}

const initialState: ConfigState = {
  isShowingSidebar: null,
}

export const getConfig = createAsyncThunk(
  'config/getConfig',
  async () => {
    const isShowingSidebar = await storage.isShowingSidebar()
    return {
      isShowingSidebar
    }
  }
)


export const configSlice = createSlice({
  name: 'config',
  initialState,
  reducers: {
    reset: () => initialState,
    setIsShowingSidebar: (state, {payload}: {payload: boolean}) => {
      state.isShowingSidebar = payload
      storage.setIsShowingSidebar(payload)
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(getConfig.fulfilled, (state, action) => {
        state.isShowingSidebar = action.payload.isShowingSidebar
      })
  },
})


export const { reset, setIsShowingSidebar } = configSlice.actions

export const selectIsShowingSidebar = (state: AppState) => state.config.isShowingSidebar ?? true

export default configSlice.reducer