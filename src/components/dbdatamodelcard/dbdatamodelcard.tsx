import styles from './dbdatamodelcard.module.scss'
import React from 'react'
import { DBConnection, DBDataModel } from '../../data/models'
import Constants from '../../constants'
import Link from 'next/link'

type DBDataModelPropType = { 
    dbConnection: DBConnection
    dataModel: DBDataModel
}

const DBDataModelCard = ({dataModel, dbConnection}: DBDataModelPropType) => {

    return (
        <Link 
            href={{pathname: Constants.APP_PATHS.DB.path, query: {mschema: dataModel.schemaName, mname: dataModel.name}}} 
            as={Constants.APP_PATHS.DB.path.replace('[id]', dbConnection.id)+"?mschema="+dataModel.schemaName+"&mname="+dataModel.name}
            >
            <a>
                <div className={"card "+styles.cardContainer}>
                    <div className="card-content">
                        <h2>{dataModel.schemaName}.{dataModel.name}</h2>
                    </div>
                </div>
            </a>
        </Link>
    )
}


export default DBDataModelCard