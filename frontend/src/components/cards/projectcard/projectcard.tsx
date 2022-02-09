import styles from './projectcard.module.scss'
import React from 'react'
import { Project } from '../../../data/models'
import Constants from '../../../constants'
import Link from 'next/link'

type ProjectCardPropType = { 
    project: Project
}

const ProjectCard = ({project}: ProjectCardPropType) => {

    return (
        <div className={"card "+styles.cardContainer}>
            <Link href={Constants.APP_PATHS.PROJECT.path} as={Constants.APP_PATHS.PROJECT.path.replace('[id]', project.id)}>
                <a className={styles.cardLink}>
                    <div className="card-content">
                        <b>{project.name}</b>
                    </div>
                </a>
            </Link>
        </div>

    )
}


export default ProjectCard