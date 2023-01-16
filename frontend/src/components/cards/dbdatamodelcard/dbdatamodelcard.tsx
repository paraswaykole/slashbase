import styles from './dbdatamodelcard.module.scss'
import { DBConnection, DBDataModel } from '../../../data/models'
import Constants from '../../../constants'
import { DBConnType } from '../../../data/defaults'
import { Link } from 'react-router-dom'

type DBDataModelPropType = {
    dbConnection: DBConnection
    dataModel: DBDataModel
}

const DBDataModelCard = ({ dataModel, dbConnection }: DBDataModelPropType) => {

    return (
        <div className={"card " + styles.cardContainer}>
            <Link
                to={Constants.APP_PATHS.DB_PATH.path.replace('[id]', dbConnection.id).replace('[path]', String('data')) + "?mschema=" + dataModel.schemaName + "&mname=" + dataModel.name}
                className={styles.cardLink}
            >
                <div className="card-content">
                    <i className={"fas fa-table"} />&nbsp;&nbsp;
                    {dbConnection.type === DBConnType.POSTGRES &&
                        <b>{dataModel.schemaName}.{dataModel.name}</b>}
                    {dbConnection.type === DBConnType.MONGO &&
                        <b>{dataModel.name}</b>}
                    {dbConnection.type === DBConnType.MYSQL &&
                        <b>{dataModel.name}</b>}
                </div>
            </Link>
        </div>

    )
}


export default DBDataModelCard