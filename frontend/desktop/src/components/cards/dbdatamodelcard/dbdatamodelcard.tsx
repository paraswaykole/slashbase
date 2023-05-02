import styles from './dbdatamodelcard.module.scss'
import { DBConnection, DBDataModel } from '../../../data/models'
import { DBConnType, TabType } from '../../../data/defaults'
import { useAppDispatch } from '../../../redux/hooks'
import { updateActiveTab } from '../../../redux/tabsSlice'

type DBDataModelPropType = {
    dbConnection: DBConnection
    dataModel: DBDataModel
}

const DBDataModelCard = ({ dataModel, dbConnection }: DBDataModelPropType) => {

    const dispatch = useAppDispatch()

    const updateActiveTabToData = () => {
        dispatch(updateActiveTab({ tabType: TabType.DATA, metadata: { schema: dataModel.schemaName, name: dataModel.name } }))
    }

    const updateActiveTabToModel = () => {
        dispatch(updateActiveTab({ tabType: TabType.MODEL, metadata: { schema: dataModel.schemaName, name: dataModel.name } }))
    }

    return (
        <div className={"card " + styles.cardContainer}>
            <div className={"card-content " + styles.cardContent}>
                <div>
                    {dbConnection.type === DBConnType.POSTGRES &&
                        <b>{dataModel.schemaName}.{dataModel.name}</b>}
                    {dbConnection.type === DBConnType.MONGO &&
                        <b>{dataModel.name}</b>}
                    {dbConnection.type === DBConnType.MYSQL &&
                        <b>{dataModel.name}</b>}
                </div>
                <div className="buttons">
                    <button className="button is-small is-white" onClick={updateActiveTabToData}>
                        <span className="icon is-small"><i className="fas fa-table" /></span>
                        <span>View Data</span>
                    </button>
                    <button className="button is-small is-white" onClick={updateActiveTabToModel}>
                        <span className="icon is-small"><i className="fas fa-list-alt" /></span>
                        <span>View Model</span>
                    </button>
                </div>
            </div>
        </div>

    )
}


export default DBDataModelCard