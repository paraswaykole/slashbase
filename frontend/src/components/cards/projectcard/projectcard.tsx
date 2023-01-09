import styles from './projectcard.module.scss'
import React, { useState } from 'react'
import { Project } from '../../../data/models'
import Constants from '../../../constants'
import OutsideClickHandler from 'react-outside-click-handler'
import { deleteProject } from '../../../redux/projectsSlice'
import ConfirmModal from '../../widgets/confirmModal'
import { Link } from 'react-router-dom'
import { useAppDispatch } from '../../../redux/hooks'

type ProjectCardPropType = {
    project: Project
}

const ProjectCard = ({ project }: ProjectCardPropType) => {

    const [showDropdown, setShowDropdown] = useState(false)
    const [isDeleting, setIsDeleting] = useState(false)
    const dispatch = useAppDispatch()


    const toggleDropdown = () => {
        setShowDropdown(!showDropdown)
    }

    const isAdmin = true
    // const isAdmin = project?.currentMember?.role.name === Constants.ROLES.ADMIN OR isLocalProject

    const onDeleteProject = async () => {
        await dispatch(deleteProject({ projectId: project.id }))
    }

    return (
        <div className={"card " + styles.cardContainer}>
            <Link to={Constants.APP_PATHS.PROJECT.path.replace('[id]', project.id)} className={styles.cardLink}>
                <div className={"card-content " + styles.cardContent}>
                    <b><i className={"fas fa-folder"} />&nbsp;&nbsp;{project.name}</b>
                    {isAdmin &&
                        <div className="dropdown is-right is-active" onClick={(e) => { e.preventDefault() }}>
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
                                            <a onClick={() => { setIsDeleting(true) }} className="dropdown-item">
                                                Delete Project
                                            </a>
                                        </div>
                                    </div>
                                </OutsideClickHandler>
                            }
                        </div>
                    }
                </div>
            </Link>
            {isDeleting && <ConfirmModal
                message={`Are you sure you want to delete  ${project.name}?`}
                onConfirm={onDeleteProject}
                onClose={() => { setIsDeleting(false) }} />}
        </div>

    )
}


export default ProjectCard