import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import Link from 'next/link'
import React, { useEffect, useState } from 'react'
import toast from 'react-hot-toast'
import InfiniteScroll from 'react-infinite-scroll-component'
import UserCard from '../../components/cards/usercard/usercard'
import AppLayout from '../../components/layouts/applayout'
import Constants from '../../constants'
import { User } from '../../data/models'
import apiService from '../../network/apiService'
import { selectCurrentUser } from '../../redux/currentUserSlice'
import { useAppSelector } from '../../redux/hooks'

const UsersPage: NextPage = () => {

  const currentUser: User = useAppSelector(selectCurrentUser)

  const router = useRouter()

  const [users, setUsers] = useState<User[]>([])
  const [usersNext, setUsersNext] = useState<number>(0) 

  useEffect(()=>{
    if (currentUser) {
      if (currentUser.isRoot) {
        fetchUsers()
      } else {
        router.push(Constants.APP_PATHS.HOME.path)
      }
    }
  }, [currentUser, usersNext])

  const fetchUsers = async () => {
    let result = await apiService.getUsers(usersNext)
    if(result.success) {
      setUsers(result.data.list)
      setUsersNext(result.data.next)
    } else {
      toast.error(result.error!)
    }
  }

  return (
    <AppLayout title="Manage Users | Slashbase">
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
              <UserCard  key={user.id} user={user}/>
            ))}
      </InfiniteScroll>
      <Link href={Constants.APP_PATHS.SETTINGS_ADD_USER.path} as={Constants.APP_PATHS.SETTINGS_ADD_USER.path}>
        <a>
          <button className="button" >
              <i className={"fas fa-user-plus"}/>
              &nbsp;&nbsp;
              Add New User
          </button>
        </a>
      </Link>
    </AppLayout>
  )
}

export default UsersPage
