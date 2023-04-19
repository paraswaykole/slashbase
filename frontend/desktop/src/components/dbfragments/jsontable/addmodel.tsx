import styles from './jsontable.module.scss'
import React, { useContext, useRef, useState } from 'react'
import { ApiResult, AddDataResponse, DBConnection, DBQueryData, Tab } from '../../../data/models'
import toast from 'react-hot-toast'
import CodeMirror, { ReactCodeMirrorRef } from '@uiw/react-codemirror'
import { javascript } from '@codemirror/lang-javascript'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { addDBData, selectQueryData, setQueryData } from '../../../redux/dataModelSlice'
import TabContext from '../../layouts/tabcontext'

type AddModal = {
    dbConnection: DBConnection
    mName: string,
    onClose: () => void
}

const AddModal = ({ dbConnection, mName, onClose }: AddModal) => {

    const dispatch = useAppDispatch()

    const activeTab: Tab = useContext(TabContext)!
    const queryData = useAppSelector(selectQueryData)

    const editorRef = useRef<ReactCodeMirrorRef | null>(null);
    const [newData, setNewData] = useState<any>(`{\n\t\n}`)

    const startAdding = async () => {
        let jsonData: any
        try {
            jsonData = JSON.parse(newData)
        } catch (e: any) {
            toast.error(e.message)
            return
        }
        const result: ApiResult<AddDataResponse> = await dispatch(addDBData({ tabId: activeTab.id, dbConnectionId: dbConnection.id, schemaName: "", name: mName, data: jsonData })).unwrap()
        if (result.success) {
            toast.success('data added')
            let mNewData = { _id: result.data.newId, ...jsonData }
            const updatedRows = [mNewData, ...queryData!.data]
            const updateQueryData: DBQueryData = { ...queryData!, data: updatedRows }
            dispatch(setQueryData({ data: updateQueryData, tabId: activeTab.id }))
            onClose()
        } else {
            toast.error(result.error!)
        }
    }

    const onChange = React.useCallback((value: any) => {
        setNewData(value)
    }, []);


    return (<div className="modal is-active">
        <div className="modal-background"></div>
        <div className="modal-card">
            <header className="modal-card-head">
                <p className="modal-card-title">Add new data to {mName}</p>
                <button className="delete" aria-label="close" onClick={onClose}></button>
            </header>
            <section className="modal-card-body">
                <CodeMirror
                    ref={editorRef}
                    value={newData}
                    extensions={[javascript()]}
                    onChange={onChange}
                />
            </section>
            <footer className="modal-card-foot">
                <button className="button is-primary" onClick={startAdding}>Add</button>
                <button className="button" onClick={onClose}>Cancel</button>
            </footer>
        </div>
    </div>)
}

export default AddModal