import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { Project } from '../data/models'
import apiService from '../network/apiService'
import { AddProjectMemberPayload } from '../network/payloads'
import { getAllDBConnections } from './allDBConnectionsSlice'
import Constants from '../constants'

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
      const { projects, isFetching } = getState()['projects'] as ProjectState
      const isFetched = projects.length > 0
      if (isFetched || isFetching) {
        return false
      }
      return true
    }
  }
)

export const createNewProject = createAsyncThunk(
  'projects/createNewProject',
  async (payload: { projectName: string }) => {
    const result = await apiService.createNewProject(payload.projectName)
    const project = result.success ? result.data : null
    return {
      project: project,
    }
  }
)

export const deleteProject = createAsyncThunk(
  'projects/deleteProject',
  async (payload: { projectId: string }, { dispatch }) => {
    const result = await apiService.deleteProject(payload.projectId)
    if (result.success) {
      const resp = await dispatch(getAllDBConnections({ force: true }))
      console.log(resp)
      return {
        success: true,
        projectId: payload.projectId,
      }
    } else {
      return {
        success: false
      }
    }
  }
)


export const projectsSlice = createSlice({
  name: 'projects',
  initialState,
  reducers: {
    reset: (state) => initialState
  },
  extraReducers: (builder) => {
    builder
      .addCase(getProjects.pending, (state) => {
        state.isFetching = true
      })
      .addCase(getProjects.fulfilled, (state, action) => {
        state.isFetching = false
        state.projects = state.projects.concat(action.payload.projects)
      })
      .addCase(createNewProject.fulfilled, (state, action) => {
        if (action.payload.project) {
          state.projects.push(action.payload.project)
        }
      })
      .addCase(deleteProject.fulfilled, (state, action) => {
        if (action.payload.success) {
          state.projects = state.projects.filter(pro => pro.id !== action.payload.projectId)
        }
      })
  },
})

export const { reset } = projectsSlice.actions

export const selectProjects = (state: AppState) => state.projects.projects

export const selectCurrentProject = (state: AppState) => state.projects.projects.find(x => x.id === state.dbConnection.dbConnection?.projectId)

export interface ProjectPermissions {
  readOnly: boolean
}

export default projectsSlice.reducer