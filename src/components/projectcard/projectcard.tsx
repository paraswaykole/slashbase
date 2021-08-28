import styles from './projectcard.module.scss'
import React from 'react'
import { Project } from '../../data/models'
import Constants from '../../constants'
import Link from 'next/link'

type ProjectCardPropType = { 
    project: Project
}

const ProjectCard = ({project}: ProjectCardPropType) => {

    return (
        <Link href={Constants.APP_PATHS.PROJECT.path} as={Constants.APP_PATHS.PROJECT.path.replace('[id]', project.id)}>
            <a>
                <div className={"card "+styles.cardContainer}>
                    <div className="card-content">
                        <h2>{project.name}</h2>
                    </div>
                </div>
            </a>
        </Link>
    )
}


export default ProjectCard