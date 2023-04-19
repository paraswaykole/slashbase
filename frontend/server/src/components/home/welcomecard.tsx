import React, { FunctionComponent, useEffect, useState } from 'react'
import logo from '../../assets/images/logo-icon.svg'
import CreateNewProjectModal from './createprojectmodal'
import Constants from '../../constants'
import { selectProjects } from '../../redux/projectsSlice'
import { Project } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { useNavigate } from 'react-router-dom'
import { loginUser, selectIsAuthenticated } from '../../redux/currentUserSlice'
import Button from '../ui/Button'


const WelcomeCard: FunctionComponent<{}> = () => {

    const navigate = useNavigate()

    const [isShowingCreateProject, setIsShowingCreateProject] = useState(false)
    const isAuthenticated = useAppSelector(selectIsAuthenticated)
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
                    {isAuthenticated ?
                        <>
                            <h1>Get started</h1>
                            <br />
                            <Button
                                className="is-white"
                                text='Create new project'
                                onClick={() => { setIsShowingCreateProject(true) }}
                                icon={<i className="fas fa-folder-plus"></i>}
                            />
                            <br />
                            <Button
                                className="is-white"
                                text='Add new db'
                                onClick={navigateToNewDB}
                                icon={<i className="fas fa-circle-plus"></i>}
                            />
                        </>
                        :
                        <>
                            <h1>Welcome to Slashbase!</h1>
                            <p>To get started, sign to Slashbase Server</p>
                            <br />
                            <LoginComponent />
                        </>
                    }
                </div>
            </div>
            {isShowingCreateProject && <CreateNewProjectModal onClose={() => { setIsShowingCreateProject(false) }} />}
        </React.Fragment >
    )
}

const LoginComponent = () => {

    const dispatch = useAppDispatch()

    const [userEmail, setUserEmail] = useState('')
    const [userPassword, setUserPassword] = useState('')
    const [loginError, setLoginError] = useState<string | undefined>(undefined)

    const onLoginBtn = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        try {
            await dispatch(loginUser({ email: userEmail, password: userPassword })).unwrap()
        } catch (e: any) {
            setLoginError(e)
        }
    }

    return (
        <div style={{ maxWidth: 500 }}>
            <form onSubmit={onLoginBtn}>
                <div className="field">
                    <label className="label">Email</label>
                    <div className="control has-icons-left">
                        <input
                            className={`input${loginError ? ' is-danger' : ''}`}
                            type="email"
                            placeholder="Enter Email"
                            value={userEmail}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setUserEmail(e.target.value) }}
                        />
                        <span className="icon is-small is-left">
                            <i className="fas fa-envelope"></i>
                        </span>
                    </div>
                </div>
                <div className="field">
                    <label className="label">Password</label>
                    <div className="control has-icons-left">
                        <input
                            className={`input${loginError ? ' is-danger' : ''}`}
                            type="password"
                            placeholder="Enter Password"
                            value={userPassword}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setUserPassword(e.target.value) }}
                        />
                        <span className="icon is-small is-left">
                            <i className="fas fa-lock"></i>
                        </span>
                    </div>
                    {loginError && <span className="help is-danger">{loginError}</span>}
                </div>
                <div className="control">
                    <button className="button is-primary">Login</button>
                </div>
            </form>
        </div>
    )
}


export default WelcomeCard