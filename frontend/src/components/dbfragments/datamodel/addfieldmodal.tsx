import styles from './datamodel.module.scss'
import React, { useRef, useState } from 'react'
import { DBConnType } from '../../../data/defaults'
import { ApiResult, DBConnection, DBQueryResult } from '../../../data/models'
import apiService from '../../../network/apiService'
import toast from 'react-hot-toast'

type AddModal = {
    dbConn: DBConnection
    mSchema: string | null,
    mName: string,
    onAddField: () => void
    onClose: () => void
}

const AddFieldModal = ({ dbConn, mSchema, mName, onAddField, onClose }: AddModal) => {

    const fieldNameRef = useRef<HTMLInputElement>(null);
    const dataTypeRef = useRef<HTMLInputElement>(null);

    const startAdding = async () => {
        const result: ApiResult<DBQueryResult> = await apiService.addDBSingleDataModelField(dbConn.id, mSchema!, mName, fieldNameRef.current!.value, dataTypeRef.current!.value)
        if (result.success) {
            toast.success('new field added')
            onAddField()
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
                    <p className="modal-card-title">Add new field to {mSchema}.{mName}</p>
                    <button className="delete" aria-label="close" onClick={onClose}></button>
                </header>
                <section className="modal-card-body">
                    <div className="field">
                        <label className="label">New Field Name:</label>
                        <div className="control">
                            <input
                                ref={fieldNameRef}
                                className="input"
                                type="text"
                                placeholder="Enter name" />
                        </div>
                    </div>
                    <div className="field">
                        <label className="label">New Field Type:</label>
                        <div className="control">
                            <input
                                ref={dataTypeRef}
                                className="input"
                                type="text"
                                placeholder="Enter type" />
                        </div>
                    </div>
                </section>
                <footer className="modal-card-foot">
                    <button className="button is-primary" onClick={startAdding}>Add</button>
                    <button className="button" onClick={onClose}>Cancel</button>
                </footer>
            </div>
        </div>
    )
}

export default AddFieldModal
