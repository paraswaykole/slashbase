import { useContext, useRef } from 'react'
import { DBConnection, Tab } from '../../../data/models'
import toast from 'react-hot-toast'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { addDBDataModelField } from '../../../redux/dataModelSlice'
import TabContext from '../../layouts/tabcontext'

type AddModal = {
    dbConn: DBConnection
    mSchema: string | null,
    mName: string,
    onAddField: () => void
    onClose: () => void
}

const AddFieldModal = ({ dbConn, mSchema, mName, onAddField, onClose }: AddModal) => {

    const dispatch = useAppDispatch()

    const activeTab: Tab = useContext(TabContext)!

    const fieldNameRef = useRef<HTMLInputElement>(null);
    const dataTypeRef = useRef<HTMLInputElement>(null);

    const startAdding = async () => {
        const result = await dispatch(addDBDataModelField({ tabId: activeTab.id, dbConnectionId: dbConn.id, schemaName: mSchema!, name: mName, fieldName: fieldNameRef.current!.value, dataType: dataTypeRef.current!.value })).unwrap()
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
