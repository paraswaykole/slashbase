import React, { FunctionComponent, useEffect, useState } from 'react'
import toast from 'react-hot-toast'
import { useNavigate } from 'react-router-dom'
import Constants from '../../constants'
import { User } from '../../data/models'
import apiService from '../../network/apiService'
import { selectCurrentUser } from '../../redux/currentUserSlice'
import { useAppSelector } from '../../redux/hooks'

const AddNewUserPage: FunctionComponent<{}> = () => {

    const navigate = useNavigate()

    const currentUser: User = useAppSelector(selectCurrentUser)

    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [adding, setAdding] = useState(false)

    useEffect(() => {
        if (currentUser && !currentUser.isRoot) {
            navigate(Constants.APP_PATHS.SETTINGS_ACCOUNT.path)
        }
    }, [currentUser])

    const startAdding = async () => {
        setAdding(true)
        const result = await apiService.addUsers(email, password)
        if (result.success) {
            navigate(Constants.APP_PATHS.SETTINGS_USERS.path)
        } else {
            toast.error(result.error!)
        }
        setAdding(false)
    }

    return (
        <React.Fragment>
            <h1>Add New User</h1>
            <div className="form-container">
                <div className="field">
                    <label className="label">Email:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="email"
                            value={email}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setEmail(e.target.value) }}
                            placeholder="Enter email for new user" />
                    </div>
                </div>
                <div className="field">
                    <label className="label">Password:</label>
                    <div className="control">
                        <input
                            className="input"
                            type="password"
                            value={password}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setPassword(e.target.value) }}
                            placeholder="Enter password for new user" />
                    </div>
                </div>
                <div className="control">
                    {!adding && <button className="button is-primary" onClick={startAdding}>Add</button>}
                    {adding && <button className="button is-primary is-loading">Adding...</button>}
                </div>
            </div>
        </React.Fragment>
    )
}

export default AddNewUserPage
