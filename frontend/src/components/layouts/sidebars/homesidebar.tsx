import styles from '../sidebar.module.scss'
import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import { DBConnection, Project } from '../../../data/models'
import { useAppSelector } from '../../../redux/hooks'
import { selectAllDBConnections } from '../../../redux/allDBConnectionsSlice'
import Constants from '../../../constants'
import { selectProjects } from '../../../redux/projectsSlice'
import CreateNewProjectModal from '../../home/createprojectmodal'

type HomeSidebarPropType = {}

const HomeSidebar = (_: HomeSidebarPropType) => {

    const [isShowingCreateProject, setIsShowingCreateProject] = useState(false)

    const allProjects: Project[] = useAppSelector(selectProjects)
    const allDBConnections: DBConnection[] = useAppSelector(selectAllDBConnections)

    return (
        <React.Fragment>
            <p className="menu-label">
                Projects & Databases
            </p>
            <ul className={"menu-list " + styles.menuList}>
                {allProjects.map((project: Project) => {
                    return (
                        <li key={project.id}>
                            <Link to={Constants.APP_PATHS.PROJECT.path.replace('[id]', project.id)}>
                                <i className={"fas fa-folder"} /> {project.name}
                            </Link>
                            <ul className={styles.subMenuList}>
                                {allDBConnections.filter((dbConn: DBConnection) => {
                                    return dbConn.projectId === project.id
                                }).map((dbConn: DBConnection) => {
                                    return (
                                        <li key={dbConn.id}>
                                            <Link to={Constants.APP_PATHS.DB.path.replace('[id]', dbConn.id)}>
                                                <i className={"fas fa-database"} /> {dbConn.name}
                                            </Link>
                                        </li>
                                    )
                                })}
                                <li>
                                    <Link to={Constants.APP_PATHS.NEW_DB.path.replace('[id]', project.id)}>
                                        <i className={"fas fa-circle-plus"} /> Add DB
                                    </Link>
                                </li>
                            </ul>
                        </li>
                    )
                })}
                <li>
                    <a onClick={() => { setIsShowingCreateProject(true) }}>
                        <i className={"fas fa-folder-plus"} /> Create new project
                    </a>
                </li>
            </ul>
            {isShowingCreateProject && <CreateNewProjectModal onClose={() => { setIsShowingCreateProject(false) }} />}
        </React.Fragment>
    )
}

export default HomeSidebar