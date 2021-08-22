import styles from './dbconncard.module.scss'
import React from 'react'
import { DBConnection } from '../../data/models'
import Constants from '../../constants'
import Link from 'next/link'

type DBConnCardPropType = { 
    dbConn: DBConnection
}

const DBConnCard = ({dbConn}: DBConnCardPropType) => {

    return (
        <Link href={Constants.APP_PATHS.DB.href} as={Constants.APP_PATHS.DB.as+dbConn.id}>
            <a>
                <div className={"card "+styles.cardContainer}>
                    <div className="card-content">
                        <h2>{dbConn.name}</h2>
                    </div>
                </div>
            </a>
        </Link>
    )
}


export default DBConnCard