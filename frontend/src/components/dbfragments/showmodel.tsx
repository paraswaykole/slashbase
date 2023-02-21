import { useContext } from 'react'
import { DBConnection, Tab } from '../../data/models'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import TabContext from '../layouts/tabcontext'
import DataModel from './datamodel/datamodel'

type DBShowModelPropType = {

}

const DBShowModelFragment = (_: DBShowModelPropType) => {

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const currentTab: Tab = useContext(TabContext)!

    const mschema = currentTab.metadata.schema
    const mname = currentTab.metadata.name

    return (
        <div className={currentTab.isActive ? "db-tab-active" : "db-tab"}>
            {mname && dbConnection &&
                <DataModel
                    dbConn={dbConnection!}
                    mschema={mschema!}
                    mname={mname}
                    isEditable={true} />
            }
        </div>
    )
}


export default DBShowModelFragment