import styles from './queryeditor.module.scss'
import React, { useState } from 'react'
import { DBDataModel } from '../../../data/models'
import ReactTooltip from 'react-tooltip'

type DataModelPropType = {
    dataModel: DBDataModel,
}

const DataModel = ({dataModel}: DataModelPropType) => {

    return (
        <React.Fragment>
            <div>
                <table className={"table is-bordered is-striped is-narrow is-hoverable"}>
                    <thead>
                        <tr>
                            <th colSpan={5}>{dataModel.schemaName}.{dataModel.name}</th>
                        </tr>
                    </thead>
                    <tbody>
                    {
                        dataModel.fields?.map(field => (
                            <tr key={field.name}>
                                <td>{
                                    field.isPrimary ? 
                                    <i className="fas fa-key fa-rotate-315" data-tip="Primary key"/> : 
                                    field.isNullable ? 
                                    <i className="fas fa-dot-circle" data-tip="Nullable"/> :
                                    <i className="fas fa-circle" data-tip="Not Nullable"/>
                                }</td>
                                <td>{field.name}</td>
                                <td>{field.type}</td>
                                <td>
                                    {field.charMaxLength ? <span className="tag is-info is-light">Max Length: {field.charMaxLength}</span> : null}
                                    {field.default ? <span className="tag is-info is-light">Default: {field.default}</span> : null}
                                </td>
                            </tr>
                        ))
                    }
                    </tbody>
                    <ReactTooltip place="bottom" type="dark" effect="solid"/>
                </table>
            </div>
        </React.Fragment>
    )
}

export default DataModel