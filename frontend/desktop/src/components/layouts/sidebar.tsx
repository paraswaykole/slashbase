import styles from './sidebar.module.scss'
import React from 'react'
import { Link, useLocation, useParams, useSearchParams } from 'react-router-dom'
import Constants from '../../constants'
import { DBConnection, DBDataModel, DBQuery } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { selectDBConnection, selectDBDataModels, selectDBDQueries } from '../../redux/dbConnectionSlice'
import { DBConnType, TabType } from '../../data/defaults'
import HomeSidebar from './sidebars/homesidebar'
import { createTab } from '../../redux/tabsSlice'

enum SidebarViewType {
    HOME = "HOME", // home sidebar
    DATABASE = "DATABASE", // Used to show elements of single database
    SETTINGS = "SETTINGS" // Used to show elements of settings screen
}

const Sidebar = () => {

    const location = useLocation()

    const sidebarView: SidebarViewType =
        (location.pathname.startsWith("/db")) ?
            SidebarViewType.DATABASE : (location.pathname.startsWith("/settings")) ? SidebarViewType.SETTINGS : SidebarViewType.HOME

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const dbQueries: DBQuery[] = useAppSelector(selectDBDQueries)

    const dispatch = useAppDispatch()


    const openDataTab = (schema: string, name: string) => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.DATA, metadata: { schema, name } }))
    }

    const openQueryTab = (queryId: string) => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.QUERY, metadata: { queryId } }))
    }

    const openConsoleTab = () => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.CONSOLE, metadata: {} }))
    }

    return (
        <aside className={"menu " + styles.sidebar}>
            <div className={styles.spacebox}>
                {sidebarView === SidebarViewType.HOME &&
                    <HomeSidebar />
                }
                {sidebarView === SidebarViewType.DATABASE && dbConnection &&
                    <React.Fragment>
                        <p className="menu-label">
                            Data Models
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            {dbDataModels.map((dataModel: DBDataModel) => {
                                const label = dbConnection.type === DBConnType.POSTGRES ? `${dataModel.schemaName}.${dataModel.name}` : `${dataModel.name}`
                                return (
                                    <li key={dataModel.schemaName + dataModel.name}>
                                        <a onClick={() => openDataTab(dataModel.schemaName ?? "", dataModel.name)}>
                                            {label}
                                        </a>
                                    </li>
                                )
                            })}
                        </ul>
                        <p className="menu-label">
                            Queries
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            {dbQueries.map((dbQuery: DBQuery) => {
                                return (
                                    <li key={dbQuery.id}>
                                        <a onClick={() => openQueryTab(dbQuery.id)}>
                                            {dbQuery.name}
                                        </a>
                                    </li>
                                )
                            })}
                            <li>
                                <a onClick={() => openQueryTab("new")}>
                                    <span className="icon">
                                        <i className="fas fa-plus-circle"></i>
                                    </span>
                                    &nbsp;New Query
                                </a>
                            </li>
                        </ul>
                        <p className="menu-label">
                            Toolbox
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            <li>
                                <a onClick={() => openConsoleTab()}>
                                    <span className="icon">
                                        <i className="fas fa-terminal"></i>
                                    </span>
                                    &nbsp;Console
                                </a>
                            </li>
                        </ul>
                    </React.Fragment>
                }
                {sidebarView === SidebarViewType.SETTINGS &&
                    <React.Fragment>
                        <p className="menu-label">
                            Settings
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            <li>
                                <Link
                                    to={Constants.APP_PATHS.SETTINGS_GENERAL.path}
                                    className={location.pathname.startsWith(Constants.APP_PATHS.SETTINGS_GENERAL.path) ? 'is-active' : ''}>
                                    General
                                </Link>
                            </li>
                        </ul>
                        <p className="menu-label">
                            Info
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            <li>
                                <Link
                                    to={Constants.APP_PATHS.SETTINGS_ABOUT.path}
                                    className={location.pathname === Constants.APP_PATHS.SETTINGS_ABOUT.path ? 'is-active' : ''}>
                                    About
                                </Link>
                            </li>
                            <li>
                                <Link
                                    to={Constants.APP_PATHS.SETTINGS_SUPPORT.path}
                                    className={location.pathname === Constants.APP_PATHS.SETTINGS_SUPPORT.path ? 'is-active' : ''}>
                                    Support
                                </Link>
                            </li>
                        </ul>
                    </React.Fragment>
                }
            </div>
        </aside>
    )
}


export default Sidebar