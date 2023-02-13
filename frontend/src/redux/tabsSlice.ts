import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import { Tab } from '../data/models'
import eventService from '../events/eventService'
import { TabType } from '../data/defaults'
import { actions } from 'react-table'

export interface TabState {
    tabs: Array<Tab>
    activeTabId: String | undefined
}

const initialState: TabState = {
    tabs: [],
    activeTabId: undefined
}

export const createTab = createAsyncThunk(
    'tabs/createTab',
    async (payload: { dbConnId: string, tabType: TabType }, { rejectWithValue }: any) => {
        const dbConnectionId = payload.dbConnId
        const result = await eventService.createTab(dbConnectionId)
        if (result.success) {
            return {
                tab: result.data,
            }
        } else {
            return rejectWithValue(result.error)
        }
    }
)

export const getTabs = createAsyncThunk(
    'tabs/getTabs',
    async (payload: { dbConnId: string }, { rejectWithValue }: any) => {
        const dbConnectionId = payload.dbConnId
        const result = await eventService.getTabsByDBConnection(dbConnectionId)
        console.log(result.data[0].id)
        if (result.success) {
            return {
                tabs: result.data,
                activeTabId: result.data[0].id
            }
        } else {
            return rejectWithValue(result.error)
        }
    }
)

export const closeTab = createAsyncThunk(
    'tabs/closeTab',
    async (payload: { dbConnId: string, tabId: string }, { rejectWithValue }: any) => {
        const dbConnectionId = payload.dbConnId
        const tabId = payload.tabId
        const result = await eventService.closeTab(dbConnectionId, tabId)
        if (result.success) {
            return {
                tabId: tabId
            }
        } else {
            return rejectWithValue(result.error)
        }
    }
)



export const tabsSlice = createSlice({
    name: 'tabs',
    initialState,
    reducers: {
        reset: () => initialState,
    },
    extraReducers: (builder) => {
        builder
            .addCase(getTabs.fulfilled, (state, action) => {
                state.tabs = action.payload.tabs
                state.activeTabId = action.payload.activeTabId
            })
            .addCase(createTab.fulfilled, (state, action) => {
                state.tabs.push(action.payload.tab)
            })
            .addCase(closeTab.fulfilled, (state, action) => {
                state.tabs = state.tabs.filter(t => t.id !== action.payload.tabId)
            })
    },
})


export const { reset } = tabsSlice.actions

export const selectTabs = (state: AppState) => state.tabs.tabs.map(t => ({ ...t, isActive: t.id === state.tabs.activeTabId }))

export default tabsSlice.reducer