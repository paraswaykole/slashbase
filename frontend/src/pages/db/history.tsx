import React, { FunctionComponent, useEffect, useState } from 'react'
import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { DBConnection } from '../../data/models'
import DBHistoryFragment from '../../components/dbfragments/history'
import { useParams } from 'react-router-dom'

const DBHistoryPage: FunctionComponent<{}> = () => {

    const { id } = useParams()
    const dispatch = useAppDispatch()
    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    const [error404, setError404] = useState(false)

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
        })()
    }, [dispatch, id])


    if (error404) {
        return (<h1>DB not found</h1>)
    }

    return (
        <DBHistoryFragment />
    )
}

export default DBHistoryPage
