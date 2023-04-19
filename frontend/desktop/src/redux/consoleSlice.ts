import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import eventService from '../events/eventService'
import { stat } from 'fs'

export interface OutputBlock {
    text: string
    cmd: boolean
}

export interface ConsoleState {
    blocks: Array<OutputBlock>
    dbConnectionId: string | undefined
}

const initialState: ConsoleState = {
    blocks: [],
    dbConnectionId: undefined
}

export const runConsoleCmd = createAsyncThunk(
    'tabs/runConsoleCmd',
    async (payload: { dbConnId: string, cmdString: string }, { rejectWithValue, getState }: any) => {
        const dbConnectionId = payload.dbConnId
        const cmdString = payload.cmdString
        const result = await eventService.runConsoleCommand(dbConnectionId, cmdString)
        if (result.success) {
            return {
                text: result.data,
            }
        } else {
            return rejectWithValue(result.error)
        }
    }
)

export const consoleSlice = createSlice({
    name: 'console',
    initialState,
    reducers: {
        reset: () => initialState,
        initConsole: (state, { payload }: { payload: string }) => {
            if (state.dbConnectionId && state.dbConnectionId !== payload) {
                state.blocks = []
                state.dbConnectionId = payload
            }
        }
    },
    extraReducers: (builder) => {
        builder
            .addCase(runConsoleCmd.fulfilled, (state, action) => {
                state.blocks.push({
                    text: action.payload.text,
                    cmd: false
                })
            })
            .addCase(runConsoleCmd.pending, (state, action) => {
                state.dbConnectionId = action.meta.arg.dbConnId
                state.blocks.push({
                    text: action.meta.arg.cmdString,
                    cmd: true
                })
            })
    },
})


export const { reset, initConsole } = consoleSlice.actions

export const selectBlocks = (state: AppState) => state.console.blocks

export default consoleSlice.reducer