import styles from './header.module.scss'
import React, { useEffect, useState } from 'react'
import Link from 'next/link'
import { DBConnection, Project, User } from '../../data/models'
import { useAppSelector } from '../../redux/hooks'
import { useRouter } from 'next/router'
import { selectCurrentUser } from '../../redux/currentUserSlice'
import Constants from '../../constants'
import { selectProjects } from '../../redux/projectsSlice'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import ProfileImage, { ProfileImageSize } from '../user/profileimage'
import { selectIsShowingSidebar, setIsShowingSidebar } from '../../redux/configSlice'
import { useDispatch } from 'react-redux'

type HeaderPropType = {

}

const Header = (_: HeaderPropType) => {
    
    const router = useRouter()

    const currentUser: User = useAppSelector(selectCurrentUser)
    const projects: Project[] = useAppSelector(selectProjects)
    const currentDBConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)

    const dispatch = useDispatch()

    const [isShowingDropDown, setIsShowingDropDown] = useState(false)
    const [isShowingNavDropDown, setIsShowingNavDropDown] = useState(false)

    const options = [
        { value: 'home', label: 'Home', path: Constants.APP_PATHS.HOME.path },
        ...projects.map((x: Project) => ({value: x.id, label: x.name, path: Constants.APP_PATHS.PROJECT.path.replace('[id]', x.id) }))
    ]

    const onNavigate = (option: {
        value: string;
        label: string;
        path: string;
    }) => {
        router.replace(option.path)
        setIsShowingNavDropDown(false)
    }

    const toggleSidebar = () => {
        dispatch(setIsShowingSidebar(!isShowingSidebar))
    }

    let currentOption = 'home'
    if (router.pathname === Constants.APP_PATHS.PROJECT.path) {
        currentOption = String(router.query.id)
    } else if (router.pathname === Constants.APP_PATHS.NEW_DB.path) {
        currentOption = String(router.query.id)
    } else if (router.pathname === Constants.APP_PATHS.PROJECT_MEMBERS.path) {
        currentOption = String(router.query.id)
    } else if (router.pathname === Constants.APP_PATHS.DB.path) {
        if (currentDBConnection)
            currentOption = currentDBConnection?.projectId
    }

    return (
        <header className={styles.header}>
            <button className={"button is-dark "+[styles.home, styles.btn].join(' ')} onClick={toggleSidebar}>
                <i className={`fas ${isShowingSidebar ? 'fa-angle-double-left' : 'fa-bars'}`}/>
            </button>
            <div className={styles.headerCenter}>
                <div className={`dropdown${isShowingNavDropDown ? ' is-active':''}`}>
                    <div className="dropdown-trigger">
                        <button className={"button is-dark " + styles.btn} aria-haspopup="true" aria-controls="dropdown-menu" onClick={()=>{setIsShowingNavDropDown(!isShowingNavDropDown)}}>
                        <span>{options.find(x => x.value === currentOption)?.label}</span>
                        <span className="icon is-small">
                            <i className="fas fa-angle-down" aria-hidden="true"></i>
                        </span>
                        </button>
                    </div>
                    <div className="dropdown-menu" id="dropdown-menu" role="menu">
                        <div className="dropdown-content">
                            {options.map((x) => {
                                return (
                                    <React.Fragment key={x.value}>
                                        <a onClick={()=>{onNavigate(x)}} className={`dropdown-item${x.value === currentOption?' is-active':''}`}>
                                            {x.label}
                                        </a>
                                        { x.value === 'home' && <hr className="dropdown-divider" /> }
                                    </React.Fragment>
                                )
                            })}
                        </div>
                    </div>
                </div>
            </div>
            <div className={styles.headerMenu}>
                { currentUser && 
                    <div className={"dropdown is-right"+(isShowingDropDown ? ' is-active' : '')}>
                        <div className="dropdown-trigger" onClick={()=>{setIsShowingDropDown(!isShowingDropDown)}}>
                            <ProfileImage imageUrl={currentUser.profileImageUrl} size={ProfileImageSize.SMALL} classes={[styles.profileImage]}/>
                        </div>
                        <div className="dropdown-menu" role="menu">
                            <div className="dropdown-content">
                                <Link href={Constants.APP_PATHS.ACCOUNT.path} as={Constants.APP_PATHS.ACCOUNT.path}>
                                    <a className="dropdown-item">
                                        Account
                                    </a>
                                </Link>
                                <hr className="dropdown-divider"/>
                                <Link href={Constants.APP_PATHS.LOGOUT.path} as={Constants.APP_PATHS.LOGOUT.path}>
                                    <a className="dropdown-item">
                                        Logout
                                    </a>
                                </Link>
                            </div>
                        </div>
                    </div>    
                }
            </div>
        </header>
    )
}


export default Header