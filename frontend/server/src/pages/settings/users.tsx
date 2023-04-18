import React, { FunctionComponent, useEffect, useState } from 'react'
import toast from 'react-hot-toast'
import InfiniteScroll from 'react-infinite-scroll-component'
import { Link, useNavigate } from 'react-router-dom'
import UserCard from '../../components/cards/usercard/usercard'
import Constants from '../../constants'
import { User } from '../../data/models'
import apiService from '../../network/apiService'
import { selectCurrentUser } from '../../redux/currentUserSlice'
import { useAppSelector } from '../../redux/hooks'

const UsersPage: FunctionComponent<{}> = () => {

    const currentUser: User = useAppSelector(selectCurrentUser)

    const navigate = useNavigate()

    const [users, setUsers] = useState<User[]>([])
    const [usersNext, setUsersNext] = useState<number>(0)

    useEffect(() => {
        if (currentUser) {
            if (currentUser.isRoot) {
                fetchUsers()
            } else {
                navigate(Constants.APP_PATHS.HOME.path)
            }
        }
    }, [currentUser, usersNext])

    const fetchUsers = async () => {
        if (usersNext === -1) {
            return
        }
        let result = await apiService.getUsers(usersNext)
        if (result.success) {
            if (usersNext === 0) {
                setUsers(result.data.list)
            } else {
                setUsers([...users, ...result.data.list])
            }
            setUsersNext(result.data.next)
        } else {
            toast.error(result.error!)
        }
    }

    return (
        <React.Fragment>
            <h1>Manage Users</h1>
            <InfiniteScroll
                dataLength={users.length}
                next={fetchUsers}
                hasMore={usersNext !== -1}
                loader={
                    <p>Loading...</p>
                }
                scrollableTarget="mainContainer"
            >
                {users.map((user: User) => (
                    <UserCard key={user.id} user={user} />
                ))}
            </InfiniteScroll>
            <Link to={Constants.APP_PATHS.SETTINGS_ADD_USER.path}>
                <button className="button" >
                    <i className={"fas fa-user-plus"} />
                    &nbsp;&nbsp;
                    Add New User
                </button>
            </Link>
        </React.Fragment>
    )
}

export default UsersPage
