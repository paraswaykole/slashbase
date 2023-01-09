import styles from './home.module.scss'
import React from 'react'
import Constants from '../../constants'
import { DBConnection, DBDataModel } from '../../data/models'
import { selectDBConnection, selectDBDataModels, selectIsFetchingDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'

import DBDataModelCard from '../cards/dbdatamodelcard/dbdatamodelcard'
import { Link } from 'react-router-dom'

type DBHomePropType = {
}

const DBHomeFragment = ({ }: DBHomePropType) => {

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)

    const isFetching: boolean = useAppSelector(selectIsFetchingDBDataModels)

    return (
        <React.Fragment>
            {dbConnection &&
                <React.Fragment>
                    <h1>Showing Data Models in {dbConnection.name}</h1>
                    {isFetching && <div className={styles.connectingMsg}>
                        <i className="fas fa-asterisk fa-spin"></i> Connecting to DB...
                    </div>
                    }
                    {dbDataModels.map(x => (
                        <DBDataModelCard key={x.schemaName + x.name} dataModel={x} dbConnection={dbConnection} />
                    ))}
                    <Link to={Constants.APP_PATHS.DB_HISTORY.path.replace('[id]', dbConnection.id)}>
                        <button className="button" >
                            <i className={"fas fa-history"} />
                            &nbsp;&nbsp;
                            View History
                        </button>
                    </Link>
                </React.Fragment>
            }
        </React.Fragment>
    )
}


export default DBHomeFragment