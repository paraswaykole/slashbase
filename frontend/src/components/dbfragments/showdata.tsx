import styles from './showdata.module.scss'
import React, { useContext, useEffect, useState } from 'react'
import { DBConnection, DBDataModel, Project, Tab } from '../../data/models'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import Table from './table/table'
import { ProjectPermissions, selectCurrentProject, selectProjectMemberPermissions } from '../../redux/projectsSlice'
import { DBConnType } from '../../data/defaults'
import { selectIsShowingSidebar } from '../../redux/configSlice'
import JsonTable from './jsontable/jsontable'
import { getDBDataInDataModel, selectIsFetchingQueryData, selectQueryData } from '../../redux/dataModelSlice'
import TabContext from '../layouts/tabcontext'
import Button from '../ui/Button'

const DBShowDataFragment = () => {

    const dispatch = useAppDispatch()

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)
    const project: Project | undefined = useAppSelector(selectCurrentProject)
    const projectMemberPermissions: ProjectPermissions = useAppSelector(selectProjectMemberPermissions)
    const currentTab: Tab = useContext(TabContext)!

    const [dataModel, setDataModel] = useState<DBDataModel>()
    const dataLoading = useAppSelector(selectIsFetchingQueryData)
    const queryData = useAppSelector(selectQueryData)
    const [queryOffset, setQueryOffset] = useState(0)
    const [queryCount, setQueryCount] = useState<number | undefined>(undefined)
    const [queryLimit] = useState(dbConnection ? (dbConnection.type === DBConnType.POSTGRES || dbConnection.type === DBConnType.MYSQL) ? 200 : 50 : 100)
    const [queryFilter, setQueryFilter] = useState<string[] | undefined>(undefined)
    const [querySort, setQuerySort] = useState<string[] | undefined>(undefined)

    const mschema = currentTab.metadata.schema
    const mname = currentTab.metadata.name

    useEffect(() => {
        const dModel = dbDataModels.find(x => x.schemaName === mschema && x.name === mname)
        if (dModel) {
            setDataModel(dModel)
        }
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


    const fetchData = async (isFirstFetch: boolean) => {
        if (!dataModel) return
        try {
            const result = await dispatch(getDBDataInDataModel({ tabId: currentTab.id, dbConnectionId: dbConnection!.id, schemaName: dataModel!.schemaName ?? '', name: dataModel!.name, queryLimit, queryOffset, isFirstFetch, queryFilter, querySort })).unwrap()
            if (isFirstFetch) {
                setQueryCount(result.data.count)
            }
        } catch (e) {
            console.log(e)
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

    const onRefresh = () => {
        fetchData(false);
    }

    const rowsLength = queryData ? (queryData.rows ? queryData.rows.length : queryData.data.length) : 0
    const queryOffsetRangeEnd = (rowsLength ?? 0) === queryLimit ?
        queryOffset + queryLimit : queryOffset + (rowsLength ?? 0)

    return (
        <div className={currentTab.isActive ? "db-tab-active" : "db-tab"}>
            {project && dbConnection && queryData && dbConnection.type === DBConnType.POSTGRES &&
                <Table
                    dbConnection={dbConnection}
                    mSchema={String(mschema)}
                    mName={String(mname)}
                    queryData={queryData}
                    querySort={querySort}
                    isInteractive={true}
                    isReadOnly={projectMemberPermissions.readOnly}
                    showHeader={true}
                    onRefresh={onRefresh}
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
                    isInteractive={true}
                    isReadOnly={projectMemberPermissions.readOnly}
                    showHeader={true}
                    onRefresh={onRefresh}
                    onFilterChanged={onFilterChanged}
                    onSortChanged={onSortChanged}
                />
            }
            {project && dbConnection && queryData && dbConnection.type === DBConnType.MONGO &&
                <JsonTable
                    dbConnection={dbConnection}
                    mName={String(mname)}
                    queryData={queryData}
                    isInteractive={true}
                    isReadOnly={projectMemberPermissions.readOnly}
                    showHeader={true}
                    onRefresh={onRefresh}
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
                        <Button className="pagination-previous" text='Previous' onClick={onPreviousPage} />
                        <Button className="pagination-next" text='Next' onClick={onNextPage} />
                        <ul className="pagination-list">
                            Showing {queryOffset} - {queryOffsetRangeEnd} of {queryCount}
                        </ul>
                    </nav>
                }
            </div>
        </div>
    )
}


export default DBShowDataFragment