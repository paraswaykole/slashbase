import styles from './query.module.scss'
import React, { useEffect, useState } from 'react'
import toast from 'react-hot-toast'
import { DBConnection, DBQuery, DBQueryData, DBQueryResult } from '../../data/models'
import QueryEditor from './queryeditor/queryeditor'
import apiService from '../../network/apiService'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import { DBConnType } from '../../data/defaults'
import JsonTable from './jsontable/jsontable'
import Table from './table/table'
import Chart from './chart/chart'


type DBQueryPropType = {
    queryId: string
    dbQuery?: DBQuery
    onQuerySaved: (newQueryId: string) => void,
}

const DBQueryFragment = ({ queryId, dbQuery, onQuerySaved }: DBQueryPropType) => {

    const [queryData, setQueryData] = useState<DBQueryData>()
    const [queryResult, setQueryResult] = useState<DBQueryResult>()
    const [isChartEnabled, setIsChartEnabled] = useState<boolean>(false)

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    useEffect(() => {
        setQueryData(undefined)
        setQueryResult(undefined)
        setIsChartEnabled(false)
    }, [queryId])

    const runQuery = async (query: string, callback: () => void) => {
        const result = await apiService.runQuery(dbConnection!.id, query)
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

    return (
        <React.Fragment>
            {(dbConnection && ((queryId === 'new' && !dbQuery) || (dbQuery && dbQuery.id === queryId))) &&
                <QueryEditor
                    initialValue={dbQuery?.query ?? ''}
                    initQueryName={dbQuery?.name ?? ''}
                    queryId={queryId === 'new' ? '' : String(queryId)}
                    dbType={dbConnection!.type ?? ''}
                    runQuery={runQuery}
                    onSave={onQuerySaved} />
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
                    {dbConnection!.type === DBConnType.POSTGRES &&
                        <Table
                            dbConnection={dbConnection!}
                            queryData={queryData}
                            mSchema={''}
                            mName={''}
                            updateCellData={() => { }}
                            onDeleteRows={() => { }}
                            onAddData={() => { }}
                            onFilterChanged={() => { }}
                            onSortChanged={() => { }}
                            isEditable={false} />
                    }
                    {dbConnection!.type === DBConnType.MONGO &&
                        <JsonTable
                            dbConnection={dbConnection!}
                            queryData={queryData}
                            mName={''}
                            updateCellData={() => { }}
                            onDeleteRows={() => { }}
                            onAddData={() => { }}
                            onFilterChanged={() => { }}
                            onSortChanged={() => { }}
                            isEditable={false} />
                    }
                </React.Fragment>
                : null
            }
            {queryResult && <span><b>Result of Query: </b>{queryResult.message}</span>}
        </React.Fragment>
    )
}

export default DBQueryFragment