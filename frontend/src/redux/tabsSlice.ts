import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import { Tab } from '../data/models'
import eventService from '../events/eventService'

export interface TabState {
    tabs: Array<Tab>
}

const initialState: TabState = {
    tabs: [],
}

export const getTabs = createAsyncThunk(
    'tabs/getTabs',
    async (payload: { dbConnId: string }, { rejectWithValue }: any) => {
        const dbConnectionId = payload.dbConnId
        const result = await eventService.getTabsByDBConnection(dbConnectionId)
        if (result.success) {
            return {
                tabs: result.data
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
            })
    },
})


export const { reset } = tabsSlice.actions

export const selectTabs = (state: AppState) => state.tabs.tabs

export default tabsSlice.reducer