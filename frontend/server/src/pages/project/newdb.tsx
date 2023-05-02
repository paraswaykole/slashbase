import React, { FunctionComponent, useState } from 'react'
import { Project } from '../../data/models'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { selectProjects } from '../../redux/projectsSlice'
import { AddDBConnPayload } from '../../network/payloads'
import { DBConnectionUseSSHType, DBConnType } from '../../data/defaults'
import { addNewDBConn } from '../../redux/allDBConnectionsSlice'
import Constants from '../../constants'
import { useNavigate, useParams } from 'react-router-dom'
import InputTextField from '../../components/ui/Input/InputField'
import PasswordInputField from '../../components/ui/Input/PasswordInputField'

const NewDBPage: FunctionComponent<{}> = () => {

    const { id } = useParams()
    const navigate = useNavigate()

    const dispatch = useAppDispatch()
    const projects: Project[] = useAppSelector(selectProjects)
    const project = projects.find(x => x.id === id)
    const [addingError, setAddingError] = useState(false)
    const [adding, setAdding] = useState(false)
    const [inputError, setInputError] = useState({
        error_1: false, error_2: false, error_3: false, error_4: false
    })
    const [data, setData] = useState({
        dbName: "",
        dbType: DBConnType.POSTGRES,
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

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
        const type = e.target.type
        const name = e.target.name;

        switch (name) {
            case "dbName":
                (e.target.value.trim().length > 0) ? setInputError({ ...inputError, error_1: false }) : setInputError({ ...inputError, error_1: true })
                break;
            case "dbHost":
                (e.target.value.trim().length > 0) ? setInputError({ ...inputError, error_2: false }) : setInputError({ ...inputError, error_2: true })
                break;
            case "dbPort":
                (e.target.value.trim().length > 0) ? setInputError({ ...inputError, error_3: false }) : setInputError({ ...inputError, error_3: true })
                break;
            case "dbDatabase":
                (e.target.value.trim().length > 0) ? setInputError({ ...inputError, error_4: false }) : setInputError({ ...inputError, error_4: true })
                break;
        }

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
            var f1 = false, f2 = false, f3 = false, f4 = false;
            (payload.name.length === 0) ? f1 = true : f1 = false;
            (payload.host.length === 0) ? f2 = true : f2 = false;
            (payload.port.length === 0) ? f3 = true : f3 = false;
            (payload.dbname.length === 0) ? f4 = true : f4 = false;
            setInputError({
                ...inputError,
                error_1: f1,
                error_2: f2,
                error_3: f3,
                error_4: f4
            })
            setAddingError(e)
        }

        setAdding(false)
    }

    const normal = {
        border: ''
    }

    const inputStyle = {
        border: '1px solid red'
    }

    return (
        <>
            <h1>Add new database connection</h1>
            <div className="form-container">
                <InputTextField
                    label='Display Name: '
                    name='dbName'
                    value={data.dbName}
                    onChange={e => handleChange(e)}
                    placeholder="Enter a display name for database"
                    style={inputError.error_1 ? inputStyle : normal}
                />
                <div className="field">
                    <label className="label">Database Type:</label>
                    <div className="control">
                        <div className="select">
                            <select name="dbType" style={normal} onChange={(e: React.ChangeEvent<HTMLSelectElement>) => { setData((prev) => ({ ...prev, [e.target.name]: e.target.value, dbScheme: "" })) }}>
                                <option value={DBConnType.POSTGRES}>PostgreSQL</option>
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
                            <select name='dbScheme' style={normal} onChange={e => handleChange(e)}>
                                <option value="default">Select scheme</option>
                                <option value="mongodb">mongodb</option>
                                <option value="mongodb+srv">mongodb+srv</option>
                            </select>
                        </div>
                    </div>
                </div>}
                <InputTextField
                    label='Host:'
                    name="dbHost"
                    value={data.dbHost}
                    onChange={e => handleChange(e)}
                    placeholder="Enter host"
                    style={inputError.error_2 ? inputStyle : normal}
                />
                <InputTextField
                    label='Port:'
                    name="dbPort"
                    value={data.dbPort}
                    onChange={e => handleChange(e)}
                    placeholder="Enter Port"
                    style={inputError.error_3 ? inputStyle : normal}
                />
                <InputTextField
                    label='Database Name:'
                    name="dbDatabase"
                    value={data.dbDatabase}
                    onChange={e => handleChange(e)}
                    placeholder="Enter Database"
                    style={inputError.error_4 ? inputStyle : normal}
                />
                <InputTextField
                    label='Database User:'
                    name="dbUsername"
                    value={data.dbUsername}
                    onChange={e => handleChange(e)}
                    placeholder="Enter Database username"
                    style={normal}
                />
                <PasswordInputField
                    label='Database Password:'
                    name='dbPassword'
                    value={data.dbPassword}
                    onChange={e => handleChange(e)}
                    placeholder="Enter database password"
                    style={normal}
                />
                <div className="field">
                    <label className="label">Use SSH:</label>
                    <div className="select">
                        <select
                            name='dbUseSSH'
                            value={data.dbUseSSH}
                            onChange={e => handleChange(e)}
                            style={normal}
                        >
                            <option
                                value={DBConnectionUseSSHType.NONE} >
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
                            onChange={e => handleChange(e)}
                        />
                        &nbsp;Enable SSL
                        <span className="help">If you are connecting to database which enforce/require SSL connection. (Example: Azure CosmosDB)</span>
                    </label>
                </div>}

                {data.dbUseSSH !== DBConnectionUseSSHType.NONE &&
                    <>
                        <InputTextField
                            label='SSH Host:'
                            name="dbSSHHost"
                            value={data.dbSSHHost}
                            onChange={e => handleChange(e)}
                            placeholder="Enter SSH Host"
                            style={normal}
                        />
                        <InputTextField
                            label='SSH User:'
                            name="dbSSHUser"
                            value={data.dbSSHUser}
                            onChange={e => handleChange(e)}
                            placeholder="Enter SSH User"
                            style={normal}
                        />
                        {(data.dbUseSSH === DBConnectionUseSSHType.PASSWORD || data.dbUseSSH === DBConnectionUseSSHType.PASSKEYFILE) &&
                            <PasswordInputField
                                label='SSH Password:'
                                name='dbSSHPassword'
                                value={data.dbSSHPassword}
                                onChange={e => handleChange(e)}
                                placeholder="Enter SSH Password"
                                style={normal}
                            />
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
                                        placeholder="Paste the contents of SSH Identity File here"
                                        style={normal}
                                    />
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
