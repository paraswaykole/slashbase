import styles from './datamodel.module.scss'
import React, { useContext, useEffect, useState } from 'react'
import { DBConnection, Tab } from '../../../data/models'
import { Tooltip } from 'react-tooltip'
import { DBConnType } from '../../../data/defaults'
import AddFieldModal from './addfieldmodal'
import ConfirmModal from '../../widgets/confirmModal'
import toast from 'react-hot-toast'
import AddIndexModal from './addindexmodal'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { deleteDBDataModelField, deleteDBDataModelIndex, getSingleDataModel, selectSingleDataModel } from '../../../redux/dataModelSlice'
import TabContext from '../../layouts/tabcontext'

type DataModelPropType = {
    dbConn: DBConnection
    mschema: string,
    mname: string,
    isEditable: boolean,
}

const DataModel = ({ dbConn, mschema, mname, isEditable }: DataModelPropType) => {

    const dispatch = useAppDispatch()

    const currentTab: Tab = useContext(TabContext)!
    const dataModel = useAppSelector(selectSingleDataModel)

    const [isEditingModel, setIsEditingModel] = useState<boolean>(false)
    const [isEditingIndex, setIsEditingIndex] = useState<boolean>(false)
    const [showingAddFieldModal, setShowingAddFieldModal] = useState<boolean>(false)
    const [showingAddIndexModal, setShowingAddIndexModal] = useState<boolean>(false)
    const [deletingField, setDeletingField] = useState<string>('')
    const [deletingIndex, setDeletingIndex] = useState<string>('')
    const [refresh, setRefresh] = useState<number>(Date.now())

    useEffect(() => {
        if (!dbConn) return
        dispatch(getSingleDataModel({ tabId: currentTab.id, dbConnectionId: dbConn!.id, schemaName: String(mschema), name: String(mname) }))
    }, [dispatch, dbConn, mschema, mname, refresh])

    const refreshModel = () => {
        setRefresh(Date.now())
    }

    if (!dataModel) {
        return null
    }
    const label = dbConn.type === DBConnType.POSTGRES ? `${dataModel.schemaName}.${dataModel.name}` : `${dataModel.name}`

    const deleteField = async () => {
        const result = await dispatch(deleteDBDataModelField({ tabId: currentTab.id, dbConnectionId: dbConn.id, schemaName: dataModel.schemaName!, name: dataModel.name, fieldName: deletingField })).unwrap()
        if (result.success) {
            toast.success(`deleted field ${deletingField}`)
            refreshModel()
            setDeletingField('')
        } else {
            toast.error(result.error!)
        }
    }

    const deleteIndex = async () => {
        const result = await dispatch(deleteDBDataModelIndex({ dbConnectionId: dbConn.id, schemaName: dataModel.schemaName!, name: dataModel.name, indexName: deletingIndex })).unwrap()
        if (result.success) {
            toast.success(`deleted index ${deletingIndex}`)
            refreshModel()
            setDeletingIndex('')
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
                            <th colSpan={(dbConn.type === DBConnType.POSTGRES || dbConn.type === DBConnType.MYSQL) ? 4 : 5}>
                                {label}
                                {isEditable && <button className="button is-small" style={{ float: 'right' }} onClick={() => { setIsEditingModel(!isEditingModel) }}>
                                    {isEditingModel && <i className={"fas fa-check"} />}
                                    {!isEditingModel && <i className={"fas fa-pen"} />}
                                </button>}
                            </th>
                            {(dbConn.type === DBConnType.POSTGRES || dbConn.type === DBConnType.MYSQL) && isEditingModel && <th>
                                <button className="button is-primary is-small" onClick={() => { setShowingAddFieldModal(true) }}>
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
                                    {(dbConn.type === DBConnType.POSTGRES || dbConn.type === DBConnType.MYSQL) && <td>
                                        {field.tags.length > 0 && field.tags.map<React.ReactNode>(tag => (
                                            <span key={tag} className="tag is-info is-light">{tag}</span>
                                        )).reduce((prev, curr) => [prev, ' ', curr])}
                                    </td>}
                                    {isEditingModel && <td>
                                        <button className="button is-danger is-small" style={{ float: 'right' }} onClick={() => { setDeletingField(field.name) }}>
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
                                <th colSpan={2}>
                                    Indexes
                                    {isEditable && <button className="button is-small" style={{ float: 'right' }} onClick={() => { setIsEditingIndex(!isEditingIndex) }}>
                                        {isEditingIndex && <i className={"fas fa-check"} />}
                                        {!isEditingIndex && <i className={"fas fa-pen"} />}
                                    </button>}
                                </th>
                                {isEditingIndex && <th>
                                    <button className="button is-primary is-small" onClick={() => { setShowingAddIndexModal(true) }}>
                                        <i className={"fas fa-plus"} />
                                    </button>
                                </th>}
                            </tr>
                        </thead>
                        <tbody>
                            {
                                dataModel.indexes?.map(idx => (
                                    <tr key={idx.name}>
                                        <td>{idx.name}</td>
                                        <td>{idx.indexDef}</td>
                                        {isEditingIndex && <td>
                                            <button className="button is-danger is-small" style={{ float: 'right' }} onClick={() => { setDeletingIndex(idx.name) }}>
                                                <i className={"fas fa-trash"} />
                                            </button>
                                        </td>}
                                    </tr>
                                ))
                            }
                        </tbody>
                    </table>}
                {(dbConn.type === DBConnType.POSTGRES || dbConn.type === DBConnType.MYSQL) && showingAddFieldModal && <AddFieldModal
                    dbConn={dbConn}
                    mSchema={dataModel.schemaName}
                    mName={dataModel.name}
                    onAddField={refreshModel}
                    onClose={() => { setShowingAddFieldModal(false) }} />}
                {showingAddIndexModal && <AddIndexModal
                    dbConn={dbConn}
                    mSchema={dataModel.schemaName}
                    mName={dataModel.name}
                    onAddIndex={refreshModel}
                    onClose={() => { setShowingAddIndexModal(false) }} />}
                {deletingField !== '' && <ConfirmModal
                    message={`Delete field: ${deletingField}?`}
                    onConfirm={deleteField}
                    onClose={() => { setDeletingField('') }} />}
                {deletingIndex !== '' && <ConfirmModal
                    message={`Delete index: ${deletingIndex}?`}
                    onConfirm={deleteIndex}
                    onClose={() => { setDeletingIndex('') }} />}
                <Tooltip place="bottom" variant="dark" />
            </div>
        </React.Fragment>
    )
}

export default DataModel