import React, { FunctionComponent } from 'react'
import { Project } from '../../data/models'
import { useAppSelector } from '../../redux/hooks'
import { selectProjects } from '../../redux/projectsSlice'
import CreateNewProjectCard from '../cards/projectcard/createprojectcard'
import ProjectCard from '../cards/projectcard/projectcard'
import emptyStateProjectImg from '../../assets/images/empty-state-project.svg'


const Projects: FunctionComponent<{}> = () => {

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
        <React.Fragment>
            <h1>Projects</h1>
            {projects.length === 0 && <div className="empty-state">
                <img className="empty-state-image" src={emptyStateProjectImg} />
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
        </React.Fragment >
    )
}


export default Projects