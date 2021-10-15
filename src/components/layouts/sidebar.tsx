import styles from './sidebar.module.scss'
import React from 'react'
import { useRouter } from 'next/router'
import { DBConnection, DBDataModel, DBQuery } from '../../data/models'
import { useAppSelector } from '../../redux/hooks'
import { selectAllDBConnections } from '../../redux/allDBConnectionsSlice'
import Constants from '../../constants'
import Link from 'next/link'
import { selectDBConnection, selectDBDataModels, selectDBDQueries } from '../../redux/dbConnectionSlice'

enum SidebarViewType {
    GENERIC = "GENERIC", // default
    DATABASE = "DATABASE" // Used to show elements of single database
}


type SidebarPropType = { }

const Sidebar = (_: SidebarPropType) => {

    const router = useRouter()

    const { mschema, mname, queryId } = router.query

    let sidebarView: SidebarViewType = 
        (router.pathname === Constants.APP_PATHS.DB.path 
            || router.pathname === Constants.APP_PATHS.DB_PATH.path 
            || router.pathname === Constants.APP_PATHS.DB_QUERY.path
            || router.pathname === Constants.APP_PATHS.DB_HISTORY.path) ?
        SidebarViewType.DATABASE : SidebarViewType.GENERIC
    
    const allDBConnections: DBConnection[] = useAppSelector(selectAllDBConnections)
    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const dbQueries: DBQuery[] = useAppSelector(selectDBDQueries)

    return (
        <aside className={"menu "+styles.sidebar}> 
            <div className={styles.spacebox}>
                {sidebarView === SidebarViewType.GENERIC &&
                    <React.Fragment>
                        <p className="menu-label">
                            All Databases
                        </p>
                        <ul className={"menu-list "+styles.menuList}>
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
                             <a title={dbConnection.name} className="nolink"><i className="fas fa-database"/> {dbConnection?.name}</a>
                        </Link>
                        <p className="menu-label">
                            Data Models
                        </p>
                        <ul className={"menu-list "+styles.menuList}>
                            {dbDataModels.map((dataModel: DBDataModel) => {
                                const label = `${dataModel.schemaName}.${dataModel.name}`
                                return (
                                    <li key={dataModel.schemaName+dataModel.name}>
                                        <Link 
                                            href={{pathname: Constants.APP_PATHS.DB_PATH.path, query: {mschema: dataModel.schemaName, mname: dataModel.name}}} 
                                            as={Constants.APP_PATHS.DB_PATH.path.replace('[id]', dbConnection!.id).replace('[path]', 'data') + "?mschema="+dataModel.schemaName + "&mname="+dataModel.name}>
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
                        <ul className={"menu-list "+styles.menuList}>
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
                                    <a className={ queryId === 'new' ? 'is-active' : ''}>
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
            </div>
        </aside>
    )
}


export default Sidebar