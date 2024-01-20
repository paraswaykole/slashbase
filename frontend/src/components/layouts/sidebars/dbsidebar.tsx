import styles from '../sidebar.module.scss'
import React from 'react'
import { DBConnType, TabType } from "../../../data/defaults"
import { DBConnection, DBDataModel, DBQuery } from "../../../data/models"
import { selectDBConnection, selectDBDQueries, selectDBDataModels } from "../../../redux/dbConnectionSlice"
import { useAppDispatch, useAppSelector } from "../../../redux/hooks"
import { createTab } from "../../../redux/tabsSlice"


type DatabaseSidebarPropType = {}

const DatabaseSidebar = (_: DatabaseSidebarPropType) => {

    const dispatch = useAppDispatch()

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


    return (
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
    )

}

export default DatabaseSidebar