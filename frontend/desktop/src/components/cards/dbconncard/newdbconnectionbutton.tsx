import React from 'react'
import { Project } from '../../../data/models'
import Constants from '../../../constants'
import { Link } from 'react-router-dom'
import Button from '../../ui/Button'

type NewDBConnButtonPropType = {
    project: Project
}

const NewDBConnButton = ({ project }: NewDBConnButtonPropType) => {

    return (
        <Link to={Constants.APP_PATHS.NEW_DB.path.replace('[id]', project.id)}>
            <a>
            <Button text="Add New DB Connection" icon={<i className={"fas fa-plus-circle"} />} />
            </a>
        </Link>
    )
}


export default NewDBConnButton
