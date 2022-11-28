import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, Project } from '../../data/models'
import apiService from '../../network/apiService'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import { ProjectPermissions, selectCurrentProject, selectProjectMemberPermissions } from '../../redux/projectsSlice'
import DataModel from './datamodel/datamodel'

type DBShowModelPropType = {

}

const DBShowModelFragment = (_: DBShowModelPropType) => {

    const router = useRouter()
    const { mschema, mname } = router.query

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const project: Project | undefined = useAppSelector(selectCurrentProject)
    const projectMemberPermissions: ProjectPermissions = useAppSelector(selectProjectMemberPermissions)

    const [dataModel, setDataModel] = useState<DBDataModel>()
    const [refresh, setRefresh] = useState<number>(Date.now())

    useEffect(() => {
        if (!dbConnection) return
        const fetchDataModel = async () => {
            const result = await apiService.getDBSingleDataModelByConnectionId(dbConnection!.id, String(mschema), String(mname))
            if (result.success) {
                setDataModel(result.data)
            }
        }
        fetchDataModel()
    }, [dbConnection, mschema, mname, refresh])

    return (
        <React.Fragment>
            {project && projectMemberPermissions && dataModel &&
                <DataModel
                    dbConn={dbConnection!}
                    dataModel={dataModel}
                    isEditable={!projectMemberPermissions.readOnly}
                    refreshModel={() => { setRefresh(Date.now()) }} />
            }
        </React.Fragment>
    )
}


export default DBShowModelFragment