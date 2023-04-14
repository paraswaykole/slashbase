import styles from './queryeditor.module.scss'
import 'react-tooltip/dist/react-tooltip.css'
import React, { useState, useRef, useEffect, useContext } from 'react'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { deleteDBQuery, saveDBQuery, selectDBConnection } from '../../../redux/dbConnectionSlice'
import { DBConnection, Tab } from '../../../data/models'
import toast from 'react-hot-toast'
import { format } from 'sql-formatter'
import { DBConnType } from '../../../data/defaults'
import { js_beautify } from 'js-beautify'
import ReactCodeMirror, { ReactCodeMirrorRef } from '@uiw/react-codemirror'
import { duotoneLight } from '@uiw/codemirror-theme-duotone'
import { javascript } from '@codemirror/lang-javascript'
import { sql } from '@codemirror/lang-sql'
import CheatSheetModal from '../cheatsheet/cheatsheet'
import { Tooltip } from 'react-tooltip'
import apiService from '../../../network/apiService'
import TabContext from '../../layouts/tabcontext'

type QueryEditorPropType = {
    initialValue: string,
    initQueryName: string,
    queryId: string,
    dbType: DBConnType
    runQuery: (query: string, callback: () => void) => void
    onSave: (queryId: string, query: string) => void
    onDelete: () => void
}

const QueryEditor = ({ initialValue, initQueryName, queryId, dbType, runQuery, onSave, onDelete }: QueryEditorPropType) => {

    const dispatch = useAppDispatch()

    const [value, setValue] = useState(initialValue)
    const [queryName, setQueryName] = useState(initQueryName)
    const [saving, setSaving] = useState(false)
    const [deleting, setDeleting] = useState(false)
    const [running, setRunning] = useState(false)
    const [showCheatsheet, setShowCheatsheet] = useState<boolean>(false)
    const editorRef = useRef<ReactCodeMirrorRef | null>(null)

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const currentTab: Tab = useContext(TabContext)!

    const onChange = React.useCallback((value: any) => {
        setValue(value)
    }, [])

    useEffect(() => {
        if (value != initialValue) {
            apiService.updateTab(dbConnection!.id, currentTab.id, currentTab.type, { queryId: currentTab.metadata.queryId, query: value })
        }
    }, [value])

    const handleKeyDown = (event: React.KeyboardEvent) => {
        if (event.ctrlKey && event.key.toLocaleLowerCase() === 'enter') {
            startRunningQuery()
        }
    }

    const startSaving = async () => {
        if (queryName === '') {
            return
        }
        setSaving(true)
        try {
            const result = await dispatch(saveDBQuery({ dbConnId: dbConnection!.id, queryId, name: queryName, query: value })).unwrap()
            toast.success("Saved Succesfully!")
            onSave(result.dbQuery.id, result.dbQuery.query)
        } catch (e) {
            toast.error("There was some problem saving! Please try again.")
        }
        setSaving(false)
    }

    const startDeleting = async () => {
        if (queryId === 'new') {
            return
        }
        setDeleting(true)
        try {
            await dispatch(deleteDBQuery({ queryId })).unwrap()
            toast.success("Query Deleted")
            onDelete()
        } catch (e) {
            toast.error("There was some problem saving! Please try again.")
        }
        setDeleting(false)
    }

    const startRunningQuery = () => {
        runQuery(value, () => {
            setRunning(false)
        })
    }

    const formatQuery = () => {
        let formattedQuery: string = value
        if (dbType === DBConnType.POSTGRES) {
            formattedQuery = format(value, {
                language: "postgresql",
                keywordCase: 'upper',
                linesBetweenQueries: 2,
            })
        } else if (dbType === DBConnType.MONGO) {
            formattedQuery = js_beautify(value)
        } else if (dbType == DBConnType.MYSQL) {
            formattedQuery = format(value, {
                language: "mysql",
                keywordCase: 'upper',
                linesBetweenQueries: 2,
            })
        }
        setValue(formattedQuery)
    }

    const placeholderText = (dbType === DBConnType.POSTGRES || dbType === DBConnType.MYSQL) ? "select * from <table name>;" : "db.<collection name>.find()"

    return (
        <React.Fragment>
            <ReactCodeMirror
                ref={editorRef}
                value={value}
                extensions={dbType === DBConnType.POSTGRES || dbType === DBConnType.MYSQL ? [sql()] : [javascript()]}
                theme={duotoneLight}
                height={"auto"}
                minHeight="80px"
                placeholder={placeholderText}
                basicSetup={{
                    autocompletion: false,
                    highlightActiveLine: false,
                }}
                onChange={onChange}
                onKeyDown={handleKeyDown}
            />
            <div className={styles.editorBottomBar}>
                <div className="columns">
                    <div className="column is-half">
                        <input
                            className="input"
                            type="name"
                            placeholder="Enter query name"
                            value={queryName}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setQueryName(e.target.value) }}
                        />
                    </div>
                    <div className={"column " + styles.buttons}>
                        {!saving && <button id="btnSaveQuery" data-tooltip-content="Save query" className="button " onClick={startSaving}>
                            <span className="icon is-small">
                                <i className="fas fa-save" aria-hidden="true"></i>
                            </span>
                        </button>}
                        {saving && <button className="button is-loading">Saving</button>}
                        &nbsp;&nbsp;
                        <button className="button" id="btnFormatQuery" data-tooltip-content="Format query" onClick={formatQuery}>
                            <span className="icon is-small">
                                <i className="fas fa-align-left" aria-hidden="true"></i>
                            </span>
                        </button>
                        &nbsp;&nbsp;
                        <button className="button" id="btnShowCheatsheet" data-tooltip-content="Show Cheatsheet" onClick={() => { setShowCheatsheet(true) }}>
                            <span className="icon is-small">
                                <i className="fas fa-book"></i>
                            </span>
                        </button>
                        &nbsp;&nbsp;
                        {!deleting && <button id="btnDelQuery" data-tooltip-content="Delete query" className="button is-danger" onClick={startDeleting}>
                            <span className="icon is-small">
                                <i className="fas fa-trash" aria-hidden="true"></i>
                            </span>
                        </button>}
                        {deleting && <button className="button is-danger is-loading">Deleting</button>}
                        &nbsp;&nbsp;
                        {!running && <button id="btnRunQuery" data-tooltip-content="Ctrl+Enter to run query" className="button is-primary" onClick={startRunningQuery}>
                            <span className="icon is-small">
                                <i className="fas fa-play-circle" aria-hidden="true"></i>
                            </span>&nbsp;&nbsp;
                            Run query
                        </button>}
                        {running && <button className="button is-primary is-loading">Running</button>}
                    </div>
                </div>
            </div>
            {showCheatsheet && <CheatSheetModal dbType={dbConnection!.type} onClose={() => { setShowCheatsheet(false) }} />}
            <Tooltip anchorId="btnRunQuery" />
            <Tooltip anchorId="btnSaveQuery" />
            <Tooltip anchorId="btnDelQuery" />
            <Tooltip anchorId="btnFormatQuery" />
            <Tooltip anchorId="btnShowCheatsheet" />
        </React.Fragment>
    )
}

export default QueryEditor