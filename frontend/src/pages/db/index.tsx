import React, { FunctionComponent, useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getDBConnection, getDBDataModels, getDBQueries } from '../../redux/dbConnectionSlice'
import { useAppDispatch } from '../../redux/hooks'
import DBHomeFragment from '../../components/dbfragments/home'
import { getTabs } from '../../redux/tabsSlice'

const DBPage: FunctionComponent<{}> = () => {

    const { id } = useParams()
    const [error404, setError404] = useState(false)
    const dispatch = useAppDispatch()

    useEffect(() => {
        (async () => {
            if (id) {
                try {
                    await dispatch(getDBConnection({ dbConnId: String(id) })).unwrap()
                } catch (e) {
                    setError404(true)
                    return
                }
                dispatch(getTabs({ dbConnId: String(id) }))
                dispatch(getDBDataModels({ dbConnId: String(id) }))
                dispatch(getDBQueries({ dbConnId: String(id) }))
            }
        })()
    }, [dispatch, id])

    if (error404) {
        return (<h1>DB not found</h1>)
    }

    return (
        <React.Fragment>
            <DBHomeFragment />
        </React.Fragment>
    )
}

export default DBPage
