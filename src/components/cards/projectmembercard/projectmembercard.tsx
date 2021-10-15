import styles from './projectmembercard.module.scss'
import React from 'react'
import { ProjectMember } from '../../../data/models'
import ProfileImage, { ProfileImageSize } from '../../user/profileimage'

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
                        <b>{member.user.name ?? member.user.email}</b>
                        { member.user.name && <b className="subtitle is-6"><br/>{member.user.email}</b>}
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