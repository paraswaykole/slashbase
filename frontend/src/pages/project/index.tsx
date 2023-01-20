import React, { FunctionComponent, useEffect } from 'react'
import { Link, useParams } from 'react-router-dom'
import DBConnCard from '../../components/cards/dbconncard/dbconncard'
import NewDBConnButton from '../../components/cards/dbconncard/newdbconnectionbutton'
import Constants from '../../constants'
import { DBConnection, Project } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { deleteDBConnectionInProject, getDBConnectionsInProjects, selectDBConnectionsInProject, selectProjects } from '../../redux/projectsSlice'
import emptyStateDatabaseImg from '../../assets/images/empty-state-database.svg'

const ProjectPage: FunctionComponent<{}> = () => {

    const { id } = useParams()

    const dispatch = useAppDispatch()

    const databases = useAppSelector(selectDBConnectionsInProject)
    const projects: Project[] = useAppSelector(selectProjects)
    const project: Project | undefined = projects.find(x => x.id === id)

    useEffect(() => {
        dispatch(getDBConnectionsInProjects({ projectId: String(id) }))
    }, [dispatch, id])

    if (!project) {
        return <h1>Project not found</h1>
    }

    const onDeleteDB = async (dbConnId: string) => {
        dispatch(deleteDBConnectionInProject({ dbConnId }))
    }

    return (
        <React.Fragment>
            <h1>Showing Databases in {project?.name}</h1>
            {project && databases.length === 0 && <div className="empty-state">
                <img className="empty-state-image" src={emptyStateDatabaseImg} />
                <h2>No Database Connections</h2>
                <p>Add a new database connection and connect to the database</p>
                <hr />
            </div>}
            {databases.map((db: DBConnection) => (
                <DBConnCard key={db.id} dbConn={db} onDeleteDB={onDeleteDB} />
            ))}
            {project && <NewDBConnButton project={project} />}
        </React.Fragment>
    )
}

export default ProjectPage
