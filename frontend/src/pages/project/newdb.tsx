import React, { FunctionComponent, useState } from 'react'
import { Project } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { selectProjects } from '../../redux/projectsSlice'
import { AddDBConnPayload } from '../../events/payloads'
import { DBConnectionUseSSHType, DBConnType } from '../../data/defaults'
import { addNewDBConn } from '../../redux/allDBConnectionsSlice'
import Constants from '../../constants'
import { useNavigate, useParams } from 'react-router-dom'

const NewDBPage: FunctionComponent<{}> = () => {

    const { id } = useParams()
    const navigate = useNavigate()

    const dispatch = useAppDispatch()
    const projects: Project[] = useAppSelector(selectProjects)
    const project = projects.find(x => x.id === id)

    const [dbName, setDBName] = useState('')
    const [dbHost, setDBHost] = useState('')
    const [dbScheme, setDBScheme] = useState('')
    const [dbType, setDBType] = useState<string>(DBConnType.POSTGRES)
    const [dbPort, setDBPort] = useState('')
    const [dbDatabase, setDBDatabase] = useState('')
    const [dbUsername, setDBUsername] = useState('')
    const [dbPassword, setDBPassword] = useState('')
    const [dbShowPassword, setDBShowPassword] = useState<boolean>(false)
    const [dbUseSSH, setUseSSH] = useState<string>(DBConnectionUseSSHType.NONE)
    const [dbSSHHost, setSSHHost] = useState('')
    const [dbSSHUser, setSSHUser] = useState('')
    const [dbSSHPassword, setSSHPassword] = useState('')
    const [dbSSHKeyFile, setSSHKeyFile] = useState('')
    const [addingError, setAddingError] = useState(false)
    const [adding, setAdding] = useState(false)
    const [dbUseSSL, setDBUseSSL] = useState(false)

    if (!project) {
        return <h1>Project not found</h1>
    }

    // if (project.currentMember?.role.name !== Constants.ROLES.ADMIN) {
    // 	return <DefaultErrorPage statusCode={401} title="Unauthorized" />
    // }

    const startAddingDB = async () => {
        setAdding(true)
        const payload: AddDBConnPayload = {
            projectId: project.id,
            name: dbName,
            type: dbType,
            scheme: dbScheme,
            host: dbHost,
            port: dbPort,
            password: dbPassword,
            user: dbUsername,
            dbname: dbDatabase,
            useSSH: dbUseSSH,
            sshHost: dbSSHHost,
            sshUser: dbSSHUser,
            sshPassword: dbSSHPassword,
            sshKeyFile: dbSSHKeyFile,
            useSSL: dbUseSSL,
        }
        try {
            await dispatch(addNewDBConn(payload)).unwrap()
            navigate(Constants.APP_PATHS.PROJECT.path.replace('[id]', project.id))
        } catch (e: any) {
            setAddingError(e)
        }
        setAdding(false)
    }

    return (
        <React.Fragment>
            <h1>Add new database connection</h1>
            <div className="form-container">
                <div className="field">
                    <label className="label">Display Name:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            value={dbName}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setDBName(e.target.value) }}
                            placeholder="Enter a display name for database" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database Type:</label>
                    <div className="control">
                        <div className="select">
                            <select onChange={(e: React.ChangeEvent<HTMLSelectElement>) => { setDBType(e.target.value); setDBScheme('') }}>
                                <option value={DBConnType.POSTGRES}>PostgresSQL</option>
                                <option value={DBConnType.MONGO}>MongoDB</option>
                                <option value={DBConnType.MYSQL}>MySQL</option>
                            </select>
                        </div>
                    </div>
                </div>
                {dbType === DBConnType.MONGO && <div className="field">
                    <label className="label">Scheme:</label>
                    <div className="control">
                        <div className="select">
                            <select onChange={(e: React.ChangeEvent<HTMLSelectElement>) => { setDBScheme(e.target.value) }}>
                                <option value="default">Select scheme</option>
                                <option value="mongodb">mongodb</option>
                                <option value="mongodb+srv">mongodb+srv</option>
                            </select>
                        </div>
                    </div>
                </div>}
                <div className="field">
                    <label className="label">Host:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            value={dbHost}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setDBHost(e.target.value) }}
                            placeholder="Enter host" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Port:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            value={dbPort}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setDBPort(e.target.value) }}
                            placeholder="Enter port" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database Name:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            value={dbDatabase}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setDBDatabase(e.target.value) }}
                            placeholder="Enter database" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database User:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            value={dbUsername}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setDBUsername(e.target.value) }}
                            placeholder="Enter database username" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database Password:</label>
                    <div className="control has-icons-right">
                        <input
                            className="input"
                            type={dbShowPassword ? "text" : "password"}
                            value={dbPassword}
                            placeholder="Enter database password"
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                                setDBPassword(e.target.value);
                            }}
                        />
                        <span
                            className="control icon is-clickable is-small is-right"
                            onClick={() => setDBShowPassword((prev) => !prev)}
                        >
                            <i className={dbShowPassword ? "fas fa-eye" : "fas fa-eye-slash"} />
                        </span>
                    </div>
                </div>
                <div className="field">
                    <label className="label">Use SSH:</label>
                    <div className="select">
                        <select
                            value={dbUseSSH}
                            onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                                setUseSSH(e.target.value)
                            }}
                        >
                            <option
                                value={DBConnectionUseSSHType.NONE}>
                                None
                            </option>
                            <option
                                value={DBConnectionUseSSHType.PASSWORD}>
                                Password
                            </option>
                            <option
                                value={DBConnectionUseSSHType.KEYFILE}>
                                Identity File
                            </option>
                            <option
                                value={DBConnectionUseSSHType.PASSKEYFILE}>
                                Identity File with Password
                            </option>
                        </select>
                    </div>
                </div>
                {dbType === DBConnType.MONGO && <div className="field">
                    <label className="checkbox">
                        <input
                            type="checkbox"
                            defaultChecked={false}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setDBUseSSL(e.target.checked) }} />
                        &nbsp;Enable SSL
                        <span className="help">If you are connecting to database which enforce/require SSL connection. (Example: Azure CosmosDB)</span>
                    </label>
                </div>}

                {dbUseSSH !== DBConnectionUseSSHType.NONE &&
                    <React.Fragment>
                        <div className="field">
                            <label className="label">SSH Host:</label>
                            <div className="control">
                                <input
                                    className="input"
                                    type="text"
                                    value={dbSSHHost}
                                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setSSHHost(e.target.value) }}
                                    placeholder="Enter SSH Host" />
                            </div>
                        </div>
                        <div className="field">
                            <label className="label">SSH User:</label>
                            <div className="control">
                                <input
                                    className="input"
                                    type="text"
                                    value={dbSSHUser}
                                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setSSHUser(e.target.value) }}
                                    placeholder="Enter SSH User" />
                            </div>
                        </div>
                        {(dbUseSSH === DBConnectionUseSSHType.PASSWORD || dbUseSSH === DBConnectionUseSSHType.PASSKEYFILE) &&
                            < div className="field">
                                <label className="label">SSH Password:</label>
                                <div className="control">
                                    <input
                                        className="input"
                                        type="password"
                                        value={dbSSHPassword}
                                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setSSHPassword(e.target.value) }}
                                        placeholder="Enter SSH Password" />
                                </div>
                            </div>
                        }
                        {(dbUseSSH === DBConnectionUseSSHType.KEYFILE || dbUseSSH === DBConnectionUseSSHType.PASSKEYFILE) &&
                            <div className="field">
                                <label className="label">SSH Identity File:</label>
                                <div className="control">
                                    <textarea
                                        className="textarea"
                                        value={dbSSHKeyFile}
                                        onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => { setSSHKeyFile(e.target.value) }}
                                        placeholder="Paste the contents of SSH Identity File here" />
                                </div>
                            </div>
                        }
                    </React.Fragment>
                }
                <div className="control">
                    {!adding && <button className="button is-primary" onClick={startAddingDB}>Add</button>}
                    {adding && <button className="button is-primary">Adding...</button>}
                    {!adding && addingError && <span className="help is-danger" style={{ display: "inline-flex" }}>&nbsp;&nbsp;{addingError}</span>}
                </div>
            </div>
        </React.Fragment>
    )
}

export default NewDBPage
