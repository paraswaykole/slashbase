import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { AppState } from './store'
import eventService from '../events/eventService'

export interface OutputBlock {
    text: string
    cmd: boolean
}

export interface ConsoleState {
    blocks: Array<OutputBlock>
}

const initialState: ConsoleState = {
    blocks: [],
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
                state.blocks.push({
                    text: action.meta.arg.cmdString,
                    cmd: true
                })
            })
    },
})


export const { reset } = consoleSlice.actions

export const selectBlocks = (state: AppState) => state.console.blocks

export default consoleSlice.reducer