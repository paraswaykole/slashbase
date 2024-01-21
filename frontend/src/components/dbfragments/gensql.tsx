import React, { useContext, useState } from "react"
import { Tab } from "../../data/models"
import { selectDBConnection } from "../../redux/dbConnectionSlice"
import { useAppDispatch, useAppSelector } from "../../redux/hooks"
import TabContext from "../layouts/tabcontext"
import styles from './gensql.module.scss'
import apiService from "../../network/apiService"
import { duotoneLight } from '@uiw/codemirror-theme-duotone'
import { toast } from "react-hot-toast"
import ReactCodeMirror from "@uiw/react-codemirror"
import { sql } from '@codemirror/lang-sql'
import { TabType } from "../../data/defaults"
import { createTab } from "../../redux/tabsSlice"

type DBGenSQLPropType = {
}

const DBGenSQLFragment = ({ }: DBGenSQLPropType) => {

    const dispatch = useAppDispatch()

    const currentTab: Tab = useContext(TabContext)!

    const [inputValue, setInputValue] = useState<string>('')
    const [generating, setGenerating] = useState<boolean>(false)
    const [outputValue, setOutputValue] = useState<string | undefined>()

    const dbConnection = useAppSelector(selectDBConnection)

    const onChange = React.useCallback((value: any) => {
        setOutputValue(value)
    }, [])

    const runGenerateSQL = async () => {
        if (generating) {
            return
        }
        setGenerating(true)
        const result = await apiService.generateSQL(dbConnection!.id, inputValue)
        if (result.success)
            setOutputValue(result.data)
        else
            toast.error(result.error!)
        setGenerating(false)
    }

    const openInQueryEditor = () => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.QUERY, metadata: { queryId: "new", query: outputValue } }))
    }

    const copyToClipboard = () => {
        navigator.clipboard.writeText(outputValue!)
        toast.success("copied")
    }

    return <div className={styles.console + " " + (currentTab.isActive ? "db-tab-active" : "db-tab")}>
        <div className={"control" + (generating ? " is-loading" : "")}>
            <textarea
                value={inputValue}
                className="textarea"
                placeholder="Enter prompt to generate SQL"
                onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => { setInputValue(e.target.value) }}
            />
        </div>
        <br />
        <div className="control">
            {!generating && <button className={"button" + (outputValue === undefined ? " is-primary" : "")} onClick={runGenerateSQL}>
                <span className="icon is-small">
                    <i className="fas fa-play-circle" aria-hidden="true"></i>
                </span>&nbsp;&nbsp;
                Generate
            </button>}
            {generating && <button className="button is-primary is-loading">Running</button>}
        </div>
        {outputValue !== undefined && <>
            <br /><br />
            <ReactCodeMirror
                value={outputValue}
                extensions={[sql()]}
                theme={duotoneLight}
                height={"auto"}
                minHeight="80px"
                placeholder={"Generated SQL"}
                basicSetup={{
                    autocompletion: false,
                    highlightActiveLine: false,
                }}
                onChange={onChange}
            />
            <br />
            <div className="buttons">
                <button className="button is-primary" onClick={openInQueryEditor}>
                    <span className="icon is-small">
                        <i className="fas fa-edit" aria-hidden="true"></i>
                    </span>&nbsp;&nbsp;
                    Open in Query Editor
                </button>
                <button className="button" onClick={copyToClipboard}>
                    <span className="icon is-small">
                        <i className="fas fa-copy" aria-hidden="true"></i>
                    </span>&nbsp;&nbsp;
                    Copy to clipboard
                </button>
            </div>
        </>}
    </div>
}

export default DBGenSQLFragment