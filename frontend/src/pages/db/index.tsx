import React, { FunctionComponent, useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import DBHomeFragment from '../../components/dbfragments/home'
import { DBConnection, Tab } from '../../data/models'
import { createTab, closeTab, getTabs, selectTabs } from '../../redux/tabsSlice'
import { TabType } from '../../data/defaults'

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

    const createNewTab = async () => {
        await dispatch(createTab({ dbConnId: String(id), tabType: TabType.BLANK }))
    }

    const handleCloseTab = async (tabId: string) => {
        await dispatch(closeTab({ dbConnId: String(id), tabId }))
    }

    return (
        <React.Fragment>
            <div className="tabs is-boxed">
                <ul>
                    {tabs.map(t => <li key={t.id} className={t.isActive ? "is-active" : ""}>
                        <a>
                            <span>
                                {t.type === TabType.BLANK && "New Tab"}
                            </span>
                            <span className="icon" onClick={() => { handleCloseTab(t.id) }}><i className="fas fa-times" aria-hidden="true"></i></span>
                        </a>
                    </li>)}
                    <li>
                        <a onClick={createNewTab}>
                            <span className="icon"><i className="fas fa-plus" aria-hidden="true"></i></span>
                        </a>
                    </li>
                </ul>
            </div>
            <DBHomeFragment />
        </React.Fragment>
    )
}

export default DBPage
