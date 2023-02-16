import styles from './home.module.scss'
import React from 'react'
import { DBConnection, DBDataModel } from '../../data/models'
import { selectDBConnection, selectDBDataModels, selectIsFetchingDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import DBDataModelCard from '../cards/dbdatamodelcard/dbdatamodelcard'
import { TabType } from '../../data/defaults'
import { updateActiveTab } from '../../redux/tabsSlice'

type DBHomePropType = {
}

const DBHomeFragment = ({ }: DBHomePropType) => {

    const dispatch = useAppDispatch()

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)

    const isFetching: boolean = useAppSelector(selectIsFetchingDBDataModels)

    const updateActiveTabToHistory = () => {
        dispatch(updateActiveTab({ tabType: TabType.HISTORY }))
    }

    return (
        <React.Fragment>
            {dbConnection &&
                <React.Fragment>
                    <h1>Showing Data Models in {dbConnection.name}</h1>
                    {isFetching && <div className={styles.connectingMsg}>
                        <i className="fas fa-asterisk fa-spin"></i> Connecting to DB...
                    </div>
                    }
                    {dbDataModels.map(x => (
                        <DBDataModelCard key={x.schemaName + x.name} dataModel={x} dbConnection={dbConnection} />
                    ))}
                    <button className="button" onClick={updateActiveTabToHistory}>
                        <i className={"fas fa-history"} />
                        &nbsp;&nbsp;
                        View History
                    </button>
                </React.Fragment>
            }
        </React.Fragment>
    )
}


export default DBHomeFragment