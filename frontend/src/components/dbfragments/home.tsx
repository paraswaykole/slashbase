import React from 'react'
import Link from 'next/link'
import Constants from '../../constants'
import { DBConnection, DBDataModel } from '../../data/models'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'

import DBDataModelCard from '../cards/dbdatamodelcard/dbdatamodelcard'

type DBHomePropType = { 
}

const DBHomeFragment = ({}: DBHomePropType) => {

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)

    return (
        <React.Fragment>
            {dbConnection && 
                <React.Fragment>
                    <h1>Showing Data Models in {dbConnection.name}</h1>
                    {dbDataModels.map(x=>(
                        <DBDataModelCard key={x.schemaName+x.name} dataModel={x} dbConnection={dbConnection}/>
                    ))}
                    <Link href={Constants.APP_PATHS.DB_HISTORY.path} as={Constants.APP_PATHS.DB_HISTORY.path.replace('[id]', dbConnection.id)}>
                        <a>
                            <button className="button" >
                                <i className={"fas fa-history"}/>
                                &nbsp;&nbsp;
                                View History
                            </button>
                        </a>
                    </Link>
                </React.Fragment>
            }
        </React.Fragment>
    )
}


export default DBHomeFragment