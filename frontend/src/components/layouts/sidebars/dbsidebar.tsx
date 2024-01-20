import styles from '../sidebar.module.scss'
import React, { useState } from 'react'
import { DBConnType, TabType } from "../../../data/defaults"
import { DBConnection, DBDataModel, DBQuery } from "../../../data/models"
import { selectDBConnection, selectDBDQueries, selectDBDataModels } from "../../../redux/dbConnectionSlice"
import { useAppDispatch, useAppSelector } from "../../../redux/hooks"
import { createTab } from "../../../redux/tabsSlice"

enum DBSidebarTabType {
    DATABASE = "DATABASE",
    QUERIES = "QUERIES",
    TOOLBOX = "TOOLBOX"
}


type DatabaseSidebarPropType = {}

const DatabaseSidebar = (_: DatabaseSidebarPropType) => {

    const dispatch = useAppDispatch()

    const [currentSidebarTab, setCurrentSidebarTab] = useState<DBSidebarTabType>(DBSidebarTabType.DATABASE)

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const dbQueries: DBQuery[] = useAppSelector(selectDBDQueries)

    const openDataTab = (schema: string, name: string) => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.DATA, metadata: { schema, name } }))
    }

    const openQueryTab = (queryId: string) => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.QUERY, metadata: { queryId } }))
    }

    const openConsoleTab = () => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.CONSOLE, metadata: {} }))
    }

    const openGenerateSQLTab = () => {
        dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.GENSQL, metadata: {} }))
    }

    const switchSidebarTab = (tabType: DBSidebarTabType) => {
        setCurrentSidebarTab(tabType)
    }


    return (
        <React.Fragment>
            <div className="tabs is-centered">
                <ul>
                    <li className={currentSidebarTab === DBSidebarTabType.DATABASE ? "is-active" : ""} onClick={() => { switchSidebarTab(DBSidebarTabType.DATABASE) }}>
                        <a>
                            <span className="icon is-small"><i className="fas fa-database" aria-hidden="true"></i></span>
                        </a>
                    </li>
                    <li className={currentSidebarTab === DBSidebarTabType.QUERIES ? "is-active" : ""} onClick={() => { switchSidebarTab(DBSidebarTabType.QUERIES) }}>
                        <a>
                            <span className="icon is-small"><i className="fas fa-file" aria-hidden="true"></i></span>
                        </a>
                    </li>
                    <li className={currentSidebarTab === DBSidebarTabType.TOOLBOX ? "is-active" : ""} onClick={() => { switchSidebarTab(DBSidebarTabType.TOOLBOX) }}>
                        <a>
                            <span className="icon is-small"><i className="fas fa-toolbox" aria-hidden="true"></i></span>
                        </a>
                    </li>
                </ul>
            </div>
            {currentSidebarTab === DBSidebarTabType.DATABASE &&
                <React.Fragment>
                    <p className="menu-label">
                        Data Models
                    </p>
                    <ul className={"menu-list " + styles.menuList}>
                        {dbDataModels.map((dataModel: DBDataModel) => {
                            const label = dbConnection!.type === DBConnType.POSTGRES ? `${dataModel.schemaName}.${dataModel.name}` : `${dataModel.name}`
                            return (
                                <li key={dataModel.schemaName + dataModel.name}>
                                    <a onClick={() => openDataTab(dataModel.schemaName ?? "", dataModel.name)}>
                                        {label}
                                    </a>
                                </li>
                            )
                        })}
                    </ul>
                </React.Fragment>
            }
            {currentSidebarTab === DBSidebarTabType.QUERIES &&
                <React.Fragment>
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
                </React.Fragment>}
            {currentSidebarTab === DBSidebarTabType.TOOLBOX &&
                <React.Fragment>
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
                        <li>
                            <a onClick={() => openGenerateSQLTab()}>
                                <span className="icon">
                                    <i className="fas fa-magic"></i>
                                </span>
                                &nbsp;Generate SQL
                            </a>
                        </li>
                    </ul>
                </React.Fragment>
            }
        </React.Fragment>
    )

}

export default DatabaseSidebar