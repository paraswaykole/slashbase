import styles from './home.module.scss'
import React, { useContext } from 'react'
import { DBConnection, DBDataModel, Tab } from '../../data/models'
import { selectDBConnection, selectDBDataModels, selectIsFetchingDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import DBDataModelCard from '../cards/dbdatamodelcard/dbdatamodelcard'
import { TabType } from '../../data/defaults'
import { updateActiveTab } from '../../redux/tabsSlice'
import TabContext from '../layouts/tabcontext'
import Button from '../ui/Button'

type DBHomePropType = {
}

const DBHomeFragment = ({ }: DBHomePropType) => {

    const dispatch = useAppDispatch()

    const currentTab: Tab = useContext(TabContext)!

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)

    const isFetching: boolean = useAppSelector(selectIsFetchingDBDataModels)

    const updateActiveTabToQuery = () => {
        dispatch(updateActiveTab({ tabType: TabType.QUERY, metadata: { queryId: 'new', query: "" } }))
    }

    const updateActiveTabToHistory = () => {
        dispatch(updateActiveTab({ tabType: TabType.HISTORY, metadata: {} }))
    }

    return (
        <div className={currentTab.isActive ? "db-tab-active" : "db-tab"}>
            {dbConnection &&
                <React.Fragment>
                    <h2>Data Models</h2>
                    {isFetching && <div className={styles.connectingMsg}>
                        <i className="fas fa-asterisk fa-spin"></i> Connecting to DB...
                    </div>
                    }
                    {dbDataModels.map(x => (
                        <DBDataModelCard key={x.schemaName + x.name} dataModel={x} dbConnection={dbConnection} />
                    ))}
                    <div className="buttons">
                        <Button
                            text='New Query'
                            icon={<i className="fas fa-circle-plus"/>} 
                            onClick={updateActiveTabToQuery}
                        />
                        <Button
                            text='View History'
                            icon={<i className="fas fa-history"/>} 
                            onClick={updateActiveTabToHistory}
                        />
                    </div>
                </React.Fragment>
            }
        </div>
    )
}


export default DBHomeFragment