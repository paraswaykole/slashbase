import React from 'react'
import { useSearchParams } from 'react-router-dom'
import { DBConnection, Project } from '../../data/models'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import { selectCurrentProject } from '../../redux/projectsSlice'
import DataModel from './datamodel/datamodel'

type DBShowModelPropType = {

}

const DBShowModelFragment = (_: DBShowModelPropType) => {

    const [searchParams] = useSearchParams()
    const mschema = searchParams.get("mschema")
    const mname = searchParams.get("mname")

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const project: Project | undefined = useAppSelector(selectCurrentProject)

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