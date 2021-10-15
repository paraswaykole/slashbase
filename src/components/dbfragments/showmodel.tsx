import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, Project } from '../../data/models'
import apiService from '../../network/apiService'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import { selectProjects } from '../../redux/projectsSlice'
import DataModel from './datamodel/datamodel'

type DBShowModelPropType = { 

}

const DBShowModelFragment = (_: DBShowModelPropType) => {

    const router = useRouter()
    const { mschema, mname } = router.query

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const projects: Project[] = useAppSelector(selectProjects)
    const project: Project|undefined = projects.find(x=> x.id === dbConnection?.projectId)
    
    const [dataModel, setDataModel] = useState<DBDataModel>()
   
    useEffect(()=>{
        if (!dbConnection) return
        const fetchDataModel = async () => {
            const result = await apiService.getDBSingleDataModelByConnectionId(dbConnection!.id, String(mschema), String(mname))
            if (result.success) {
                setDataModel(result.data)
            }
        }
        fetchDataModel()
    }, [dbConnection, mschema, mname])    

    return (
        <React.Fragment>
            { dataModel && <DataModel dataModel={dataModel} /> }
        </React.Fragment>
    )
}


export default DBShowModelFragment