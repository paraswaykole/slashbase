import React, { FunctionComponent, useEffect, useState } from 'react'
import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import Constants from '../../constants'
import DBQueryFragment from '../../components/dbfragments/query'
import { useNavigate, useParams } from 'react-router-dom'
import { getDBQuery, selectDBQuery, setDBQuery } from '../../redux/dbQuerySlice'


const DBQueryPage: FunctionComponent<{}> = () => {

    const { id, queryId } = useParams()
    const navigate = useNavigate()

    const dbQuery = useAppSelector(selectDBQuery)
    const [error404, setError404] = useState(false)

    const dispatch = useAppDispatch()

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
                const result = await dispatch(getDBQuery({ queryId: String(queryId) })).unwrap()
                if (!result.success) {
                    setError404(false)
                    return
                }
            }
            if (queryId === 'new') {
                dispatch(setDBQuery(undefined))
            }
        })()
    }, [dispatch, queryId, id])

    const onQuerySaved = (newQueryId: string) => {
        if (newQueryId !== queryId)
            navigate(Constants.APP_PATHS.DB_QUERY.path.replace('[id]', String(id)).replace('[queryId]', newQueryId))
    }

    const onDelete = () => {
        navigate(Constants.APP_PATHS.DB_QUERY.path.replace('[id]', String(id)).replace('[queryId]', 'new'))
    }

    if (error404) {
        return (<h1>404 page not found</h1>)
    }

    return (
        <React.Fragment>
            <DBQueryFragment
                queryId={String(queryId)}
                dbQuery={dbQuery}
                onQuerySaved={onQuerySaved}
                onDelete={onDelete}
            />
        </React.Fragment>
    )
}

export default DBQueryPage
