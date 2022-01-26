import styles from './projectmembercard.module.scss'
import React, { useRef, useState } from 'react'
import InfiniteScroll from 'react-infinite-scroll-component'
import { Project, ProjectMember, User } from '../../../data/models'
import { useAppSelector } from '../../../redux/hooks'
import { selectCurrentUser } from '../../../redux/currentUserSlice'
import { ProjectMemberRole } from '../../../data/defaults'
import apiService from '../../../network/apiService'
import { AddProjectMemberPayload } from '../../../network/payloads'
import ProfileImage from '../../user/profileimage'

type AddNewProjectMemberCardPropType = {
    project: Project
    onAdded: (newMember: ProjectMember)=> void
}

const AddNewProjectMemberCard = ({ project, onAdded }: AddNewProjectMemberCardPropType) => {

    const currentUser: User = useAppSelector(selectCurrentUser)

    const [showing, setShowing] = useState(false)
    const searchTerm = useRef<HTMLInputElement>(null)
    const [searchResults, setSearchResults] = useState<User[]>([])
    const [searchOffset, setSearchOffset] = useState<number>(0)
    const [memberEmail, setMemberEmail] = useState('')
    const [memberRole, setMemberRole] = useState<string>(ProjectMemberRole.ANALYST)
    const [searching, setSearching] = useState(false)
    const [adding, setAdding] = useState(false)

    if (!currentUser || (currentUser && !currentUser.isRoot)){
        return null
    } 

    const onSearch = async () => {
        if (searchTerm.current!.value === ''){
            if (searchResults.length !== 0) {
                setSearchResults([])
                setSearchOffset(0)
            }
            return
        }
        if (searching) return
        fetchSearchUsers(0)
    }

    const fetchSearchUsers = async (offset: number) => {
        setSearching(true)
        await new Promise(resolve => setTimeout(resolve, 1000))
        if (searchTerm.current!.value === ''){
            setSearching(false)
            return
        }
        let results = await apiService.searchUsers(searchTerm.current!.value, offset)
        if (results.success) {
            if (offset == 0) {
                setSearchResults(results.data.list)   
            } else {
                setSearchResults([...searchResults, ...results.data.list])
            }
            setSearchOffset(results.data.next)
            setSearching(false)
        }
    }

    const startAddingMember = async () => {
        if (adding){ 
            return
        }
        setAdding(true)
        const payload: AddProjectMemberPayload = {
            email: memberEmail,
            role: memberRole,
        }
        let response = await apiService.addNewProjectMember(project.id, payload)
        if(response.success){
            onAdded(response.data)
        }
        clearAndExit()
    }

    const clearAndExit = () => {
        setAdding(false)
        setShowing(false)
        setMemberEmail('')
        setSearchResults([])
        setSearchOffset(0)
        setMemberRole(ProjectMemberRole.ANALYST)
    }

    return (
        <React.Fragment>
            {!showing && 
                <button 
                    className="button" 
                    onClick={()=>{
                        setShowing(true)
                    }}>
                    <i className={"fas fa-user-plus"}/>
                    &nbsp;&nbsp;
                    Add New Project Member
                </button>
            }
            {
                showing && 
                    <div className={"card "+styles.cardContainer}>
                        <div className="card-content">
                            { memberEmail == '' ? 
                            <>
                                <div className="field has-addons">
                                    <div className="control is-expanded">
                                        <input 
                                            className="input" 
                                            type="text"
                                            placeholder="Search users by name or email"
                                            ref={searchTerm}
                                            onChange={onSearch}/>
                                    </div>
                                    <div className="control">
                                        <button className={"delete "+styles.cancelBtn} onClick={clearAndExit}></button>
                                    </div>
                                </div>
                                { searchResults.length >  0 && 
                                        <InfiniteScroll
                                            dataLength={searchResults.length}
                                            next={()=>fetchSearchUsers(searchOffset)}
                                            hasMore={searchOffset !== -1}
                                            loader={
                                                <p>Loading...</p>
                                            }
                                            height={215}
                                            className={styles.searchResults}
                                            >
                                                {searchResults.map((user: User) => 
                                                    (<div className={"columns is-2 "+styles.searchItem} onClick={()=>{setMemberEmail(user.email)}}>
                                                        <div className="column">
                                                            <ProfileImage imageUrl={user.profileImageUrl} />
                                                        </div>
                                                        <div className="column is-10">
                                                            <b>{user.name ?? user.email}</b>
                                                            { user.name && <b className="subtitle is-6"><br/>{user.email}</b>}
                                                        </div>
                                                    </div>)
                                                )}
                                        </InfiniteScroll>
                                   
                                }
                            </>
                            :
                            <div className="field has-addons">
                                <div className="control is-expanded">
                                    <input 
                                        className="input" 
                                        type="text"
                                        placeholder="No user selected"
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
                                        { adding ? 'Adding' : 'Add'}
                                    </button>
                                    <button className={"delete "+styles.cancelBtn} onClick={clearAndExit}></button>
                                </div>
                            </div>
                            }
                        </div>
                    </div>
            }
        </React.Fragment>
    )
}


export default AddNewProjectMemberCard