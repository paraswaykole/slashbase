import React from 'react'
import { Project } from '../../../data/models'
import Constants from '../../../constants'
import Link from 'next/link'

type NewDBConnButtonPropType = {
    project: Project
}

const NewDBConnButton = ({ project }: NewDBConnButtonPropType) => {

    if (project.currentMember?.role.name !== Constants.ROLES.ADMIN) {
        return null
    }

    return (
        <Link href={Constants.APP_PATHS.NEW_DB.path} as={Constants.APP_PATHS.NEW_DB.path.replace('[id]', project.id)}>
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
