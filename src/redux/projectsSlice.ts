import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { Project } from '../data/models'
import apiService from '../network/apiService'

export interface ProjectState {
    projects: Array<Project>
    isFetching: boolean
}

const initialState: ProjectState = {
  projects: [],
  isFetching: false,
}

export const getProjects = createAsyncThunk(
  'projects/getProjects',
  async () => {
    const result = await apiService.getProjects()
    const projects = result.success ? result.data : []
    return {
      projects: projects,
    }
  },
  {
    condition: (_, { getState }: any) => {
      const { projects, isFetching} = getState()['projects'] as ProjectState
      const isFetched = projects.length > 0
      if (isFetched || isFetching) {
        return false
      }
      return true
    }
  }
)

export const projectsSlice = createSlice({
  name: 'projects',
  initialState,
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(getProjects.pending, (state) => {
        state.isFetching = true
      })
      .addCase(getProjects.fulfilled, (state,  action) => {
        state.isFetching = false
        state.projects = state.projects.concat(action.payload.projects)
      })
  },
})

export const selectProjects = (state: AppState) => state.projects.projects

export default projectsSlice.reducer