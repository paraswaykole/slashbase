import styles from './table.module.scss'
import React, { useState } from 'react'
import { ApiResult, CTIDResponse, DBConnection, DBQueryData } from '../../../data/models'
import apiService from '../../../network/apiService'
import toast from 'react-hot-toast'

type AddModal = { 
    queryData: DBQueryData
    dbConnection: DBConnection
    mSchema: string,
    mName: string,
    onAddData: (newData: any) => void
    onClose: () => void
}

const AddModal = ({queryData, dbConnection, mSchema, mName, onAddData, onClose}: AddModal) => {

    const [newData, setNewData] = useState<any>({})

    const onFieldChange = (e: React.ChangeEvent<HTMLInputElement>, col: string) => {
        let tmpData = {...newData}
        tmpData[col] = e.target.value
        setNewData(tmpData)
    }

    const startAdding = async () => {
        const result: ApiResult<CTIDResponse> = await apiService.addDBData(dbConnection.id, mSchema, mName, newData)
        if(result.success) {
            toast.success('data added')
            let mNewData = {...newData, ctid: result.data.ctid}
            queryData.columns.forEach((col, i)=>{
                const colIdx = i.toString()
                if (mNewData[col] === undefined) {
                    mNewData[colIdx] = null
                } else {
                    mNewData[colIdx] = mNewData[col]
                    delete mNewData[col]
                }
            })
            onAddData(mNewData)
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
                                onChange={(e)=>{onFieldChange(e, col)}}
                                placeholder="Enter input"/>
                        </div>
                    </div>
                  )
              })}
          </section>
          <footer className="modal-card-foot">
            <button className="button is-primary" onClick={startAdding}>Add</button>
            <button className="button" onClick={onClose}>Cancel</button>
          </footer>
        </div>
      </div>
    )
}

export default AddModal
  