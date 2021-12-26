import styles from './queryeditor.module.scss'
import React, { useState, useRef } from 'react'
import dynamic from 'next/dynamic'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { saveDBQuery, selectDBConnection } from '../../../redux/dbConnectionSlice'
import { DBConnection } from '../../../data/models'
import toast from 'react-hot-toast'
import {format} from 'sql-formatter'


const WrappedCodeMirror = dynamic(() => {
    // @ts-ignore
    import('codemirror/mode/sql/sql')
    return import('../wrappedcodemirror')
}, {ssr: false})

const ForwardRefCodeMirror = React.forwardRef<
  ReactCodeMirror.ReactCodeMirror,
  ReactCodeMirror.ReactCodeMirrorProps
>((props, ref) => {
  return <WrappedCodeMirror {...props} editorRef={ref} />;
});

type QueryEditorPropType = {
    initialValue: string,
    initQueryName: string,
    queryId: string,
    runQuery: (query: string, callback: ()=>void) => void
    onSave: (queryId: string) => void
}


const QueryEditor = ({initialValue, initQueryName, queryId, runQuery, onSave}: QueryEditorPropType) => {

    const dispatch = useAppDispatch()

    const [value, setValue] = useState(initialValue)
    const [queryName, setQueryName] = useState(initQueryName)
    const [saving, setSaving] = useState(false)
    const [running, setRunning] = useState(false)
    const editorRef = useRef<ReactCodeMirror.ReactCodeMirror | null>(null);

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    const startSaving = async () => {
        setSaving(true)
        try{
            let formattedQuery = format(value, {
              language: "postgresql", // Defaults to "sql" (see the above list of supported dialects)
              uppercase: true, // Defaults to false
              linesBetweenQueries: 2,
            });
            setValue(formattedQuery);
            editorRef.current?.getCodeMirror().setValue(formattedQuery);
            const result = await dispatch(saveDBQuery({dbConnId: dbConnection!.id, queryId, name: queryName, query: formattedQuery})).unwrap()
            toast.success("Saved Succesfully!")
            onSave(result.dbQuery.id)
        } catch(e) {
            toast.error("There was some problem saving! Please try again.")
        }
        setSaving(false)
    }

    const startRunningQuery = () => {
      let formattedQuery = format(value, {
        language: "postgresql", // Defaults to "sql" (see the above list of supported dialects)
        uppercase: true, // Defaults to false
        linesBetweenQueries: 2,
      });
      setValue(formattedQuery);
      editorRef.current?.getCodeMirror().setValue(formattedQuery);
        runQuery(value, ()=>{
            setRunning(false)
        })
    }

    return (
        <React.Fragment>
        <ForwardRefCodeMirror 
            ref={editorRef}
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