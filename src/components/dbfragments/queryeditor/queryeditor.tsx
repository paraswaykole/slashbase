import styles from './queryeditor.module.scss'
import React, { useState } from 'react'
import dynamic from 'next/dynamic'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { saveDBQuery, selectDBConnection } from '../../../redux/dbConnectionSlice'
import { DBConnection } from '../../../data/models'
import toast from 'react-hot-toast'


const CodeMirror = dynamic(() => {
    // @ts-ignore
    import('codemirror/mode/sql/sql')
    return import('react-codemirror')
}, {ssr: false})

type QueryEditorPropType = {
    initialValue: string,
    initQueryName: string,
    queryId: string,
    runQuery: (query: string, callback: ()=>void) => void
}


const QueryEditor = ({initialValue, initQueryName, queryId, runQuery}: QueryEditorPropType) => {

    const dispatch = useAppDispatch()

    const [value, setValue] = useState(initialValue)
    const [queryName, setQueryName] = useState(initQueryName)
    const [saving, setSaving] = useState(false)
    const [running, setRunning] = useState(false)

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    const startSaving = async () => {
        setSaving(true)
        try{
            await dispatch(saveDBQuery({dbConnId: dbConnection!.id, queryId, name: queryName, query: value})).unwrap()
            toast.success("Saved Succesfully!")
        } catch(e) {
            toast.error("There was some problem saving! Please try again.")
        }
        setSaving(false)
    }

    const startRunningQuery = () => {
        runQuery(value, ()=>{
            setRunning(false)
        })
    }

    return (
        <React.Fragment>
        <CodeMirror 
            value={value}
            options={{
                mode: 'sql',
                theme: 'duotone-light',
                lineNumbers: true
            }}
            onChange={(newValue) => {
                setValue(newValue)
            }}/>
            <div className={styles.editorBottomBar}>
                <div className="columns">
                    <div className="column is-two-thirds">
                            <input 
                                className="input" 
                                type="name" 
                                placeholder="Enter query name" 
                                value={queryName} 
                                onChange={(e: React.ChangeEvent<HTMLInputElement>)=>{setQueryName(e.target.value)}}
                            />
                    </div>
                    <div className={"column "+styles.buttons}>
                        { !running && <button className="button" onClick={startRunningQuery}>Run Query</button>}
                        { running && <button className="button is-loading">Running</button>}
                        &nbsp;&nbsp;
                        { !saving && <button className="button is-primary" onClick={startSaving}>Save Query</button>}
                        { saving && <button className="button is-primary is-loading">Saving</button>}
                    </div>
                </div>
            </div>
        </React.Fragment>
    )
}

export default QueryEditor