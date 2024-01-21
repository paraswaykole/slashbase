import styles from '../sidebar.module.scss'
import React from 'react'
import { Link } from "react-router-dom"
import Constants from "../../../constants"
import { User } from '../../../data/models'
import { useAppSelector } from '../../../redux/hooks'
import { selectCurrentUser } from '../../../redux/currentUserSlice'


type SettingSidebarPropType = {}

const SettingSidebar = (_: SettingSidebarPropType) => {

    const currentUser: User = useAppSelector(selectCurrentUser)

    return (
        <React.Fragment>
            <p className="menu-label">
                Settings
            </p>
            <ul className={"menu-list " + styles.menuList}>
                {Constants.Build === 'server' &&
                    <li>
                        <Link
                            to={Constants.APP_PATHS.SETTINGS_ACCOUNT.path}
                            className={location.pathname.startsWith(Constants.APP_PATHS.SETTINGS_ACCOUNT.path) ? 'is-active' : ''}>
                            Account
                        </Link>
                    </li>
                }
                <li>
                    <Link
                        to={Constants.APP_PATHS.SETTINGS_GENERAL.path}
                        className={location.pathname.startsWith(Constants.APP_PATHS.SETTINGS_GENERAL.path) ? 'is-active' : ''}>
                        General
                    </Link>
                </li>
                <li>
                    <Link
                        to={Constants.APP_PATHS.SETTINGS_ADVANCED.path}
                        className={location.pathname.startsWith(Constants.APP_PATHS.SETTINGS_ADVANCED.path) ? 'is-active' : ''}>
                        Advanced
                    </Link>
                </li>
            </ul>
            {Constants.Build === 'server' && currentUser && currentUser.isRoot &&
                <React.Fragment>
                    <p className="menu-label">
                        Manage Team
                    </p>
                    <ul className={"menu-list " + styles.menuList}>
                        <li>
                            <Link
                                to={Constants.APP_PATHS.SETTINGS_USERS.path}
                                className={location.pathname.startsWith(Constants.APP_PATHS.SETTINGS_USERS.path) ? 'is-active' : ''}>
                                Manage Users
                            </Link>
                        </li>
                        <li>
                            <Link
                                to={Constants.APP_PATHS.SETTINGS_ROLES.path}
                                className={location.pathname === Constants.APP_PATHS.SETTINGS_ROLES.path ? 'is-active' : ''}>
                                Manage Roles
                            </Link>
                        </li>
                    </ul>
                </React.Fragment>
            }
            <p className="menu-label">
                Info
            </p>
            <ul className={"menu-list " + styles.menuList}>
                <li>
                    <Link
                        to={Constants.APP_PATHS.SETTINGS_ABOUT.path}
                        className={location.pathname === Constants.APP_PATHS.SETTINGS_ABOUT.path ? 'is-active' : ''}>
                        About
                    </Link>
                </li>
                <li>
                    <Link
                        to={Constants.APP_PATHS.SETTINGS_SUPPORT.path}
                        className={location.pathname === Constants.APP_PATHS.SETTINGS_SUPPORT.path ? 'is-active' : ''}>
                        Support
                    </Link>
                </li>
            </ul>
        </React.Fragment>
    )

}

export default SettingSidebar