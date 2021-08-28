import styles from './projectmembercard.module.scss'
import React from 'react'
import { ProjectMember } from '../../data/models'
import ProfileImage, { ProfileImageSize } from '../user/profileimage'

type ProjectMemberCardPropType = { 
    member: ProjectMember
}

const ProjectMemberCard = ({member}: ProjectMemberCardPropType) => {

    return (
        <div className={"card "+styles.cardContainer}>
            <div className="card-content">
                <div className="columns is-2">
                    <div className="column">
                        <ProfileImage imageUrl={member.user.profileImageUrl} />
                    </div>
                    <div className="column is-8">
                        <h2>{member.user.name ?? member.user.email}</h2>
                        { member.user.name && <h6 className="subtitle is-6">{member.user.email}</h6>}
                    </div>
                    <div className="column is-2">
                        <span className={"tag is-primary "+styles.roleTag}>{member.role}</span>
                    </div>
                </div>
            </div>
        </div>
    )
}


export default ProjectMemberCard