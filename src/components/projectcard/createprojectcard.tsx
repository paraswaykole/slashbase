import styles from './projectcard.module.scss'
import React, { useState } from 'react'
import { User } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { selectCurrentUser } from '../../redux/currentUserSlice'
import { createNewProject } from '../../redux/projectsSlice'

type CreateNewProjectCardPropType = {}

const CreateNewProjectCard = (_: CreateNewProjectCardPropType) => {

    const currentUser: User = useAppSelector(selectCurrentUser)

    const [creating, setCreating] = useState(false)
    const [projectName, setProjectName] = useState('')
    const [loading, setLoading] = useState(false)

    if (!currentUser || (currentUser && !currentUser.isRoot)){
        return null
    }

    const dispatch = useAppDispatch()

    const startCreatingProject = async () => {
        if (loading){ 
            return
        }
        setLoading(true)
        await dispatch(createNewProject({projectName}))
        setLoading(false)
        setCreating(false)
        setProjectName('')
    }

    return (
        <React.Fragment>
            {!creating && 
                <button 
                    className="button" 
                    onClick={()=>{
                        setCreating(true)
                    }}>
                    <i className={"fas fa-folder-plus"}/>
                    &nbsp;&nbsp;
                    Create New Project
                </button>
            }
            {
                creating && 
                    <div className={"card "+styles.cardContainer}>
                        <div className="card-content">
                            <div className="field has-addons">
                                <div className="control is-expanded">
                                    <input 
                                        className="input" 
                                        type="text"
                                        placeholder="Enter Project Name"
                                        value={projectName}
                                        onChange={(e: React.ChangeEvent<HTMLInputElement>)=>{setProjectName(e.target.value)}}/>
                                </div>
                                <div className="control">
                                    <button className="button is-primary" onClick={startCreatingProject}>
                                        { loading ? 'Creating' : 'Create'}
                                    </button>
                                </div>
                            </div>
                            <span
                                className={styles.cancelBtn}  
                                onClick={()=>{setCreating(false)}}>
                                Cancel
                            </span>
                        </div>
                    </div>
            }
        </React.Fragment>
    )
}


export default CreateNewProjectCard