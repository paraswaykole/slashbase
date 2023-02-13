import { FunctionComponent, useEffect, useState } from 'react'

import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import DBHomeFragment from '../../components/dbfragments/home'
import { DBConnection, Tab } from '../../data/models'
import { useParams } from 'react-router-dom'
import { getTabs, selectTabs } from '../../redux/tabsSlice'

const DBPage: FunctionComponent<{}> = () => {

    const { id } = useParams()
    const [error404, setError404] = useState(false)
    const dispatch = useAppDispatch()

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const tabs: Tab[] = useAppSelector(selectTabs)

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
        <>
            <p>Tabs: {tabs.map(t => t.type)}</p>
            <DBHomeFragment />
        </>
    )
}

export default DBPage
