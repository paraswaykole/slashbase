import styles from './showdata.module.scss'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, DBQueryData, Project } from '../../data/models'
import apiService from '../../network/apiService'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import Table from './table/table'
import { selectProjects } from '../../redux/projectsSlice'
import { DBConnType, ProjectMemberRole } from '../../data/defaults'
import { selectIsShowingSidebar } from '../../redux/configSlice'
import JsonTable from './jsontable/jsontable'

type DBShowDataPropType = {

}

const DBShowDataFragment = (_: DBShowDataPropType) => {

    const router = useRouter()
    const { mschema, mname } = router.query

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)
    const projects: Project[] = useAppSelector(selectProjects)
    const project: Project | undefined = projects.find(x => x.id === dbConnection?.projectId)

    const [dataModel, setDataModel] = useState<DBDataModel>()
    const [queryData, setQueryData] = useState<DBQueryData>()
    const [queryOffset, setQueryOffset] = useState(0)
    const [queryCount, setQueryCount] = useState<number | undefined>(undefined)
    const [queryLimit] = useState(dbConnection ? dbConnection.type === DBConnType.POSTGRES ? 200 : 50 : 100)
    const [queryFilter, setQueryFilter] = useState<string[] | undefined>(undefined)
    const [querySort, setQuerySort] = useState<string[] | undefined>(undefined)
    const [dataLoading, setDataLoading] = useState(false)


    useEffect(() => {
        const dModel = dbDataModels.find(x => x.schemaName == mschema && x.name == mname)
        if (dModel) {
            setDataModel(dModel)
        }
        // else redirect to home fragment             
    }, [dbDataModels])

    useEffect(() => {
        if (dataModel && !queryCount) {
            fetchData(true)
        }
    }, [dataModel, queryCount])

    useEffect(() => {
        fetchData(false)
    }, [queryOffset, querySort])

    useEffect(() => {
        fetchData(true)
    }, [queryFilter])


    const fetchData = async (fetchCount: boolean) => {
        if (!dataModel || dataLoading) return
        setDataLoading(true)
        const result = await apiService.getDBDataInDataModel(dbConnection!.id, dataModel!.schemaName ?? '', dataModel!.name, queryLimit, queryOffset, fetchCount, queryFilter, querySort)
        if (result.success) {
            setQueryData(result.data)
            if (fetchCount) {
                setQueryCount(result.data.count)
            }
        }
        setDataLoading(false)
    }

    const onPreviousPage = () => {
        let previousOffset = queryOffset - queryLimit
        if (previousOffset < 0) {
            previousOffset = 0
        }
        setQueryOffset(previousOffset)
    }
    const onNextPage = () => {
        let nextOffset = queryOffset + queryLimit
        if (nextOffset > (queryCount ?? 0)) {
            return
        }
        setQueryOffset(nextOffset)
    }
    const onFilterChanged = (newFilter: string[] | undefined) => {
        setQueryFilter(newFilter)
        setQueryOffset(0)
    }

    const onSortChanged = (newSort: string[] | undefined) => {
        setQuerySort(newSort)
    }

    const updateCellData = (oldCtid: string, newCtid: string, columnIdx: string, newValue: string | null | boolean) => {
        const rowIdx = queryData!.rows.findIndex(x => x["0"] == oldCtid)
        if (rowIdx) {
            const newQueryData: DBQueryData = { ...queryData! }
            newQueryData!.rows[rowIdx] = { ...newQueryData!.rows[rowIdx], ctid: newCtid }
            newQueryData!.rows[rowIdx][columnIdx] = newValue
            setQueryData(newQueryData)
        } else {
            fetchData(false)
        }
    }

    const onDeleteRows = (indexes: number[]) => {
        const filteredRows = queryData!.rows.filter((_, i) => !indexes.includes(i))
        const newQueryData: DBQueryData = { ...queryData!, rows: filteredRows }
        setQueryData(newQueryData)
    }

    const onAddData = (newData: any) => {
        if (dbConnection!.type === DBConnType.POSTGRES) {
            const updatedRows = [newData, ...queryData!.rows]
            const updateQueryData: DBQueryData = { ...queryData!, rows: updatedRows }
            setQueryData(updateQueryData)
        } else if (dbConnection!.type === DBConnType.MONGO) {
            const updatedRows = [newData, ...queryData!.data]
            const updateQueryData: DBQueryData = { ...queryData!, data: updatedRows }
            setQueryData(updateQueryData)
        }
    }

    const rowsLength = queryData ? (queryData.rows ? queryData.rows.length : queryData.data.length) : 0
    const queryOffsetRangeEnd = (rowsLength ?? 0) === queryLimit ?
        queryOffset + queryLimit : queryOffset + (rowsLength ?? 0)

    return (
        <React.Fragment>
            {project && dbConnection && queryData && dbConnection.type === DBConnType.POSTGRES &&
                <Table
                    dbConnection={dbConnection}
                    mSchema={String(mschema)}
                    mName={String(mname)}
                    queryData={queryData}
                    querySort={querySort}
                    isEditable={project.currentMember?.role !== ProjectMemberRole.ANALYST}
                    showHeader={true}
                    updateCellData={updateCellData}
                    onDeleteRows={onDeleteRows}
                    onAddData={onAddData}
                    onFilterChanged={onFilterChanged}
                    onSortChanged={onSortChanged}
                />
            }
            {project && dbConnection && queryData && dbConnection.type === DBConnType.MONGO &&
                <JsonTable
                    dbConnection={dbConnection}
                    mName={String(mname)}
                    queryData={queryData}
                    isEditable={project.currentMember?.role !== ProjectMemberRole.ANALYST}
                    showHeader={true}
                    onAddData={onAddData}
                />
            }
            <br /><br /><br />
            <div className={styles.bottomBar + (isShowingSidebar ? ' withsidebar' : '')}>
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