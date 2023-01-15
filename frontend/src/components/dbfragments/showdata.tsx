import styles from './showdata.module.scss'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, DBQueryData, Project } from '../../data/models'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import Table from './table/table'
import { ProjectPermissions, selectCurrentProject } from '../../redux/projectsSlice'
import { DBConnType } from '../../data/defaults'
import { selectIsShowingSidebar } from '../../redux/configSlice'
import JsonTable from './jsontable/jsontable'
import { useSearchParams } from 'react-router-dom'
import { getDBDataInDataModel, selectIsFetchingQueryData, selectQueryData, setQueryData } from '../../redux/dataModelSlice'

type DBShowDataPropType = {

}

const DBShowDataFragment = (_: DBShowDataPropType) => {

    const [searchParams] = useSearchParams()
    const mschema = searchParams.get("mschema")
    const mname = searchParams.get("mname")

    const dispatch = useAppDispatch()

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)
    const project: Project | undefined = useAppSelector(selectCurrentProject)

    const [dataModel, setDataModel] = useState<DBDataModel>()
    const dataLoading = useAppSelector(selectIsFetchingQueryData)
    const queryData = useAppSelector(selectQueryData)
    const [queryOffset, setQueryOffset] = useState(0)
    const [queryCount, setQueryCount] = useState<number | undefined>(undefined)
    const [queryLimit] = useState(dbConnection ? dbConnection.type === DBConnType.POSTGRES ? 200 : 50 : 100)
    const [queryFilter, setQueryFilter] = useState<string[] | undefined>(undefined)
    const [querySort, setQuerySort] = useState<string[] | undefined>(undefined)


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
        if (!dataModel) return
        const result = await dispatch(getDBDataInDataModel({ dbConnectionId: dbConnection!.id, schemaName: dataModel!.schemaName ?? '', name: dataModel!.name, queryLimit, queryOffset, fetchCount, queryFilter, querySort })).unwrap()
        if (fetchCount) {
            setQueryCount(result.data.count)
        }
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
                    isEditable={true}
                    showHeader={true}
                    onFilterChanged={onFilterChanged}
                    onSortChanged={onSortChanged}
                />
            }
            {project && dbConnection && queryData && dbConnection.type === DBConnType.MYSQL &&
                <Table
                    dbConnection={dbConnection}
                    mSchema={String(mschema)}
                    mName={String(mname)}
                    queryData={queryData}
                    querySort={querySort}
                    isEditable={true}
                    showHeader={true}
                    onFilterChanged={onFilterChanged}
                    onSortChanged={onSortChanged}
                />
            }
            {project && dbConnection && queryData && dbConnection.type === DBConnType.MONGO &&
                <JsonTable
                    dbConnection={dbConnection}
                    mName={String(mname)}
                    queryData={queryData}
                    isEditable={true}
                    showHeader={true}
                    onFilterChanged={onFilterChanged}
                    onSortChanged={onSortChanged}
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