import type { NextPage } from 'next'
import React, { useRef, useState } from 'react'
import AppLayout from '../../../components/layouts/applayout'
import apiService from '../../../network/apiService'

const AccountPage: NextPage = () => {

  const oldPasswordRef = useRef<HTMLInputElement>(null)
  const newPasswordRef = useRef<HTMLInputElement>(null)
  const newPasswordAgainRef = useRef<HTMLInputElement>(null)

  const [saving, setSaving] = useState(false)
  const [savingError, setSavingError] = useState('')
  const [success, setSuccess] = useState(false)

  const startSaving = async () => {
    if (newPasswordRef.current?.value !== newPasswordAgainRef.current?.value) {
      setSavingError('new password is not same as new password again')
      return
    }
    setSaving(true)
    const response = await apiService.changeUserPassword(oldPasswordRef.current!.value!, newPasswordRef.current!.value!)
    if (response.success) {
      oldPasswordRef.current!.value = ''
      newPasswordRef.current!.value = ''
      newPasswordAgainRef.current!.value = ''
      setSavingError('')
      setSuccess(true)
    } else {
      setSuccess(false)
      setSavingError(response.error!)
    }
    setSaving(false)
  }

  return (
    <AppLayout title="Change password | Slashbase">
      <h1>Change your account password</h1>
      <br />
      <div className="form-container">
        <div className="field">
          <label className="label">Old Password:</label>
          <div className="control">
            <input
              ref={oldPasswordRef}
              className="input"
              type="password"
              placeholder="Enter your old password" />
          </div>
        </div>
        <div className="field">
          <label className="label">New Password:</label>
          <div className="control">
            <input
              ref={newPasswordRef}
              className="input"
              type="password"
              placeholder="Enter your new password" />
          </div>
        </div>
        <div className="field">
          <label className="label">New Password again:</label>
          <div className="control">
            <input
              ref={newPasswordAgainRef}
              className="input"
              type="password"
              placeholder="Enter your new password again" />
          </div>
        </div>
        {savingError !== '' && <span className="help is-danger">{savingError}</span>}
        {success && <span className="help is-success">Password saved successfully!</span>}
        <div className="control">
          {!saving && <button className="button is-primary" onClick={startSaving}>Save</button>}
          {saving && <button className="button is-primary is-loading">Saving...</button>}
        </div>
      </div>
    </AppLayout>
  )
}

export default AccountPage
