import styles from './query.module.scss'
import React, { useContext, useEffect, useState } from 'react'
import toast from 'react-hot-toast'
import { DBConnection, DBQuery, DBQueryData, DBQueryResult, Tab } from '../../data/models'
import QueryEditor from './queryeditor/queryeditor'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { DBConnType } from '../../data/defaults'
import JsonTable from './jsontable/jsontable'
import Table from './table/table'
import Chart from './chart/chart'
import { getDBQuery, runQuery, selectDBQuery, setDBQuery } from '../../redux/dbQuerySlice'
import TabContext from '../layouts/tabcontext'


type DBQueryPropType = {
}

const DBQueryFragment = (_: DBQueryPropType) => {

    const dispatch = useAppDispatch()

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbQuery = useAppSelector(selectDBQuery)

    const currentTab: Tab = useContext(TabContext)!

    const [queryData, setQueryData] = useState<DBQueryData>()
    const [queryResult, setQueryResult] = useState<DBQueryResult>()
    const [isChartEnabled, setIsChartEnabled] = useState<boolean>(false)

    const queryId = currentTab.metadata.queryId

    useEffect(() => {
        (async () => {
            if (queryId && queryId !== 'new') {
                dispatch(getDBQuery({ queryId: String(queryId), tabId: currentTab.id }))
            }
            if (queryId === 'new') {
                dispatch(setDBQuery({ data: undefined, tabId: currentTab.id }))
            }
        })()
    }, [dispatch, queryId])


    useEffect(() => {
        setQueryData(undefined)
        setQueryResult(undefined)
        setIsChartEnabled(false)
    }, [queryId])

    const onRunQueryBtn = async (query: string, callback: () => void) => {
        const result = await dispatch(runQuery({ dbConnectionId: dbConnection!.id, query })).unwrap()
        if (result.success) {
            toast.success('Success')
            if ((result.data as DBQueryResult).message) {
                setQueryResult(result.data as DBQueryResult)
                setQueryData(undefined)
            } else {
                setQueryData(result.data as DBQueryData)
                setQueryResult(undefined)
            }
        } else {
            toast.error(result.error!)
        }
        callback()
    }

    const toggleIsChartEnabled = () => {
        setIsChartEnabled(!isChartEnabled)
    }

    const onQuerySaved = () => {
        //TODO: not implemented
    }

    const onDelete = () => {
        // TODO: not implemented
    }

    return (
        <div className={currentTab.isActive ? "db-tab-active" : "db-tab"}>
            {(dbConnection && ((queryId === 'new' && !dbQuery) || (dbQuery && dbQuery.id === queryId))) &&
                <QueryEditor
                    initialValue={dbQuery?.query ?? ''}
                    initQueryName={dbQuery?.name ?? ''}
                    queryId={queryId === 'new' ? '' : String(queryId)}
                    dbType={dbConnection!.type}
                    runQuery={onRunQueryBtn}
                    onSave={onQuerySaved}
                    onDelete={onDelete} />
            }
            <br />
            {queryData && <div className="tabs is-small is-centered is-toggle is-toggle-rounded tabs-set ">
                <ul>
                    <li className={!isChartEnabled ? 'is-active' : ''}>
                        <a onClick={toggleIsChartEnabled}>
                            <span className="icon is-small"><i className="fas fa-table" aria-hidden="true" /></span>
                            <span>Data</span>
                        </a>
                    </li>
                    <li className={isChartEnabled ? 'is-active' : ''}>
                        <a onClick={toggleIsChartEnabled}>
                            <span className="icon is-small"><i className="fas fa-chart-bar" aria-hidden="true" /></span>
                            <span>Chart</span>
                        </a>
                    </li>

                </ul>
            </div>}
            {queryData ? isChartEnabled ?
                <React.Fragment>
                    <Chart dbConn={dbConnection!} queryData={queryData} />
                </React.Fragment>
                :
                <React.Fragment>
                    {(dbConnection!.type === DBConnType.POSTGRES || dbConnection!.type === DBConnType.MYSQL) &&
                        <Table
                            dbConnection={dbConnection!}
                            queryData={queryData}
                            mSchema={''}
                            mName={''}
                            onFilterChanged={() => { }}
                            onSortChanged={() => { }}
                            isEditable={false} />
                    }
                    {dbConnection!.type === DBConnType.MONGO &&
                        <JsonTable
                            dbConnection={dbConnection!}
                            queryData={queryData}
                            mName={''}
                            onFilterChanged={() => { }}
                            onSortChanged={() => { }}
                            isEditable={false} />
                    }
                </React.Fragment>
                : null
            }
            {queryResult && <span><b>Result of Query: </b>{queryResult.message}</span>}
        </div>
    )
}

export default DBQueryFragment