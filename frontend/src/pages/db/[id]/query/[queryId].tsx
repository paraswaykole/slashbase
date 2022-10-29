import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../../../components/layouts/applayout'
import DefaultErrorPage from 'next/error'
import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../../../redux/hooks'
import { DBConnection, DBQuery, DBQueryData, DBQueryResult } from '../../../../data/models'
import QueryEditor from '../../../../components/dbfragments/queryeditor/queryeditor'
import apiService from '../../../../network/apiService'
import toast from 'react-hot-toast'
import Table from '../../../../components/dbfragments/table/table'
import Constants from '../../../../constants'
import { DBConnType } from '../../../../data/defaults'
import JsonTable from '../../../../components/dbfragments/jsontable/jsontable'


const DBQueryPage: NextPage = () => {

    const router = useRouter()
    const { id, queryId } = router.query

    const [dbQuery, setDBQuery] = useState<DBQuery>()
    const [queryData, setQueryData] = useState<DBQueryData>()
    const [queryResult, setQueryResult] = useState<DBQueryResult>()
    const [error404, setError404] = useState(false)

    const dispatch = useAppDispatch()
    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    useEffect(() => {
        (async () => {
            if (id) {
                try {
                    await dispatch((getDBConnection({ dbConnId: String(id) }))).unwrap()
                } catch (e) {
                    setError404(true)
                    return
                }
                dispatch((getDBDataModels({ dbConnId: String(id) })))
                dispatch((getDBQueries({ dbConnId: String(id) })))
            }
            if (queryId && queryId !== 'new') {
                const result = await apiService.getSingleDBQuery(String(queryId))
                if (result.success) {
                    setDBQuery(result.data)
                } else {
                    setError404(false)
                    return
                }
            }
            if (queryId === 'new') {
                setDBQuery(undefined)
            }
            setQueryData(undefined)
        })()
    }, [dispatch, router, queryId, id])

    const runQuery = async (query: string, callback: () => void) => {
        const result = await apiService.runQuery(dbConnection!.id, query)
        if (result.success) {
            toast.success('Success')
            if ((result.data as DBQueryResult).message) {
                setQueryResult(result.data as DBQueryResult)
            } else {
                setQueryData(result.data as DBQueryData)
            }
        } else {
            toast.error(result.error!)
        }
        callback()
    }

    const onQuerySaved = (newQueryId: string) => {
        if (newQueryId !== queryId)
            router.replace(Constants.APP_PATHS.DB_QUERY.path.replace('[id]', String(id)).replace('[queryId]', newQueryId))
    }

    if (error404) {
        return (<DefaultErrorPage statusCode={404} />)
    }

    return (
        <AppLayout title={(dbQuery ? dbQuery.name + " | " : " New Query | ") + (dbConnection ? dbConnection.name + " | Slashbase" : "Slashbase")} key={String(queryId)}>
            {((queryId === 'new' && !dbQuery) || (dbQuery && dbQuery.id === queryId)) &&
                <QueryEditor
                    initialValue={dbQuery?.query ?? ''}
                    initQueryName={dbQuery?.name ?? ''}
                    queryId={queryId === 'new' ? '' : String(queryId)}
                    dbType={dbConnection!.type}
                    runQuery={runQuery}
                    onSave={onQuerySaved} />
            }
            <br />
            {queryData && dbConnection!.type === DBConnType.POSTGRES &&
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
            {queryData && dbConnection!.type === DBConnType.MONGO &&
                <JsonTable
                    dbConnection={dbConnection!}
                    queryData={queryData}
                    mName={''}
                    updateCellData={() => { }}
                    onDeleteRows={() => { }}
                    onAddData={() => { }}
                    isEditable={false} />
            }
            {queryResult && <span><b>Result of Query: </b>{queryResult.message}</span>}
        </AppLayout>
    )
}

export default DBQueryPage
