import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../../../components/layouts/applayout'
import DefaultErrorPage from 'next/error'
import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../../../redux/hooks'
import { DBConnection, DBQuery } from '../../../../data/models'
import apiService from '../../../../network/apiService'
import Constants from '../../../../constants'
import DBQueryFragment from '../../../../components/dbfragments/query'


const DBQueryPage: NextPage = () => {

    const router = useRouter()
    const { id, queryId } = router.query

    const [dbQuery, setDBQuery] = useState<DBQuery>()
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
        })()
    }, [dispatch, router, queryId, id])

    const onQuerySaved = (newQueryId: string) => {
        if (newQueryId !== queryId)
            router.replace(Constants.APP_PATHS.DB_QUERY.path.replace('[id]', String(id)).replace('[queryId]', newQueryId))
    }

    if (error404) {
        return (<DefaultErrorPage statusCode={404} />)
    }

    return (
        <AppLayout title={(dbQuery ? dbQuery.name + " | " : " New Query | ") + (dbConnection ? dbConnection.name + " | Slashbase" : "Slashbase")} key={String(queryId)}>
            <DBQueryFragment
                queryId={String(queryId)}
                dbQuery={dbQuery}
                onQuerySaved={onQuerySaved}
            />
        </AppLayout>
    )
}

export default DBQueryPage
