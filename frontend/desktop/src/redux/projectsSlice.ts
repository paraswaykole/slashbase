import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import type { AppState } from './store'
import { DBConnection, Project } from '../data/models'
import { getAllDBConnections } from './allDBConnectionsSlice'
import { toast } from 'react-hot-toast'
import eventService from '../events/eventService'

export interface ProjectState {
  projects: Array<Project>
  isFetching: boolean
  dbConnectionsInProject: Array<DBConnection>
}

const initialState: ProjectState = {
  projects: [],
  isFetching: false,
  dbConnectionsInProject: [],
}

export const getProjects = createAsyncThunk(
  'projects/getProjects',
  async (_, { }: any) => {
    const result = await eventService.getProjects()
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
  async (payload: { projectName: string }, { }: any) => {
    if (payload.projectName.trim().length === 0) {
      toast.error("Project Name cannot be empty!");
      return {
        success: false,
        project: null,
      }
    }
    const result = await eventService.createNewProject(payload.projectName)
    console.log(result)
    const project = result.success ? result.data : null
    return {
      success: true,
      project: project,
    }
  }
)

export const deleteProject = createAsyncThunk(
  'projects/deleteProject',
  async (payload: { projectId: string }, { dispatch }: any) => {
    const result = await eventService.deleteProject(payload.projectId)
    if (result.success) {
      await dispatch(getAllDBConnections({ force: true }))
      return {
        success: true,
        projectId: payload.projectId,
      }
    } else {
      return {
        success: false,
        projectId: ''
      }
    }
  }
)

export const getDBConnectionsInProjects = createAsyncThunk(
  'projects/getDBConnectionsInProjects',
  async (payload: { projectId: string }, { }: any) => {
    const result = await eventService.getDBConnectionsByProject(payload.projectId)
    const dbConnections = result.success ? result.data : []
    return {
      dbConnectionsInProject: dbConnections,
    }
  }
)

export const deleteDBConnectionInProject = createAsyncThunk(
  'projects/deleteDBConnectionInProject',
  async (payload: { dbConnId: string }, { dispatch }: any) => {
    const result = await eventService.deleteDBConnection(payload.dbConnId)
    if (result.success) {
      await dispatch(getAllDBConnections({ force: true }))
      return {
        success: true,
        dbConnId: payload.dbConnId,
      }
    } else {
      return {
        success: false,
        dbConnId: ''
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
      .addCase(getDBConnectionsInProjects.fulfilled, (state, action) => {
        state.dbConnectionsInProject = action.payload.dbConnectionsInProject
      })
      .addCase(deleteDBConnectionInProject.fulfilled, (state, action) => {
        if (action.payload.success) {
          state.dbConnectionsInProject = state.dbConnectionsInProject.filter(dbConn => dbConn.id !== action.payload.dbConnId)
        }
      })
  },
})

export const { reset } = projectsSlice.actions

export const selectProjects = (state: AppState) => state.projects.projects

export const selectCurrentProject = (state: AppState) => state.projects.projects.find(x => x.id === state.dbConnection.dbConnection?.projectId)

export const selectDBConnectionsInProject = (state: AppState) => state.projects.dbConnectionsInProject

export interface ProjectPermissions {
  readOnly: boolean
}

export default projectsSlice.reducer