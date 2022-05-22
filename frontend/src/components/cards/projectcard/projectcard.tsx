import styles from './projectcard.module.scss'
import React, { useState } from 'react'
import { Project } from '../../../data/models'
import Constants from '../../../constants'
import Link from 'next/link'
import OutsideClickHandler from 'react-outside-click-handler'
import { ProjectMemberRole } from '../../../data/defaults'
import { deleteProject } from '../../../redux/projectsSlice'
import { useDispatch } from 'react-redux'

type ProjectCardPropType = {
    project: Project
}

const ProjectCard = ({ project }: ProjectCardPropType) => {

    const [showDropdown, setShowDropdown] = useState(false)
    const dispatch = useDispatch()


    const toggleDropdown = () => {
        setShowDropdown(!showDropdown)
    }

    const isAdmin = project?.currentMember?.role === ProjectMemberRole.ADMIN

    const onDeleteProject = async () => {
        await dispatch(deleteProject({ projectId: project.id }))
    }

    return (
        <div className={"card " + styles.cardContainer}>
            <Link href={Constants.APP_PATHS.PROJECT.path} as={Constants.APP_PATHS.PROJECT.path.replace('[id]', project.id)}>
                <a className={styles.cardLink}>
                    <div className={"card-content " + styles.cardContent}>
                        <b>{project.name}</b>
                        {isAdmin &&
                            <div className="dropdown is-active" onClick={(e) => { e.preventDefault() }}>
                                <div className="dropdown-trigger">
                                    <button className="button" aria-haspopup="true" aria-controls="dropdown-menu" onClick={toggleDropdown}>
                                        <span className="icon is-small">
                                            <i className="fas fa-ellipsis-v" aria-hidden="true"></i>
                                        </span>
                                    </button>
                                </div>
                                {showDropdown &&
                                    <OutsideClickHandler onOutsideClick={() => { setShowDropdown(false) }}>
                                        <div className="dropdown-menu" id="dropdown-menu" role="menu">
                                            <div className="dropdown-content">
                                                <a onClick={onDeleteProject} className="dropdown-item">
                                                    Delete Project
                                                </a>
                                            </div>
                                        </div>
                                    </OutsideClickHandler>
                                }
                            </div>
                        }
                    </div>
                </a>
            </Link>
        </div>

    )
}


export default ProjectCard