import styles from './showdata.module.scss'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, DBQueryData, Project } from '../../data/models'
import apiService from '../../network/apiService'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import Table from './table/table'
import { selectProjects } from '../../redux/projectsSlice'
import { ProjectMemberRole } from '../../data/defaults'
import { selectIsShowingSidebar } from '../../redux/configSlice'

type DBShowDataPropType = { 

}

const DBShowDataFragment = (_: DBShowDataPropType) => {

    const router = useRouter()
    const { mschema, mname } = router.query

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)
    const projects: Project[] = useAppSelector(selectProjects)
    const project: Project|undefined = projects.find(x=> x.id === dbConnection?.projectId)
    
    const [dataModel, setDataModel] = useState<DBDataModel>()
    const [queryData, setQueryData] = useState<DBQueryData>()
    const [queryOffset, setQueryOffset] = useState(0)
    const [queryCount, setQueryCount] = useState<number|undefined>(undefined)
    const [queryLimit] = useState(200)
    const [dataLoading, setDataLoading] = useState(false)
    
    useEffect(()=>{
        const dModel = dbDataModels.find(x => x.schemaName == mschema && x.name == mname)
        if(dModel){
            setDataModel(dModel)
        }
        // else redirect to home fragment             
    }, [dbDataModels])

    useEffect(()=>{
        if (dataModel && !queryCount){ 
            fetchData(true)
        }
    }, [dataModel, queryCount])
    
    useEffect(()=>{
        fetchData(false)
    }, [queryOffset])
    
    const fetchData = async (fetchCount: boolean) => {
        if(!dataModel || dataLoading) return
        setDataLoading(true)
        const result = await apiService.getDBDataInDataModel(dbConnection!.id, dataModel!.schemaName ?? '', dataModel!.name, queryOffset, fetchCount)
        if (result.success) {
            setQueryData(result.data)
            if (!queryCount){
                setQueryCount(result.data.count)
            }
        }
        setDataLoading(false)
    }

    const onPreviousPage = () => {
        let previousOffset = queryOffset - queryLimit
        if(previousOffset < 0) {
            previousOffset = 0
        }
        setQueryOffset(previousOffset)
    }
    const onNextPage = () => {
        let nextOffset = queryOffset + queryLimit
        if (nextOffset > (queryCount ?? 0)){
            return
        }
        setQueryOffset(nextOffset)
    }

    const updateCellData = (oldCtid: string, newCtid: string, columnName: string, newValue: string|null|boolean) => {
        const rowIdx = queryData!.rows.findIndex(x => x.ctid == oldCtid)
        if (rowIdx) {
            const newQueryData: DBQueryData = {...queryData!}
            newQueryData!.rows[rowIdx] = {...newQueryData!.rows[rowIdx], ctid: newCtid}
            newQueryData!.rows[rowIdx][columnName] = newValue
            setQueryData(newQueryData)
        } else {
            fetchData(false)
        }
    }

    const onDeleteRows = (indexes: number[]) => {
        const filteredRows = queryData!.rows.filter((_,i) => !indexes.includes(i))     
        const newQueryData: DBQueryData = {...queryData!, rows: filteredRows}
        setQueryData(newQueryData)
    }

    const onAddData = (newData: any) => {
        const updatedRows = [newData, ...queryData!.rows]   
        const updateQueryData: DBQueryData = {...queryData!, rows: updatedRows}
        setQueryData(updateQueryData)
    }

    const queryOffsetRangeEnd = (queryData?.rows.length ?? 0) === queryLimit ? 
        queryOffset + queryLimit : queryOffset + (queryData?.rows.length ?? 0)

    return (
        <React.Fragment>
            { project && dbConnection && queryData && 
                <Table 
                    dbConnection={dbConnection} 
                    isEditable={project.currentMember?.role !== ProjectMemberRole.ANALYST}
                    queryData={queryData} 
                    updateCellData={updateCellData}
                    onDeleteRows={onDeleteRows}
                    onAddData={onAddData}
                    heading={`Showing ${dataModel?.schemaName}.${dataModel?.name}`}
                    mSchema={String(mschema)}
                    mName={String(mname)}
                />
            }
            <br/><br/><br/>
            <div className={styles.bottomBar+(isShowingSidebar?' withsidebar':'')}>
                {dataLoading ? 
                    <progress className="progress is-primary" max="100">loading</progress>
                    :
                    <nav className="pagination is-centered is-rounded" role="navigation" aria-label="pagination">
                        <button className="button pagination-previous" onClick={onPreviousPage}>Previous</button>
                        <button className="button pagination-next" onClick={onNextPage}>Next</button>
                        <ul className="pagination-list">
                            Showing {queryOffset} - {queryOffsetRangeEnd} of {queryCount}
                        </ul>
                    </nav>
                }
            </div>
        </React.Fragment>
    )
}


export default DBShowDataFragment