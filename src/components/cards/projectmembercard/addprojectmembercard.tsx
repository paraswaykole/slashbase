import styles from './projectmembercard.module.scss'
import React, { useState } from 'react'
import { Project, ProjectMember, User } from '../../../data/models'
import { useAppSelector } from '../../../redux/hooks'
import { selectCurrentUser } from '../../../redux/currentUserSlice'
import { ProjectMemberRole } from '../../../data/defaults'
import apiService from '../../../network/apiService'
import { AddProjectMemberPayload } from '../../../network/payloads'

type AddNewProjectMemberCardPropType = {
    project: Project
    onAdded: (newMember: ProjectMember)=> void
}

const AddNewProjectMemberCard = ({ project, onAdded }: AddNewProjectMemberCardPropType) => {

    const currentUser: User = useAppSelector(selectCurrentUser)

    const [adding, setAdding] = useState(false)
    const [memberEmail, setMemberEmail] = useState('')
    const [memberRole, setMemberRole] = useState<string>(ProjectMemberRole.ANALYST)
    const [loading, setLoading] = useState(false)

    if (!currentUser || (currentUser && !currentUser.isRoot)){
        return null
    }

    const startAddingMember = async () => {
        if (loading){ 
            return
        }
        setLoading(true)
        const payload: AddProjectMemberPayload = {
            email: memberEmail,
            role: memberRole,
        }
        let response = await apiService.addNewProjectMember(project.id, payload)
        if(response.success){
            onAdded(response.data)
        }
        setLoading(false)
        setAdding(false)
        setMemberEmail('')
        setMemberRole(ProjectMemberRole.ANALYST)
    }

    return (
        <React.Fragment>
            {!adding && 
                <button 
                    className="button" 
                    onClick={()=>{
                        setAdding(true)
                    }}>
                    <i className={"fas fa-user-plus"}/>
                    &nbsp;&nbsp;
                    Add New Project Member
                </button>
            }
            {
                adding && 
                    <div className={"card "+styles.cardContainer}>
                        <div className="card-content">
                            <div className="field has-addons">
                                <div className="control is-expanded">
                                    <input 
                                        className="input" 
                                        type="text"
                                        placeholder="Enter member email"
                                        value={memberEmail}
                                        onChange={(e: React.ChangeEvent<HTMLInputElement>)=>{setMemberEmail(e.target.value)}}/>
                                </div>
                                <div className="control">
                                    <div className="select">
                                        <select
                                            value={memberRole}
                                            onChange={(e: React.ChangeEvent<HTMLSelectElement>)=>{
                                                setMemberRole(e.target.value)
                                            }}
                                        >
                                            <option 
                                                value={ProjectMemberRole.ADMIN}
                                                >
                                                Admin
                                            </option>
                                            <option 
                                                value={ProjectMemberRole.DEVELOPER}
                                                >
                                                Developer
                                            </option>
                                            <option 
                                                value={ProjectMemberRole.ANALYST}
                                                >
                                                Analyst
                                            </option>
                                        </select>
                                    </div>
                                </div>
                                <div className="control">
                                    <button className="button is-primary" onClick={startAddingMember}>
                                        { loading ? 'Adding' : 'Add'}
                                    </button>
                                    <button className={"delete "+styles.cancelBtn} onClick={()=>{setAdding(false)}}></button>
                                </div>
                            </div>
                        </div>
                    </div>
            }
        </React.Fragment>
    )
}


export default AddNewProjectMemberCard