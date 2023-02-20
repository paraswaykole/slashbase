import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import { Tab } from '../data/models'
import eventService from '../events/eventService'
import { TabType } from '../data/defaults'
import { DBConnectionState } from './dbConnectionSlice'

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
    async (payload: { dbConnId: string, tabType: TabType, metadata?: any | undefined }, { rejectWithValue }: any) => {
        const dbConnectionId = payload.dbConnId
        const tabType = payload.tabType
        let mSchema = ""
        let mName = ""
        let queryId = ""
        if (tabType === TabType.DATA || tabType === TabType.MODEL) {
            mSchema = payload.metadata.schema
            mName = payload.metadata.name
        } else if (tabType === TabType.QUERY) {
            queryId = payload.metadata.queryId
        }
        const result = await eventService.createTab(dbConnectionId, tabType, mSchema, mName, queryId)
        if (result.success) {
            return {
                tab: result.data,
            }
        } else {
            return rejectWithValue(result.error)
        }
    }
)

export const updateActiveTab = createAsyncThunk(
    'tabs/updateActiveTab',
    async (payload: { tabType: TabType, metadata: Object }, { getState, rejectWithValue }: any) => {
        const { activeTabId } = getState()['tabs'] as TabState
        const { dbConnection } = getState()['dbConnection'] as DBConnectionState
        if (!activeTabId) {
            return rejectWithValue('no active tab')
        }
        if (!dbConnection) {
            return rejectWithValue('no db connection active')
        }
        const tabType = payload.tabType
        const metadata = payload.metadata
        const result = await eventService.updateTab(dbConnection.id, String(activeTabId), tabType, metadata)
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
        setActiveTab: (state, { payload }: { payload: string }) => {
            state.activeTabId = payload
        },
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
            .addCase(updateActiveTab.fulfilled, (state, action) => {
                const idx = state.tabs.findIndex(t => t.id === action.payload.tab.id)
                state.tabs[idx] = action.payload.tab
            })
            .addCase(closeTab.fulfilled, (state, action) => {
                state.tabs = state.tabs.filter(t => t.id !== action.payload.tabId)
            })
    },
})


export const { reset, setActiveTab } = tabsSlice.actions

export const selectTabs = (state: AppState) => state.tabs.tabs.map(t => ({ ...t, isActive: t.id === state.tabs.activeTabId }))

export const selectActiveTab = (state: AppState) => state.tabs.tabs.find(t => t.id === state.tabs.activeTabId)!

export default tabsSlice.reducer