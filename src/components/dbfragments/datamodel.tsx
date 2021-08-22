import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBDataModel } from '../../data/models'
import { selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'

type DBHomePropType = { 

}

const DBDataModelFragment = (_: DBHomePropType) => {

    const router = useRouter()
    const { mschema, mname } = router.query

    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const [dataModel, setDataModel] = useState<DBDataModel>()

    useEffect(()=>{
        const dataModel = dbDataModels.find(x => x.schemaName == mschema && x.name == mname)
        if(dataModel)
            setDataModel(dataModel)
        // else redirect to home fragment 
            
    }, [dbDataModels, router])

    return (
        <React.Fragment>
            <h1>Showing {dataModel?.schemaName}.{dataModel?.name}</h1>
        </React.Fragment>
    )
}


export default DBDataModelFragment