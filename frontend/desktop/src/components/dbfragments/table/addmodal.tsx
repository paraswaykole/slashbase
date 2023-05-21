import styles from './table.module.scss'
import React, { useContext, useState } from 'react'
import { ApiResult, AddDataResponse, DBConnection, DBQueryData, Tab } from '../../../data/models'
import toast from 'react-hot-toast'
import { useAppDispatch } from '../../../redux/hooks'
import { addDBData, setQueryData } from '../../../redux/dataModelSlice'
import { DBConnType } from '../../../data/defaults'
import TabContext from '../../layouts/tabcontext'
import Button from '../../ui/Button'

type AddModal = {
    queryData: DBQueryData
    dbConnection: DBConnection
    mSchema: string,
    mName: string,
    onClose: () => void
}

const AddModal = ({ queryData, dbConnection, mSchema, mName, onClose }: AddModal) => {

    const dispatch = useAppDispatch()

    const activeTab: Tab = useContext(TabContext)!

    const [newData, setNewData] = useState<any>({})

    const onFieldChange = (e: React.ChangeEvent<HTMLInputElement>, col: string) => {
        let tmpData = { ...newData }
        tmpData[col] = e.target.value
        setNewData(tmpData)
    }

    const startAdding = async () => {
        const result: ApiResult<AddDataResponse> = await dispatch(addDBData({ dbConnectionId: dbConnection.id, schemaName: mSchema, name: mName, data: newData })).unwrap()
        if (result.success) {
            toast.success('data added')
            let mNewData: any
            if (dbConnection.type === DBConnType.POSTGRES && result.data.data) {
                mNewData = { ...result.data.data, 0: result.data.newId }
            } else {
                mNewData = { ...newData, ctid: result.data.newId }
                queryData.columns.forEach((col, i) => {
                    const colIdx = i.toString()
                    if (mNewData[col] === undefined) {
                        mNewData[colIdx] = null
                    } else {
                        mNewData[colIdx] = mNewData[col]
                        delete mNewData[col]
                    }
                })
            }
            const updatedRows = [mNewData, ...queryData!.rows]
            const updateQueryData: DBQueryData = { ...queryData!, rows: updatedRows }
            dispatch(setQueryData({ data: updateQueryData, tabId: activeTab.id }))
            onClose()
        } else {
            toast.error(result.error!)
        }
    }

    return (
        <div className="modal is-active">
            <div className="modal-background"></div>
            <div className="modal-card">
                <header className="modal-card-head">
                    <p className="modal-card-title">Add new {mSchema}.{mName}</p>
                    <button className="delete" aria-label="close" onClick={onClose}></button>
                </header>
                <section className="modal-card-body">
                    {queryData.columns.filter(col => col !== 'ctid').map(col => {
                        return (
                            <div className="field" key={col}>
                                <label className="label">{col}</label>
                                <div className="control">
                                    <input
                                        className="input"
                                        type="text"
                                        value={newData[col]}
                                        onChange={(e) => { onFieldChange(e, col) }}
                                        placeholder="Enter input" />
                                </div>
                            </div>
                        )
                    })}
                </section>
                <footer className="modal-card-foot">
                    <Button text='Add' className="is-primary" onClick={startAdding} />
                    <Button text='Cancel' onClick={onClose} />
                </footer>
            </div>
        </div>
    )
}

export default AddModal
