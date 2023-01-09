import { FunctionComponent, useEffect, useState } from 'react'

import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import DBHomeFragment from '../../components/dbfragments/home'
import { DBConnection } from '../../data/models'
import { useParams } from 'react-router-dom'

const DBPage: FunctionComponent<{}> = () => {

    const { id } = useParams()
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
        })()
    }, [dispatch, id])

    if (error404) {
        return (<h1>DB not found</h1>)
    }

    return (
        <DBHomeFragment />
    )
}

export default DBPage
