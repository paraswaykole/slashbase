import type { NextPage } from 'next'
import React from 'react'
import AppLayout from '../components/layouts/applayout'
import CreateNewProjectCard from '../components/cards/projectcard/createprojectcard'
import ProjectCard from '../components/cards/projectcard/projectcard'
import { Project } from '../data/models'
import { useAppSelector } from '../redux/hooks'
import { selectProjects } from '../redux/projectsSlice'

const HomePage: NextPage = () => {

  const projects: Project[] = useAppSelector(selectProjects)
  const projectsGrid: Project[][] = []

  projects.forEach((project: Project, index: number) => {
    let x = Math.floor(index / 3)
    let y = index % 3
    if (!projectsGrid[x]) {
      projectsGrid[x] = []
    }
    projectsGrid[x][y] = project
  })

  return (
    <AppLayout title="Home">
      <h1>All Projects</h1>
      {projects.length === 0 && <div className="empty-state">
        <img className="empty-state-image" src="/static/images/empty-state-project.svg" />
        <h2>No Projects</h2>
        <p>Create a new project to get started</p>
        <hr />
      </div>}
      {projectsGrid.map((list: Project[], index: number) => (
        <div className="columns" key={index}>
          {list.map((project: Project) => (
            <div className="column" key={project.id}>
              <ProjectCard key={project.id} project={project} />
            </div>
          ))}
          {list.length !== 3 && [1, 2, 3].splice(0, 3 - list.length).map((num) => (
            <div className="column" key={num}>
            </div>
          ))}
        </div>
      ))}
      <CreateNewProjectCard />
    </AppLayout >
  )
}

export default HomePage
