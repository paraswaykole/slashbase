import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import toast from 'react-hot-toast'
import AppLayout from '../../../components/layouts/applayout'
import Constants from '../../../constants'
import { User } from '../../../data/models'
import apiService from '../../../network/apiService'
import { selectCurrentUser } from '../../../redux/currentUserSlice'
import { useAppSelector } from '../../../redux/hooks'

const AddNewUserPage: NextPage = () => {

  const router = useRouter()

  const currentUser: User = useAppSelector(selectCurrentUser)

  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [adding, setAdding] = useState(false)

  useEffect(() => {
    if (currentUser && !currentUser.isRoot) {
      router.push(Constants.APP_PATHS.HOME.path)
    }
  }, [currentUser])

  const startAdding = async () => {
    setAdding(true)
    const result = await apiService.addUser(email, password)
    if (result.success) {
      router.push(Constants.APP_PATHS.SETTINGS_USER.path)
    } else {
      toast.error(result.error!)
    }
    setAdding(false)
  }

  return (
    <AppLayout title="Add New User | Slashbase">
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
    </AppLayout>
  )
}

export default AddNewUserPage
