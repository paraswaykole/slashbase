import React from 'react'
import { DBConnection, Project, Tab } from '../../data/models'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import { selectCurrentProject } from '../../redux/projectsSlice'
import { selectActiveTab } from '../../redux/tabsSlice'
import DataModel from './datamodel/datamodel'

type DBShowModelPropType = {

}

const DBShowModelFragment = (_: DBShowModelPropType) => {

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const project: Project | undefined = useAppSelector(selectCurrentProject)
    const activeTab: Tab = useAppSelector(selectActiveTab)

    const mschema = activeTab.metadata.schema
    const mname = activeTab.metadata.name

    return (
        <React.Fragment>
            {mname && project &&
                <DataModel
                    dbConn={dbConnection!}
                    mschema={mschema!}
                    mname={mname}
                    isEditable={true} />
            }
        </React.Fragment>
    )
}


export default DBShowModelFragment