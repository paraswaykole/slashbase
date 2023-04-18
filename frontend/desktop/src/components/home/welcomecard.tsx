import React, { FunctionComponent, useState } from 'react'
import logo from '../../assets/images/logo-icon.svg'
import CreateNewProjectModal from './createprojectmodal'
import utils from '../../lib/utils'
import Constants from '../../constants'
import { selectProjects } from '../../redux/projectsSlice'
import { Project } from '../../data/models'
import { useAppSelector } from '../../redux/hooks'
import { useNavigate } from 'react-router-dom'


const WelcomeCard: FunctionComponent<{}> = () => {

    const navigate = useNavigate()

    const [isShowingCreateProject, setIsShowingCreateProject] = useState(false)
    const projects: Project[] = useAppSelector(selectProjects)

    const navigateToNewDB = () => {
        if (projects.length > 0) {
            navigate(Constants.APP_PATHS.NEW_DB.path.replace('[id]', projects[0].id))
        }
    }

    return (
        <React.Fragment>
            <div className='card'>
                <div className="card-content">
                    <img src={logo} width={45} alt="slashbase logo" />
                    <h1>Get started</h1>
                    <br />
                    <button className="button is-white" onClick={() => { setIsShowingCreateProject(true) }}>
                        <span className="icon is-small">
                            <i className="fas fa-folder-plus"></i>
                        </span>
                        <span>Create new project</span>
                    </button><br />
                    <button className="button is-white" onClick={navigateToNewDB}>
                        <span className="icon is-small">
                            <i className="fas fa-circle-plus"></i>
                        </span>
                        <span>Add new db</span>
                    </button>
                    <hr />
                    <div>
                        <h3>Have any feedback?</h3>
                        <p>Use any of the channels below to share your feedback or feature requests.</p>
                        <div className="buttons">
                            <button className="button is-small is-white" onClick={() => { utils.openInBrowser(Constants.EXTERNAL_PATHS.DISCORD_COMMUNITY) }}>
                                <span className="icon is-small">
                                    <i className="fab fa-discord"></i>
                                </span>
                                <span>Discord</span>
                            </button>
                            <button className="button is-small is-white" onClick={() => { utils.openInBrowser(Constants.EXTERNAL_PATHS.REPORT_BUGS) }}>
                                <span className="icon is-small">
                                    <i className="fab fa-github"></i>
                                </span>
                                <span>GitHub</span>
                            </button>
                            <button className="button is-small is-white" onClick={() => { utils.openInBrowser(Constants.EXTERNAL_PATHS.SUPPORT_MAIL) }}>
                                <span className="icon is-small">
                                    <i className="fas fa-envelope"></i>
                                </span>
                                <span>E-mail</span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            {isShowingCreateProject && <CreateNewProjectModal onClose={() => { setIsShowingCreateProject(false) }} />}
        </React.Fragment >
    )
}


export default WelcomeCard