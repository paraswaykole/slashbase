import styles from './sidebar.module.scss'
import React from 'react'
import { useRouter } from 'next/router'
import { DBConnection, DBDataModel } from '../../data/models'
import { useAppSelector } from '../../redux/hooks'
import { selectAllDBConnections } from '../../redux/allDBConnectionsSlice'
import Constants from '../../constants'
import Link from 'next/link'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'

enum SidebarViewType {
    GENERIC = "GENERIC", // default
    DATABASE = "DATABASE" // Used to show elements of single database
}


type SidebarPropType = { }

const Sidebar = (_: SidebarPropType) => {

    const router = useRouter()

    const { mschema, mname } = router.query

    let sidebarView: SidebarViewType = 
        (router.pathname === Constants.APP_PATHS.DB.href) ?
        SidebarViewType.DATABASE : SidebarViewType.GENERIC
    
    const allDBConnections: DBConnection[] = useAppSelector(selectAllDBConnections)
    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)

    return (
        <aside className={"menu "+styles.sidebar}> 
            <div className={styles.spacebox}>
                {sidebarView === SidebarViewType.GENERIC &&
                    <React.Fragment>
                        <p className="menu-label">
                            Databases
                        </p>
                        <ul className="menu-list">
                            {allDBConnections.map((dbConn: DBConnection) => {
                                return (
                                    <li key={dbConn.id}>
                                        <Link href={Constants.APP_PATHS.DB.href} as={Constants.APP_PATHS.DB.as + dbConn.id}>
                                            <a>{dbConn.name}</a>
                                        </Link>
                                    </li>
                                )
                            })}
                        </ul>
                    </React.Fragment>
                }
                {sidebarView === SidebarViewType.DATABASE &&
                    <React.Fragment>
                        <p className="menu-label">
                            Data Models
                        </p>
                        <ul className="menu-list">
                            {dbDataModels.map((dataModel: DBDataModel) => {
                                return (
                                    <li  key={dataModel.schemaName+dataModel.name}>
                                        <Link 
                                            href={{pathname: Constants.APP_PATHS.DB.href, query: {mschema: dataModel.schemaName, mname: dataModel.name}}} 
                                            as={Constants.APP_PATHS.DB.as + dbConnection!.id + "?mschema="+dataModel.schemaName + "&mname="+dataModel.name}>
                                            <a className={dataModel.schemaName == mschema && dataModel.name == mname ? 'is-active' : ''}>
                                                {dataModel.schemaName}.{dataModel.name}
                                            </a>
                                        </Link>
                                    </li>
                                )
                            })}
                        </ul>
                    </React.Fragment>
                }
            </div>
        </aside>
    )
}


export default Sidebar