import styles from './header.module.scss'
import { Link, useLocation, useNavigate, useParams } from 'react-router-dom'
import Constants from '../../constants'
import React, { useState } from 'react'
import OutsideClickHandler from 'react-outside-click-handler'
import { DBConnection, Project } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { selectIsShowingSidebar, setIsShowingSidebar } from '../../redux/configSlice'
import { selectProjects } from '../../redux/projectsSlice'
import { selectAllDBConnections } from '../../redux/allDBConnectionsSlice'
import { selectDBConnection, getDBDataModels, resetDBDataModels } from '../../redux/dbConnectionSlice'
import utils from '../../lib/utils'
import { Tooltip } from 'react-tooltip'
import 'react-tooltip/dist/react-tooltip.css'

const Header = () => {

    let location = useLocation()
    const navigate = useNavigate()
    const params = useParams()

    const dispatch = useAppDispatch()

    const projects: Project[] = useAppSelector(selectProjects)
    const dbConnections: DBConnection[] = useAppSelector(selectAllDBConnections)
    const currentDBConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)


    const [isShowingDropDown, setIsShowingDropDown] = useState(false)
    const [isShowingNavDropDown, setIsShowingNavDropDown] = useState(false)
    const [isShowingDBDropDown, setIsShowingDBDropdown] = useState(false)


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

    let currentProjectOption: String | undefined = undefined;
    let currentDBOption: string | undefined = undefined;
    if (location.pathname.startsWith('/project'))
        currentProjectOption = String(params.id)
    else if (location.pathname.startsWith('/db')) {
        if (currentDBConnection)
            currentProjectOption = currentDBConnection?.projectId
        currentDBOption = currentDBConnection?.id
    }

    const projectOptions = [

        ...projects.map((x: Project) => ({ value: x.id, label: x.name, path: Constants.APP_PATHS.PROJECT.path.replace('[id]', x.id) }))
    ]
    const dbOptions = [
        ...dbConnections.filter((x: DBConnection) => (x.projectId === currentProjectOption)).map((x: DBConnection) => ({ value: x.id, label: x.name, path: Constants.APP_PATHS.DB.path.replace('[id]', x.id) }))
    ]

    const refreshDataModels = () => {
        dispatch(resetDBDataModels())
        dispatch(getDBDataModels({ dbConnId: currentDBConnection!.id }))
    }

    return (
        <header className={styles.header}>
            <div className={styles.leftBtns}>
                {!isShowingSidebar ? (<button className={"button is-dark " + [styles.btn].join(' ')} onClick={toggleSidebar}>
                    <i className="fas fa-bars" />
                </button>) : (<button className={"button is-dark " + [styles.btn].join(' ')} onClick={toggleSidebar}>
                    <i className="fas fa-bars" />
                </button>)}
                <Link to={Constants.APP_PATHS.HOME.path}>
                    <button className={`button is-dark ` + [styles.btn, currentProjectOption !== undefined ? styles.home : ''].join(' ')}>
                        <span className="icon">
                            <i className={`fas fa-home`} />
                        </span>
                    </button>
                </Link>
                <div className={styles.headerCenter}>
                    {currentProjectOption !== undefined && <div className={`dropdown${isShowingNavDropDown ? ' is-active' : ''}`}>
                        <div className={`dropdown-trigger`}>
                            <button className={"button is-dark " + [styles.btn, styles.bread, currentDBOption === undefined ? styles.breadEnds : ''].join(' ')} aria-haspopup="true" aria-controls="dropdown-menu" onClick={() => { setIsShowingNavDropDown(!isShowingNavDropDown) }}>
                                <span className='icon'>
                                    <i className="fas fa-folder" aria-hidden="true"></i>
                                </span>
                                <span>{projectOptions.find(x => x.value === currentProjectOption)?.label}</span>
                                <span className="icon">
                                    <i className="fas fa-angle-down" aria-hidden="true"></i>
                                </span>
                            </button>
                        </div>
                        <OutsideClickHandler onOutsideClick={() => { setIsShowingNavDropDown(false) }}>
                            <div className="dropdown-menu" id="dropdown-menu" role="menu">
                                <div className="dropdown-content">
                                    {projectOptions.map((x) => {
                                        return (
                                            <React.Fragment key={x.value}>
                                                <a onClick={() => { onNavigate(x) }} className={`dropdown-item${x.value === currentProjectOption ? ' is-active' : ''}`}>
                                                    {x.label}
                                                </a>
                                                {x.value === 'home' && <hr className="dropdown-divider" />}
                                            </React.Fragment>
                                        )
                                    })}
                                </div>
                            </div>
                        </OutsideClickHandler>
                    </div>}
                    {currentProjectOption !== undefined && currentDBOption !== undefined && <div className={`dropdown${isShowingDBDropDown ? ' is-active' : ''}`}>
                        <div className={`dropdown-trigger`}>
                            <button className={"button is-dark " + [styles.btn, styles.bread, styles.dbBread, styles.breadEnds].join(' ')} aria-haspopup="true" aria-controls="dropdown-menu" onClick={() => { setIsShowingDBDropdown(!isShowingDBDropDown) }}>
                                <span className='icon'>
                                    <i className="fas fa-database" aria-hidden="true"></i>
                                </span>
                                <span>{dbOptions.find(x => x.value === currentDBOption)?.label}</span>
                                <span className="icon">
                                    <i className="fas fa-angle-down" aria-hidden="true"></i>
                                </span>
                            </button>
                        </div>
                        <OutsideClickHandler onOutsideClick={() => { setIsShowingDBDropdown(false) }}>
                            <div className="dropdown-menu" id="dropdown-menu" role="menu">
                                <div className="dropdown-content">
                                    {dbOptions.map((x) => {
                                        return (
                                            <React.Fragment key={x.value}>
                                                <a onClick={() => { onNavigate(x) }} className={`dropdown-item${x.value === currentDBOption ? ' is-active' : ''}`}>
                                                    {x.label}
                                                </a>
                                                {x.value === 'home' && <hr className="dropdown-divider" />}
                                            </React.Fragment>
                                        )
                                    })}
                                </div>
                            </div>
                        </OutsideClickHandler>
                        {currentDBOption !== undefined &&
                            <div>
                                <button id="refreshBtn" data-tooltip-content="Refresh data models" className={" button is-dark is-small" + [styles.btn].join(' ')} onClick={refreshDataModels} >
                                    <span className="icon is-small">
                                        <i className="fas fa-sync" />
                                    </span>
                                </button>
                                <Tooltip anchorId="refreshBtn" />
                            </div>
                        }
                    </div>
                    }
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
                                {Constants.Build === 'server' &&
                                    <React.Fragment>
                                        <hr className="dropdown-divider" />
                                        <Link to={Constants.APP_PATHS.LOGOUT.path} className="dropdown-item">
                                            Logout
                                        </Link>
                                    </React.Fragment>
                                }
                            </div>
                        </div>
                    </OutsideClickHandler>
                </div>
            </div>
        </header>
    )
}


export default Header