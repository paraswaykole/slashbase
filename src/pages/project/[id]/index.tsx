import type { NextPage } from 'next'
import Link from 'next/link'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import DBConnCard from '../../../components/dbconncard/dbconncard'
import NewDBConnButton from '../../../components/dbconncard/newdbconnectionbutton'
import AppLayout from '../../../components/layouts/applayout'
import Constants from '../../../constants'
import { DBConnection, Project } from '../../../data/models'
import apiService from '../../../network/apiService'
import { useAppSelector } from '../../../redux/hooks'
import { selectProjects } from '../../../redux/projectsSlice'

const ProjectPage: NextPage = () => {

  const router = useRouter()
  const { id } = router.query

  const [databases, setDatabases] = useState<DBConnection[]>([])

  const projects: Project[] = useAppSelector(selectProjects)
  const project: Project | undefined = projects.find(x => x.id === id)

  useEffect(()=>{
    (async () => {
      let response = await apiService.getDBConnectionsByProject(String(id))
      if(response.success){
        setDatabases(response.data)
      }
    })()
  }, [router])

  return (
    <AppLayout title={project ? project.name + " | Slashbase": "Slashbase"}>
      <main className="maincontainer">
        <h1>Showing Databases in {project?.name}</h1>
        {databases.map((db: DBConnection) => (
          <DBConnCard key={db.id} dbConn={db}/>
        ))}
        { project && <NewDBConnButton project={project}/> }
        &nbsp;&nbsp;
        { project && <Link href={Constants.APP_PATHS.PROJECT_MEMBERS.path} as={Constants.APP_PATHS.PROJECT_MEMBERS.path.replace('[id]', project.id)}>
          <a>
            <button className="button" >
                <i className={"fas fa-users"}/>
                &nbsp;&nbsp;
                View Project Members
            </button>
          </a>
        </Link> }
      </main>
    </AppLayout>
  )
}

export default ProjectPage
