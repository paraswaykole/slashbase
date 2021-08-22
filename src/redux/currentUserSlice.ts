import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { User } from '../data/models'
import storage from '../data/storage'
import apiService from '../network/apiService'

export interface CurrentUserState {
  user?: User
  isAuthenticated: boolean | null
}

const initialState: CurrentUserState = {
  isAuthenticated: null
}

export const getUser = createAsyncThunk(
  'currentUser/getUser',
  async () => {
    const isAuthenticated = await storage.isUserAuthenticated()
    const currentUser = await storage.getCurrentUser()
    return {
      currentUser,
      isAuthenticated
    }
  }
)

export const loginUser = createAsyncThunk(
  'currentUser/loginUser',
  async (payload: {email: string, password: string}, { rejectWithValue }) => {
    let response = await apiService.loginUser(payload.email, payload.password)
    if (response.success) {
      await storage.loginCurrentUser(response.data.user, response.data.token)
      return {
        currentUser: response.data.user,
        isAuthenticated: true,
      }
    } else {
      return rejectWithValue(response.error)
    }
  }
)

export const logoutUser = createAsyncThunk(
  'currentUser/logoutUser',
  async () => {
    //TODO: make logout api call
    await storage.logoutUser()
    return {
      currentUser: undefined,
      isAuthenticated: false,
    }
  }
)

export const userSlice = createSlice({
  name: 'currentUser',
  initialState,
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(getUser.fulfilled, (state, action) => {
        state.user = action.payload.currentUser ? action.payload.currentUser : undefined
        state.isAuthenticated = action.payload.isAuthenticated
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        if (action.payload){
          state.user = action.payload.currentUser
          state.isAuthenticated = action.payload.isAuthenticated
        }
      })
      .addCase(logoutUser.fulfilled, (state, action) => {
        state.user = action.payload.currentUser
        state.isAuthenticated = action.payload.isAuthenticated
      })
  },
})


export const { } = userSlice.actions

export const selectCurrentUser = (state: AppState) => state.currentUser.user!

export const selectIsAuthenticated = (state: AppState) => state.currentUser.isAuthenticated

export default userSlice.reducer