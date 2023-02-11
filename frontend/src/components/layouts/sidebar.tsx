import styles from './sidebar.module.scss'
import React from 'react'
import { Link, useLocation, useParams, useSearchParams } from 'react-router-dom'
import Constants from '../../constants'
import { DBConnection, DBDataModel, DBQuery } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { selectDBConnection, selectDBDataModels, selectDBDQueries } from '../../redux/dbConnectionSlice'
import { selectIsShowingSidebar, setIsShowingSidebar } from '../../redux/configSlice'
import { DBConnType } from '../../data/defaults'
import HomeSidebar from './sidebars/homesidebar'

enum SidebarViewType {
    HOME = "HOME", // home sidebar
    DATABASE = "DATABASE", // Used to show elements of single database
    SETTINGS = "SETTINGS" // Used to show elements of settings screen
}

type SidebarPropType = {}

const Sidebar = (_: SidebarPropType) => {

    const location = useLocation()
    const { queryId } = useParams()
    const [searchParams] = useSearchParams()
    const mschema = searchParams.get("mschema")
    const mname = searchParams.get("mname")

    let sidebarView: SidebarViewType =
        (location.pathname.startsWith("/db")) ?
            SidebarViewType.DATABASE : (location.pathname.startsWith("/settings")) ? SidebarViewType.SETTINGS : SidebarViewType.HOME

    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)
    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const dbQueries: DBQuery[] = useAppSelector(selectDBDQueries)

    const dispatch = useAppDispatch()

    const toggleSidebar = () => {
        dispatch(setIsShowingSidebar(!isShowingSidebar))
    }

    return (
        <aside className={"menu " + styles.sidebar}>
            <div className={styles.spacebox}>
                {sidebarView === SidebarViewType.HOME &&
                    <HomeSidebar />
                }
                {sidebarView === SidebarViewType.DATABASE && dbConnection &&
                    <React.Fragment>
                        <Link to={Constants.APP_PATHS.DB.path.replace('[id]', dbConnection?.id)} className="nolink">
                            <i className="fas fa-database" /> {dbConnection?.name}
                        </Link>
                        <p className="menu-label">
                            Data Models
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            {dbDataModels.map((dataModel: DBDataModel) => {
                                const label = dbConnection.type === DBConnType.POSTGRES ? `${dataModel.schemaName}.${dataModel.name}` : `${dataModel.name}`
                                return (
                                    <li key={dataModel.schemaName + dataModel.name}>
                                        <Link
                                            to={Constants.APP_PATHS.DB_PATH.path.replace('[id]', dbConnection!.id).replace('[path]', 'data') + "?mschema=" + dataModel.schemaName + "&mname=" + dataModel.name}
                                            className={dataModel.schemaName === mschema && dataModel.name === mname ? 'is-active' : ''}>
                                            {label}
                                        </Link>
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
                                        <Link
                                            to={Constants.APP_PATHS.DB_QUERY.path.replace('[id]', dbConnection!.id).replace('[queryId]', dbQuery.id)}
                                            className={queryId === dbQuery.id ? 'is-active' : ''}>
                                            {dbQuery.name}
                                        </Link>
                                    </li>
                                )
                            })}
                            <li>
                                <Link
                                    to={Constants.APP_PATHS.DB_QUERY.path.replace('[id]', dbConnection!.id).replace('[queryId]', 'new')}
                                    className={queryId === 'new' ? 'is-active' : ''}>
                                    <span className="icon">
                                        <i className="fas fa-plus-circle"></i>
                                    </span>
                                    &nbsp;New Query
                                </Link>
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
            <div>
                <button className={"button " + [styles.btn, styles.sidebarHideBtn].join(' ')} onClick={toggleSidebar}>
                    <i className={"fas fa-angle-double-left"} />
                    {/* <span className={styles.btnMsg}>&nbsp;&nbsp;hide sidebar</span> */}
                </button>
            </div>
        </aside>
    )
}


export default Sidebar