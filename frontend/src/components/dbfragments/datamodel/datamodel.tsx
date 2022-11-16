import styles from './datamodel.module.scss'
import React, { useState } from 'react'
import { DBConnection, DBDataModel } from '../../../data/models'
import ReactTooltip from 'react-tooltip'
import { DBConnType } from '../../../data/defaults'
import AddFieldModal from './addfieldmodal'
import ConfirmModal from '../../widgets/confirmModal'
import apiService from '../../../network/apiService'
import toast from 'react-hot-toast'

type DataModelPropType = {
    dbConn: DBConnection
    dataModel: DBDataModel,
    isEditable: boolean,
    refreshModel?: () => void
}

const DataModel = ({ dbConn, dataModel, isEditable, refreshModel }: DataModelPropType) => {

    const label = dbConn.type === DBConnType.POSTGRES ? `${dataModel.schemaName}.${dataModel.name}` : `${dataModel.name}`

    const [isEditing, setIsEditing] = useState<boolean>(false)
    const [showingAddModal, setShowingAddModal] = useState<boolean>(false)
    const [isDeletingField, setIsDeletingField] = useState<string>('')

    const deleteField = async () => {
        const result = await apiService.deleteDBSingleDataModelField(dbConn.id, dataModel.schemaName!, dataModel.name, isDeletingField)
        if (result.success) {
            toast.success(`deleted field ${isDeletingField}`)
            refreshModel?.()
            setIsDeletingField('')
        } else {
            toast.error(result.error!)
        }
    }

    return (
        <React.Fragment>
            <div>
                <table className={"table is-bordered is-striped is-narrow is-hoverable"}>
                    <thead>
                        <tr>
                            <th colSpan={dbConn.type === DBConnType.POSTGRES ? 4 : 5}>
                                {label}
                                {isEditable && <button className="button is-small" style={{ float: 'right' }} onClick={() => { setIsEditing(!isEditing) }}>
                                    {isEditing && <i className={"fas fa-check"} />}
                                    {!isEditing && <i className={"fas fa-pen"} />}
                                </button>}
                            </th>
                            {dbConn.type === DBConnType.POSTGRES && isEditing && <th>
                                <button className="button is-primary is-small" onClick={() => { setShowingAddModal(true) }}>
                                    <i className={"fas fa-plus"} />
                                </button>
                            </th>}
                        </tr>
                    </thead>
                    <tbody>
                        {
                            dataModel.fields?.map(field => (
                                <tr key={field.name}>
                                    <td>{
                                        field.isPrimary ?
                                            <i className="fas fa-key fa-rotate-315" data-tip="Primary key" /> :
                                            field.isNullable ?
                                                <i className="fas fa-dot-circle" data-tip="Nullable" /> :
                                                <i className="fas fa-circle" data-tip="Not Nullable" />
                                    }</td>
                                    <td>{field.name}</td>
                                    <td colSpan={dbConn.type === DBConnType.MONGO ? 2 : 1}>{field.type}</td>
                                    {dbConn.type === DBConnType.POSTGRES && <td>
                                        {field.tags.length > 0 && field.tags.map<React.ReactNode>(tag => (
                                            <span key={tag} className="tag is-info is-light">{tag}</span>
                                        )).reduce((prev, curr) => [prev, ' ', curr])}
                                    </td>}
                                    {isEditing && <td>
                                        <button className="button is-danger is-small" style={{ float: 'right' }} onClick={() => { setIsDeletingField(field.name) }}>
                                            <i className={"fas fa-trash"} />
                                        </button>
                                    </td>}
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
                {dataModel.indexes && dataModel.indexes.length > 0 &&
                    <table className={"table is-bordered is-striped is-narrow is-hoverable"}>
                        <thead>
                            <tr>
                                <th colSpan={2}>Indexes</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                dataModel.indexes?.map(field => (
                                    <tr key={field.name}>
                                        <td>{field.name}</td>
                                        <td>{field.indexDef}</td>
                                    </tr>
                                ))
                            }
                        </tbody>
                    </table>}
                {dbConn.type === DBConnType.POSTGRES && showingAddModal && <AddFieldModal
                    dbConn={dbConn}
                    mSchema={dataModel.schemaName}
                    mName={dataModel.name}
                    onAddField={() => { refreshModel?.() }}
                    onClose={() => { setShowingAddModal(false) }} />}
                {isDeletingField !== '' && <ConfirmModal
                    message={`Delete ${isDeletingField}?`}
                    onConfirm={deleteField}
                    onClose={() => { setIsDeletingField('') }} />}
                <ReactTooltip place="bottom" type="dark" effect="solid" />
            </div>
        </React.Fragment>
    )
}

export default DataModel