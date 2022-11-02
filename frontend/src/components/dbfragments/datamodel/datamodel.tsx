import styles from './datamodel.module.scss'
import React, { useState } from 'react'
import { DBDataModel } from '../../../data/models'
import ReactTooltip from 'react-tooltip'
import { DBConnType } from '../../../data/defaults'

type DataModelPropType = {
    dbType: DBConnType
    dataModel: DBDataModel,
}

const DataModel = ({ dbType, dataModel }: DataModelPropType) => {

    const label = dbType === DBConnType.POSTGRES ? `${dataModel.schemaName}.${dataModel.name}` : `${dataModel.name}`

    return (
        <React.Fragment>
            <div>
                <table className={"table is-bordered is-striped is-narrow is-hoverable"}>
                    <thead>
                        <tr>
                            <th colSpan={5}>{label}</th>
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
                                    <td>{field.type}</td>
                                    {dbType === DBConnType.POSTGRES && <td>
                                        {field.charMaxLength ? <span className="tag is-info is-light">Max Length: {field.charMaxLength}</span> : null}
                                        {field.default ? <span className="tag is-info is-light">Default: {field.default}</span> : null}
                                    </td>}
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
                {dataModel.constraints && dataModel.constraints.length > 0 &&
                    <table className={"table is-bordered is-striped is-narrow is-hoverable"}>
                        <thead>
                            <tr>
                                <th colSpan={2}>Constraints</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                dataModel.constraints?.map(field => (
                                    <tr key={field.name}>
                                        <td>{field.name}</td>
                                        <td>{field.constraintDef}</td>
                                    </tr>
                                ))
                            }
                        </tbody>
                    </table>}
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
                <ReactTooltip place="bottom" type="dark" effect="solid" />
            </div>
        </React.Fragment>
    )
}

export default DataModel