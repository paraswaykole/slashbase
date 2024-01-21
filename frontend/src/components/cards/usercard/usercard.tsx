import styles from './usercard.module.scss'
import { User } from '../../../data/models'
import ProfileImage from '../../user/profileimage'

type UserCardPropType = {
    user: User
}

const UserCard = ({ user }: UserCardPropType) => {

    return (
        <div className={"card " + styles.cardContainer}>
            <div className="card-content">
                <div className="columns is-2">
                    <div className="column">
                        <ProfileImage imageUrl={user.profileImageUrl} />
                    </div>
                    <div className="column is-10">
                        <b>{user.name ?? user.email}</b>
                        {user.name && <b className="subtitle is-6"><br />{user.email}</b>}
                    </div>
                </div>
            </div>
        </div>
    )
}


export default UserCard