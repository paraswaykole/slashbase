import styles from './queryeditor.module.scss'
import React, { useState, useRef } from 'react'
import dynamic from 'next/dynamic'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { saveDBQuery, selectDBConnection } from '../../../redux/dbConnectionSlice'
import { DBConnection } from '../../../data/models'
import toast from 'react-hot-toast'
import { format } from 'sql-formatter'
import { DBConnType } from '../../../data/defaults'


const WrappedCodeMirror = dynamic(() => {
    // @ts-ignore
    import('codemirror/mode/sql/sql')
    // @ts-ignore
    import('codemirror/mode/javascript/javascript')
    return import('../../lib/wrappedcodemirror')
}, { ssr: false })

const ForwardRefCodeMirror = React.forwardRef<
    ReactCodeMirror.ReactCodeMirror,
    ReactCodeMirror.ReactCodeMirrorProps
>((props, ref) => {
    return <WrappedCodeMirror {...props} editorRef={ref} />;
});

ForwardRefCodeMirror.displayName = 'ForwardRefCodeMirror';

type QueryEditorPropType = {
    initialValue: string,
    initQueryName: string,
    queryId: string,
    dbType: DBConnType
    runQuery: (query: string, callback: () => void) => void
    onSave: (queryId: string) => void
}

const QueryEditor = ({ initialValue, initQueryName, queryId, dbType, runQuery, onSave }: QueryEditorPropType) => {

    const dispatch = useAppDispatch()

    const [value, setValue] = useState(initialValue)
    const [queryName, setQueryName] = useState(initQueryName)
    const [saving, setSaving] = useState(false)
    const [running, setRunning] = useState(false)
    const editorRef = useRef<ReactCodeMirror.ReactCodeMirror | null>(null);

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    const startSaving = async () => {
        if (queryName === '') {
            return
        }
        setSaving(true)
        try {
            const result = await dispatch(saveDBQuery({ dbConnId: dbConnection!.id, queryId, name: queryName, query: value })).unwrap()
            toast.success("Saved Succesfully!")
            onSave(result.dbQuery.id)
        } catch (e) {
            toast.error("There was some problem saving! Please try again.")
        }
        setSaving(false)
    }

    const startRunningQuery = () => {
        runQuery(value, () => {
            setRunning(false)
        })
    }

    const formatQuery = () => {
        let formattedQuery: string = value
        if (dbType == DBConnType.POSTGRES) {
            formattedQuery = format(value, {
                language: "postgresql", // Defaults to "sql" (see the above list of supported dialects)
                uppercase: true, // Defaults to false
                linesBetweenQueries: 2,
            })
        }
        setValue(formattedQuery)
        editorRef.current?.getCodeMirror().setValue(formattedQuery)
    }

    return (
        <React.Fragment>
            <ForwardRefCodeMirror
                ref={editorRef}
                value={value}
                options={{
                    mode: dbType == DBConnType.POSTGRES ? 'sql' : 'javascript',
                    theme: 'duotone-light',
                    lineNumbers: true
                }}
                onChange={(newValue) => {
                    setValue(newValue)
                }} />
            <div className={styles.editorBottomBar}>
                <div className="columns">
                    <div className="column is-two-thirds">
                        <input
                            className="input"
                            type="name"
                            placeholder="Enter query name"
                            value={queryName}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setQueryName(e.target.value) }}
                        />
                    </div>
                    <div className={"column " + styles.buttons}>
                        {!saving && <button className="button " onClick={startSaving}>
                            <span className="icon is-small">
                                <i className="fas fa-save" aria-hidden="true"></i>
                            </span>
                        </button>}
                        {saving && <button className="button is-loading">Saving</button>}
                        &nbsp;&nbsp;
                        <button className="button" onClick={formatQuery}>
                            <span className="icon is-small">
                                <i className="fas fa-align-left" aria-hidden="true"></i>
                            </span>
                        </button>
                        &nbsp;&nbsp;
                        {!running && <button className="button is-primary" onClick={startRunningQuery}>
                            <span className="icon is-small">
                                <i className="fas fa-play-circle" aria-hidden="true"></i>
                            </span>&nbsp;&nbsp;
                            Run query
                        </button>}
                        {running && <button className="button is-primary is-loading">Running</button>}
                    </div>
                </div>
            </div>
        </React.Fragment>
    )
}

export default QueryEditor