import styles from './sidebar.module.scss'
import React from 'react'
import { useRouter } from 'next/router'
import { DBConnection, DBDataModel, DBQuery, User } from '../../data/models'
import { useAppSelector } from '../../redux/hooks'
import { selectAllDBConnections } from '../../redux/allDBConnectionsSlice'
import Constants from '../../constants'
import Link from 'next/link'
import { selectDBConnection, selectDBDataModels, selectDBDQueries } from '../../redux/dbConnectionSlice'
import { selectIsShowingSidebar, setIsShowingSidebar } from '../../redux/configSlice'
import { useDispatch } from 'react-redux'
import { DBConnType } from '../../data/defaults'
import { selectCurrentUser } from '../../redux/currentUserSlice'

enum SidebarViewType {
    GENERIC = "GENERIC", // default
    DATABASE = "DATABASE", // Used to show elements of single database
    SETTINGS = "SETTINGS" // Used to show elements of settings screen
}

type SidebarPropType = {}

const Sidebar = (_: SidebarPropType) => {

    const router = useRouter()
    const currentUser: User = useAppSelector(selectCurrentUser)

    const { mschema, mname, queryId } = router.query

    let sidebarView: SidebarViewType =
        (router.pathname === Constants.APP_PATHS.DB.path
            || router.pathname === Constants.APP_PATHS.DB_PATH.path
            || router.pathname === Constants.APP_PATHS.DB_QUERY.path
            || router.pathname === Constants.APP_PATHS.DB_HISTORY.path) ?
            SidebarViewType.DATABASE : (router.pathname.startsWith("/settings")) ? SidebarViewType.SETTINGS : SidebarViewType.GENERIC

    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)
    const allDBConnections: DBConnection[] = useAppSelector(selectAllDBConnections)
    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const dbQueries: DBQuery[] = useAppSelector(selectDBDQueries)


    const dispatch = useDispatch()

    const toggleSidebar = () => {
        dispatch(setIsShowingSidebar(!isShowingSidebar))
    }

    return (
        <aside className={"menu " + styles.sidebar}>
            <div className={styles.spacebox}>
                {sidebarView === SidebarViewType.GENERIC &&
                    <React.Fragment>
                        <p className="menu-label">
                            All Databases
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            {allDBConnections.map((dbConn: DBConnection) => {
                                return (
                                    <li key={dbConn.id}>
                                        <Link href={Constants.APP_PATHS.DB.path} as={Constants.APP_PATHS.DB.path.replace('[id]', dbConn.id)}>
                                            <a title={dbConn.name}>{dbConn.name}</a>
                                        </Link>
                                    </li>
                                )
                            })}
                        </ul>
                    </React.Fragment>
                }
                {sidebarView === SidebarViewType.DATABASE && dbConnection &&
                    <React.Fragment>
                        <Link href={Constants.APP_PATHS.DB.path} as={Constants.APP_PATHS.DB.path.replace('[id]', dbConnection?.id)}>
                            <a title={dbConnection.name} className="nolink"><i className="fas fa-database" /> {dbConnection?.name}</a>
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
                                            href={{ pathname: Constants.APP_PATHS.DB_PATH.path, query: { mschema: dataModel.schemaName, mname: dataModel.name } }}
                                            as={Constants.APP_PATHS.DB_PATH.path.replace('[id]', dbConnection!.id).replace('[path]', 'data') + "?mschema=" + dataModel.schemaName + "&mname=" + dataModel.name}>
                                            <a className={dataModel.schemaName == mschema && dataModel.name == mname ? 'is-active' : ''} title={label}>
                                                {label}
                                            </a>
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
                                            href={Constants.APP_PATHS.DB_QUERY.path}
                                            as={Constants.APP_PATHS.DB_QUERY.path.replace('[id]', dbConnection!.id).replace('[queryId]', dbQuery.id)}>
                                            <a className={queryId == dbQuery.id ? 'is-active' : ''} title={dbQuery.name}>
                                                {dbQuery.name}
                                            </a>
                                        </Link>
                                    </li>
                                )
                            })}
                            <li>
                                <Link
                                    href={Constants.APP_PATHS.DB_QUERY.path}
                                    as={Constants.APP_PATHS.DB_QUERY.path.replace('[id]', dbConnection!.id).replace('[queryId]', 'new')}>
                                    <a className={queryId === 'new' ? 'is-active' : ''}>
                                        <span className="icon">
                                            <i className="fas fa-plus-circle"></i>
                                        </span>
                                        &nbsp;New Query
                                    </a>
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
                                <Link href={Constants.APP_PATHS.SETTINGS_ACCOUNT.path} as={Constants.APP_PATHS.SETTINGS_ACCOUNT.path}>
                                    <a className={router.pathname.startsWith(Constants.APP_PATHS.SETTINGS_ACCOUNT.path) ? 'is-active' : ''}>Account</a>
                                </Link>
                            </li>
                        </ul>
                        {currentUser && currentUser.isRoot && <React.Fragment>
                            <p className="menu-label">
                                Admin (for root-user)
                            </p>
                            <ul className={"menu-list " + styles.menuList}>
                                <li>
                                    <Link href={Constants.APP_PATHS.SETTINGS_USERS.path} as={Constants.APP_PATHS.SETTINGS_USERS.path}>
                                        <a className={router.pathname.startsWith(Constants.APP_PATHS.SETTINGS_USERS.path) ? 'is-active' : ''}>Manage Users</a>
                                    </Link>
                                </li>
                                <li>
                                    <Link href={Constants.APP_PATHS.SETTINGS_ROLES.path} as={Constants.APP_PATHS.SETTINGS_ROLES.path}>
                                        <a className={router.pathname == Constants.APP_PATHS.SETTINGS_ROLES.path ? 'is-active' : ''}>Manage Roles</a>
                                    </Link>
                                </li>
                            </ul>
                        </React.Fragment>}
                        <p className="menu-label">
                            Info
                        </p>
                        <ul className={"menu-list " + styles.menuList}>
                            <li>
                                <Link href={Constants.APP_PATHS.SETTINGS_ABOUT.path} as={Constants.APP_PATHS.SETTINGS_ABOUT.path}>
                                    <a className={router.pathname === Constants.APP_PATHS.SETTINGS_ABOUT.path ? 'is-active' : ''}>About</a>
                                </Link>
                            </li>
                            <li>
                                <Link href={Constants.APP_PATHS.SETTINGS_SUPPORT.path} as={Constants.APP_PATHS.SETTINGS_SUPPORT.path}>
                                    <a className={router.pathname === Constants.APP_PATHS.SETTINGS_SUPPORT.path ? 'is-active' : ''}>Support</a>
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