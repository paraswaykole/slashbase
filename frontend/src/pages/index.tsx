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
      {projectsGrid.map((list: Project[]) => (
        <div className="columns">
          {list.map((project: Project) => (
            <div className="column">
              <ProjectCard key={project.id} project={project} />
            </div>
          ))}
        </div>
      ))}

      <CreateNewProjectCard />
    </AppLayout >
  )
}

export default HomePage
