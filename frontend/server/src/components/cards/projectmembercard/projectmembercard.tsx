import styles from './projectmembercard.module.scss'
import React, { useState } from 'react'
import OutsideClickHandler from 'react-outside-click-handler'
import { ProjectMember } from '../../../data/models'
import ProfileImage from '../../user/profileimage'

type ProjectMemberCardPropType = {
    member: ProjectMember
    isAdmin: Boolean
    onDeleteMember: (dbConnId: string) => void
}

const ProjectMemberCard = ({ member, isAdmin, onDeleteMember }: ProjectMemberCardPropType) => {

    const [showDropdown, setShowDropdown] = useState(false)

    const toggleDropdown = () => {
        setShowDropdown(!showDropdown)
    }

    return (
        <div className={"card " + styles.cardContainer}>
            <div className="card-content">
                <div className="columns">
                    <div className="column is-2">
                        <ProfileImage imageUrl={member.user.profileImageUrl} />
                    </div>
                    <div className="column is-7">
                        <b>{member.user.name ?? member.user.email}</b>
                        {member.user.name && <b className="subtitle is-6"><br />{member.user.email}</b>}
                    </div>
                    <div className="column is-2">
                        <span className={"tag is-primary " + styles.roleTag}>{member.role.name}</span>
                    </div>
                    <div className="column is-1">
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
                                                <a onClick={() => { onDeleteMember(member.user.id) }} className="dropdown-item">
                                                    Remove Member
                                                </a>
                                            </div>
                                        </div>
                                    </OutsideClickHandler>
                                }
                            </div>
                        }
                    </div>
                </div>
            </div>
        </div>
    )
}


export default ProjectMemberCard