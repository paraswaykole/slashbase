import React, { FunctionComponent, useEffect, useState } from 'react'
import AddNewProjectMemberCard from '../../components/cards/projectmembercard/addprojectmembercard'
import ProjectMemberCard from '../../components/cards/projectmembercard/projectmembercard'
import { Project, ProjectMember } from '../../data/models'
import apiService from '../../network/apiService'
import { useAppSelector } from '../../redux/hooks'
import { selectProjects } from '../../redux/projectsSlice'
import Constants from '../../constants'
import { useParams } from 'react-router-dom'

const ProjectMembersPage: FunctionComponent<{}> = () => {

    const { id } = useParams()

    const [projectMembers, setProjectMembers] = useState<ProjectMember[]>([])

    const projects: Project[] = useAppSelector(selectProjects)
    const project: Project | undefined = projects.find(x => x.id === id)

    useEffect(() => {
        (async () => {
            let response = await apiService.getProjectMembers(String(id))
            if (response.success) {
                setProjectMembers(response.data)
            }
        })()
    }, [id])

    const onDeleteMember = async (userId: string) => {
        let response = await apiService.deleteProjectMember(project!.id, userId)
        if (response.success) {
            setProjectMembers(projectMembers.filter(pm => pm.user.id !== userId))
        }
    }

    return (
        <React.Fragment>
            <h1>Showing Members in {project?.name}</h1>
            {projectMembers.map((pm: ProjectMember) => (
                <ProjectMemberCard key={pm.id} member={pm} isAdmin={project?.currentMember?.role.name == Constants.ROLES.ADMIN} onDeleteMember={onDeleteMember} />
            ))}
            {project &&
                <AddNewProjectMemberCard
                    project={project}
                    onAdded={(newMember: ProjectMember) => {
                        setProjectMembers([...projectMembers, newMember])
                    }} />
            }
        </React.Fragment>
    )
}

export default ProjectMembersPage
