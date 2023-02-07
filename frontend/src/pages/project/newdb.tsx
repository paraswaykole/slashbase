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

    const [dbShowPassword, setDBShowPassword] = useState<boolean>(false)
    const [addingError, setAddingError] = useState(false)
    const [adding, setAdding] = useState(false)


    const [data, setData] = useState({
        dbName: "",
        dbType: DBConnType.POSTGRES ,
        dbScheme: "",
        dbHost: "",
        dbPort: "",
        dbDatabase: "",
        dbUsername: "",
        dbPassword: "",
        dbUseSSH: DBConnectionUseSSHType.NONE,
        dbSSHHost: "",
        dbSSHUser: "",
        dbSSHPassword: "",
        dbSSHKeyFile: "",
        dbUseSSL: false,
    })

    const handleChange = (e:React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement >) => {
        const type = e.target.type

        const name = e.target.name

        const value = type === "checkbox"
            ? (e.target as HTMLInputElement).checked 
            : e.target.value

        setData(prevData => ({
            ...prevData,
            [name]: value
        }))
    }



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
            name: data.dbName,
            type: data.dbType,
            scheme: data.dbScheme,
            host: data.dbHost,
            port: data.dbPort,
            password: data.dbPassword,
            user: data.dbUsername,
            dbname: data.dbDatabase,
            useSSH: data.dbUseSSH,
            sshHost: data.dbSSHHost,
            sshUser: data.dbSSHUser,
            sshPassword: data.dbSSHPassword,
            sshKeyFile: data.dbSSHKeyFile,
            useSSL: data.dbUseSSL,
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
        <>
            <h1>Add new database connection</h1>
            <div className="form-container">
                <div className="field">
                    <label className="label">Display Name:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            name='dbName'
                            value={data.dbName}
                            onChange={e => handleChange(e)}
                            placeholder="Enter a display name for database" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database Type:</label>
                    <div className="control">
                        <div className="select">
                            <select name="dbType" onChange={(e: React.ChangeEvent<HTMLSelectElement>) => { setData((prev)=> ({...prev, [e.target.name]:e.target.value, dbScheme :""}))}}>
                                <option value={DBConnType.POSTGRES}>PostgresSQL</option>
                                <option value={DBConnType.MONGO}>MongoDB</option>
                                <option value={DBConnType.MYSQL}>MySQL</option>
                            </select>
                        </div>
                    </div>
                </div>
                {data.dbType === DBConnType.MONGO && <div className="field">
                    <label className="label">Scheme:</label>
                    <div className="control">
                        <div className="select">
                            <select name='dbScheme'  onChange={e => handleChange(e)}>
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
                            name="dbHost"
                            value={data.dbHost}
                            onChange={e => handleChange(e)}
                            placeholder="Enter host" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Port:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            name='dbPort'
                            value={data.dbPort}
                            onChange={e => handleChange(e)}
                            placeholder="Enter port" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database Name:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            name='dbDatabase'
                            value={data.dbDatabase}
                            onChange={e => handleChange(e)}
                            placeholder="Enter database" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database User:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="text"
                            name='dbUsername'
                            value={data.dbUsername}
                            onChange={e => handleChange(e)}
                            placeholder="Enter database username" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Database Password:</label>
                    <div className="control has-icons-right">
                        <input
                            className="input"
                            type={dbShowPassword ? "text" : "password"}
                            name='dbPassword'
                            value={data.dbPassword}
                            placeholder="Enter database password"
                            onChange={e => handleChange(e)}
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
                            name='dbUseSSH'
                            value={data.dbUseSSH}
                            onChange={e => handleChange(e)}
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
                {data.dbType === DBConnType.MONGO && <div className="field">
                    <label className="checkbox">
                        <input
                            name='dbUseSSL'
                            type="checkbox"
                            defaultChecked={false}
                            onChange={e=>handleChange(e)} />
                        &nbsp;Enable SSL
                        <span className="help">If you are connecting to database which enforce/require SSL connection. (Example: Azure CosmosDB)</span>
                    </label>
                </div>}

                {data.dbUseSSH !== DBConnectionUseSSHType.NONE &&
                    <>
                        <div className="field">
                            <label className="label">SSH Host:</label>
                            <div className="control">
                                <input
                                    className="input"
                                    type="text"
                                    name='dbSSHHost'
                                    value={data.dbSSHHost}
                                    onChange={e => handleChange(e)}
                                    placeholder="Enter SSH Host" />
                            </div>
                        </div>
                        <div className="field">
                            <label className="label">SSH User:</label>
                            <div className="control">
                                <input
                                    className="input"
                                    type="text"
                                    name='dbSSHUser'
                                    value={data.dbSSHUser}
                                    onChange={e => handleChange(e)}
                                    placeholder="Enter SSH User" />
                            </div>
                        </div>
                        {(data.dbUseSSH === DBConnectionUseSSHType.PASSWORD || data.dbUseSSH === DBConnectionUseSSHType.PASSKEYFILE) &&
                            < div className="field">
                                <label className="label">SSH Password:</label>
                                <div className="control">
                                    <input
                                        className="input"
                                        type="password"
                                        name='dbSSHPassword'
                                        value={data.dbSSHPassword}
                                        onChange={e => handleChange(e)}
                                        placeholder="Enter SSH Password" />
                                </div>
                            </div>
                        }
                        {(data.dbUseSSH === DBConnectionUseSSHType.KEYFILE || data.dbUseSSH === DBConnectionUseSSHType.PASSKEYFILE) &&
                            <div className="field">
                                <label className="label">SSH Identity File:</label>
                                <div className="control">
                                    <textarea
                                        className="textarea"
                                        name='dbSSHKeyFile'
                                        value={data.dbSSHKeyFile}
                                        onChange={e => handleChange(e)}
                                        placeholder="Paste the contents of SSH Identity File here" />
                                </div>
                            </div>
                        }
                    </>
                }
                <div className="control">
                    {!adding && <button className="button is-primary" onClick={startAddingDB}>Add</button>}
                    {adding && <button className="button is-primary">Adding...</button>}
                    {!adding && addingError && <span className="help is-danger" style={{ display: "inline-flex" }}>&nbsp;&nbsp;{addingError}</span>}
                </div>
            </div>
        </>
    )
}

export default NewDBPage
