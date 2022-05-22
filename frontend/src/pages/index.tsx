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

  return (
    <AppLayout title="Home">
      <h1>All Projects</h1>
      {projects.map((project: Project) => (
        <ProjectCard key={project.id} project={project} />
      ))}
      <CreateNewProjectCard />
    </AppLayout>
  )
}

export default HomePage
