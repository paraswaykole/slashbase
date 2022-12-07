import styles from './datamodel.module.scss'
import React, { useRef, useState } from 'react'
import { DBConnType } from '../../../data/defaults'
import { ApiResult, DBConnection, DBQueryResult } from '../../../data/models'
import apiService from '../../../network/apiService'
import toast from 'react-hot-toast'

type AddIndexModal = {
    dbConn: DBConnection
    mSchema: string | null,
    mName: string,
    onAddIndex: () => void
    onClose: () => void
}

const AddIndexModal = ({ dbConn, mSchema, mName, onAddIndex, onClose }: AddIndexModal) => {

    const indexNameRef = useRef<HTMLInputElement>(null);
    const fieldNamesRef = useRef<HTMLInputElement>(null);
    const isUnqiueRef = useRef<HTMLInputElement>(null);

    const startAdding = async () => {
        const result: ApiResult<DBQueryResult> = await apiService.addDBSingleDataModelIndex(dbConn.id, mSchema!, mName, indexNameRef.current!.value, fieldNamesRef.current!.value.split(","), isUnqiueRef.current!.checked)
        if (result.success) {
            toast.success('new index added')
            onAddIndex()
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
                    <p className="modal-card-title">Add new index to {mSchema}.{mName}</p>
                    <button className="delete" aria-label="close" onClick={onClose}></button>
                </header>
                <section className="modal-card-body">
                    <div className="field">
                        <label className="label">New Index Name:</label>
                        <div className="control">
                            <input
                                ref={indexNameRef}
                                className="input"
                                type="text"
                                placeholder="Enter name" />
                        </div>
                    </div>
                    <div className="field">
                        <label className="label">Field Names (comma seperated):</label>
                        <div className="control">
                            <input
                                ref={fieldNamesRef}
                                className="input"
                                type="text"
                                placeholder="Enter type" />
                        </div>
                    </div>
                    <div className="field">
                        <label className="label">Is Unique Index?</label>
                        <div className="control">
                            <input
                                ref={isUnqiueRef}
                                type="checkbox" />&nbsp;
                            Select
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

export default AddIndexModal
