import React from 'react'
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
                    <h1>Connected to {dbConnection.name}</h1>
                    {dbDataModels.map(x=>(
                        <DBDataModelCard key={x.schemaName+x.name} dataModel={x} dbConnection={dbConnection}/>
                    ))}
                </React.Fragment>
            }
        </React.Fragment>
    )
}


export default DBHomeFragment