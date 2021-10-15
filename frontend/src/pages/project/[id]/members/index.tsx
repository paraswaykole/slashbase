import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../../../components/layouts/applayout'
import AddNewProjectMemberCard from '../../../../components/cards/projectmembercard/addprojectmembercard'
import ProjectMemberCard from '../../../../components/cards/projectmembercard/projectmembercard'
import { Project, ProjectMember } from '../../../../data/models'
import apiService from '../../../../network/apiService'
import { useAppSelector } from '../../../../redux/hooks'
import { selectProjects } from '../../../../redux/projectsSlice'

const ProjectMembersPage: NextPage = () => {

  const router = useRouter()
  const { id } = router.query

  const [projectMembers, setProjectMembers] = useState<ProjectMember[]>([])

  const projects: Project[] = useAppSelector(selectProjects)
  const project: Project | undefined = projects.find(x => x.id === id)

  useEffect(()=>{
    (async () => {
      let response = await apiService.getProjectMembers(String(id))
      if(response.success){
        setProjectMembers(response.data)
      }
    })()
  }, [router])

  return (
    <AppLayout title={project ? project.name + " | Slashbase": "Slashbase"}>
      <h1>Showing Members in {project?.name}</h1>
      {projectMembers.map((pm: ProjectMember) => (
        <ProjectMemberCard key={pm.id} member={pm}/>
      ))}
      { project && 
          <AddNewProjectMemberCard 
            project={project} 
            onAdded={(newMember: ProjectMember )=>{
              setProjectMembers([...projectMembers, newMember])
            }}/> 
      }
    </AppLayout>
  )
}

export default ProjectMembersPage
