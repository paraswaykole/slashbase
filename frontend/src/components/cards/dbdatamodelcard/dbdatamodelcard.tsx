import styles from './dbdatamodelcard.module.scss'
import React from 'react'
import { DBConnection, DBDataModel } from '../../../data/models'
import Constants from '../../../constants'
import Link from 'next/link'

type DBDataModelPropType = { 
    dbConnection: DBConnection
    dataModel: DBDataModel
}

const DBDataModelCard = ({dataModel, dbConnection}: DBDataModelPropType) => {

    return (
        <div className={"card "+styles.cardContainer}>
            <Link 
                href={{pathname: Constants.APP_PATHS.DB_PATH.path, query: {mschema: dataModel.schemaName, mname: dataModel.name}}} 
                as={Constants.APP_PATHS.DB_PATH.path.replace('[id]', dbConnection.id).replace('[path]', String('data'))+"?mschema="+dataModel.schemaName+"&mname="+dataModel.name}
                >
                <a className={styles.cardLink}>
                    <div className="card-content">
                        <b>{dataModel.schemaName}.{dataModel.name}</b>
                    </div>
                </a>
            </Link>
        </div>

    )
}


export default DBDataModelCard