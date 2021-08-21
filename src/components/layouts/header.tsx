import styles from './header.module.scss'
import React, { useEffect } from 'react'
import Link from 'next/link'
import { Project, User } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { useRouter } from 'next/router'
import { selectCurrentUser } from '../../redux/currentUserSlice'
import Constants from '../../constants'

type HeaderPropType = {

}

const Header = (_: HeaderPropType) => {
    
    const router = useRouter()

    const currentUser: User = useAppSelector(selectCurrentUser)

    const options = [
        { value: 'home', label: 'Home', path: Constants.APP_PATHS.HOME.as }
    ]

    return (
        <header className={styles.header}>
            <Link {...Constants.APP_PATHS.HOME}>
                <a>
                    <div className={styles.home}>
                        <i className={"fas fa-home"}/>
                    </div>
                </a>
            </Link>
            <div className={styles.headerCenter}>
                <select className={styles.headerSelect} value={router.pathname === '/' ? 'home' : 'home'}>
                    {options.map((x)=>{
                        return <option key={x.value} value={x.value} label={x.label} />
                    })}
                </select>
            </div>
            <div className={styles.headerMenu}>
                { currentUser && 
                <Link {...Constants.APP_PATHS.LOGOUT}>
                    <a>
                        <img className={styles.profileImage} src={currentUser.profileImageUrl} width={40} height={40} /> 
                    </a>
                </Link>
                }
            </div>
        </header>
    )
}


export default Header