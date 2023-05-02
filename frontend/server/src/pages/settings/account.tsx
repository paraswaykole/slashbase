import axios from 'axios'
import React, { FunctionComponent, useEffect, useState } from 'react'
import ProfileImage, { ProfileImageSize } from '../../components/user/profileimage'
import { User } from '../../data/models'
import { editUser, selectCurrentUser } from '../../redux/currentUserSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'

const AccountPage: FunctionComponent<{}> = () => {

    const currentUser: User = useAppSelector(selectCurrentUser)

    const [editableUser, setEditableUser] = useState<User | null>(null)
    const [saving, setSaving] = useState(false)
    const [savingError, setSavingError] = useState(false)

    const dispatch = useAppDispatch()

    useEffect(() => {
        setEditableUser(currentUser)
    }, [currentUser])

    if (!currentUser) {
        return <h1>Not logged in. Check other settings.</h1>
    }

    const startSaving = async () => {
        setSaving(true)
        try {
            await dispatch(editUser({ name: editableUser!.name ?? '', profileImageUrl: editableUser!.profileImageUrl })).unwrap()
        } catch (e) {
            setSavingError(true)
        }
        setSaving(false)
    }

    const refreshImageUrl = async () => {
        const response = await axios.get(`https://picsum.photos/seed/${Date.now()}/200/200`)
        const newImageURL = response.request.responseURL
        setEditableUser({ ...editableUser!, profileImageUrl: newImageURL })
    }

    const revertImageUrl = async () => {
        setEditableUser({ ...editableUser!, profileImageUrl: currentUser.profileImageUrl })
    }

    return (
        <React.Fragment>
            <h1>Your Account</h1>
            <br />
            {editableUser &&
                <div className="form-container">
                    <div className="field">
                        <label className="label">Profile Image:</label>
                        <ProfileImage imageUrl={editableUser.profileImageUrl} size={ProfileImageSize.MEDIUM} />
                        &nbsp;&nbsp;&nbsp;
                        <button className="button is-small" onClick={refreshImageUrl}>
                            <span className="icon is-small">
                                <i className="fas fa-sync"></i>
                            </span>
                            <span>Change</span>
                        </button>
                        &nbsp;&nbsp;
                        <button className="button is-small" onClick={revertImageUrl}>
                            <span className="icon is-small">
                                <i className="fas fa-undo"></i>
                            </span>
                            <span>Revert</span>
                        </button>
                    </div>
                    <div className="field">
                        <label className="label">Email:</label>
                        <div className="control">
                            <input
                                className="input"
                                type="email"
                                value={editableUser.email}
                                disabled={true}
                                onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setEditableUser({ ...editableUser, email: e.target.value }) }}
                                placeholder="Enter your email" />
                        </div>
                    </div>
                    <div className="field">
                        <label className="label">Name:</label>
                        <div className="control">
                            <input
                                className="input"
                                type="text"
                                value={editableUser.name ?? ''}
                                onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setEditableUser({ ...editableUser, name: e.target.value }) }}
                                placeholder="Enter your name" />
                        </div>
                    </div>
                    {savingError && <span className="help is-danger">There was some problem saving!</span>}
                    <div className="control">
                        {!saving && <button className="button is-primary" onClick={startSaving}>Save</button>}
                        {saving && <button className="button is-primary is-loading">Saving...</button>}
                    </div>
                </div>
            }
        </React.Fragment>
    )
}

export default AccountPage
