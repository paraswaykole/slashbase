import React from 'react'
import { Project } from '../../../data/models'
import Constants from '../../../constants'
import { Link } from 'react-router-dom'

type NewDBConnButtonPropType = {
    project: Project
}

const NewDBConnButton = ({ project }: NewDBConnButtonPropType) => {

    return (
        <Link to={Constants.APP_PATHS.NEW_DB.path.replace('[id]', project.id)}>
            <a>
                <button className="button" >
                    <i className={"fas fa-plus-circle"} />
                    &nbsp;&nbsp;
                    Add New DB Connection
                </button>
            </a>
        </Link>
    )
}


export default NewDBConnButton
