import { useContext, useRef } from 'react'
import { DBConnection, Tab } from '../../../data/models'
import toast from 'react-hot-toast'
import { addDBDataModelIndex } from '../../../redux/dataModelSlice'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import TabContext from '../../layouts/tabcontext'

type AddIndexModal = {
    dbConn: DBConnection
    mSchema: string | null,
    mName: string,
    onAddIndex: () => void
    onClose: () => void
}

const AddIndexModal = ({ dbConn, mSchema, mName, onAddIndex, onClose }: AddIndexModal) => {

    const dispatch = useAppDispatch()

    const activeTab: Tab = useContext(TabContext)!

    const indexNameRef = useRef<HTMLInputElement>(null);
    const fieldNamesRef = useRef<HTMLInputElement>(null);
    const isUnqiueRef = useRef<HTMLInputElement>(null);

    const startAdding = async () => {
        const result = await dispatch(addDBDataModelIndex({ tabId: activeTab.id, dbConnectionId: dbConn.id, schemaName: mSchema!, name: mName, indexName: indexNameRef.current!.value, fieldNames: fieldNamesRef.current!.value.split(","), isUnique: isUnqiueRef.current!.checked })).unwrap()
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
