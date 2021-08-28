import styles from './projectmembercard.module.scss'
import React from 'react'
import { ProjectMember } from '../../data/models'

type ProjectMemberCardPropType = { 
    member: ProjectMember
}

const ProjectMemberCard = ({member}: ProjectMemberCardPropType) => {

    return (
        <div className={"card "+styles.cardContainer}>
            <div className="card-content">
                <h2>{member.user.email}<span className={"tag is-primary "+styles.roleTag}>{member.role}</span></h2>
            </div>
        </div>
    )
}


export default ProjectMemberCard