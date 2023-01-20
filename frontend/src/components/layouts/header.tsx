import styles from './header.module.scss'
import { Link, useLocation, useNavigate, useParams } from 'react-router-dom'
import Constants from '../../constants'
import React, { useState } from 'react'
import OutsideClickHandler from 'react-outside-click-handler'
import { DBConnection, Project } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { selectIsShowingSidebar, setIsShowingSidebar } from '../../redux/configSlice'
import { selectProjects } from '../../redux/projectsSlice'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import utils from '../../lib/utils'

declare var window: any;

type HeaderPropType = {}

const Header = (_: HeaderPropType) => {

    let location = useLocation()
    const navigate = useNavigate()
    const params = useParams()

    const dispatch = useAppDispatch()

    const projects: Project[] = useAppSelector(selectProjects)
    const currentDBConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)


    const [isShowingDropDown, setIsShowingDropDown] = useState(false)
    const [isShowingNavDropDown, setIsShowingNavDropDown] = useState(false)

    const options = [
        { value: 'home', label: 'Home', path: Constants.APP_PATHS.HOME.path },
        ...projects.map((x: Project) => ({ value: x.id, label: x.name, path: Constants.APP_PATHS.PROJECT.path.replace('[id]', x.id) }))
    ]

    const onNavigate = (option: {
        value: string;
        label: string;
        path: string;
    }) => {
        navigate(option.path)
        setIsShowingNavDropDown(false)
    }

    const toggleSidebar = () => {
        dispatch(setIsShowingSidebar(!isShowingSidebar))
    }

    let currentOption = 'home'
    if (location.pathname.startsWith('/project'))
        currentOption = String(params.id)
    else if (location.pathname.startsWith('/db')) {
        if (currentDBConnection)
            currentOption = currentDBConnection?.projectId
    }

    return (
        <header className={styles.header}>
            <div className={styles.leftBtns}>
                {!isShowingSidebar && <button className={"button is-dark " + [styles.btn].join(' ')} onClick={toggleSidebar}>
                    <i className="fas fa-bars" />
                </button>}
                <Link to={Constants.APP_PATHS.HOME.path}>
                    <button className={"button is-dark " + [styles.btn].join(' ')}>
                        <span className="icon">
                            <i className={`fas fa-home`} />
                        </span>
                    </button>
                </Link>
            </div>
            <div className={styles.headerCenter}>
                <div className={`dropdown${isShowingNavDropDown ? ' is-active' : ''}`}>
                    <div className="dropdown-trigger">
                        <button className={"button is-dark " + styles.btn} aria-haspopup="true" aria-controls="dropdown-menu" onClick={() => { setIsShowingNavDropDown(!isShowingNavDropDown) }}>
                            <span>{options.find(x => x.value === currentOption)?.label}</span>
                            <span className="icon is-small">
                                <i className="fas fa-angle-down" aria-hidden="true"></i>
                            </span>
                        </button>
                    </div>
                    <OutsideClickHandler onOutsideClick={() => { setIsShowingNavDropDown(false) }}>
                        <div className="dropdown-menu" id="dropdown-menu" role="menu">
                            <div className="dropdown-content">
                                {options.map((x) => {
                                    return (
                                        <React.Fragment key={x.value}>
                                            <a onClick={() => { onNavigate(x) }} className={`dropdown-item${x.value === currentOption ? ' is-active' : ''}`}>
                                                {x.label}
                                            </a>
                                            {x.value === 'home' && <hr className="dropdown-divider" />}
                                        </React.Fragment>
                                    )
                                })}
                            </div>
                        </div>
                    </OutsideClickHandler>
                </div>
            </div>
            <div className={styles.headerMenu}>
                <div className={"dropdown is-right" + (isShowingDropDown ? ' is-active' : '')}>
                    <div className="dropdown-trigger" onClick={() => { setIsShowingDropDown(!isShowingDropDown) }}>
                        <button className={"button is-dark " + [styles.btn].join(' ')}>
                            <span className="icon">
                                <i className="fas fa-gear" />
                            </span>
                        </button>
                    </div>
                    <OutsideClickHandler onOutsideClick={() => { setIsShowingDropDown(false) }}>
                        <div className="dropdown-menu" role="menu">
                            <div className="dropdown-content">
                                <a onClick={() => { utils.openInBrowser(Constants.EXTERNAL_PATHS.CHANGELOG) }} className="dropdown-item">
                                    What&apos;s New?
                                </a>
                                <hr className="dropdown-divider" />
                                <Link to={Constants.APP_PATHS.SETTINGS_GENERAL.path} className="dropdown-item">
                                    Settings
                                </Link>
                                <Link to={Constants.APP_PATHS.SETTINGS_SUPPORT.path} className="dropdown-item">
                                    Support
                                </Link>
                            </div>
                        </div>
                    </OutsideClickHandler>
                </div>
            </div>
        </header>
    )
}


export default Header