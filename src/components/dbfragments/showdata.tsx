import styles from './showdata.module.scss'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, DBQueryData } from '../../data/models'
import apiService from '../../network/apiService'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'

type DBShowDataPropType = { 

}

const DBShowDataFragment = (_: DBShowDataPropType) => {

    const router = useRouter()
    const { mschema, mname } = router.query

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    const [dataModel, setDataModel] = useState<DBDataModel>()
    const [queryData, setQueryModel] = useState<DBQueryData>()

    useEffect(()=>{
        const dataModel = dbDataModels.find(x => x.schemaName == mschema && x.name == mname)
        if(dataModel){
            setDataModel(dataModel)
            {(async () => {
                const result = await apiService.getDBDataInDataModel(dbConnection!.id, dataModel?.schemaName ?? '', dataModel.name)
                if (result.success) {
                    setQueryModel(result.data)
                }
            })()}
        }
        // else redirect to home fragment 


            
    }, [dbConnection, dbDataModels, router])

    return (
        <React.Fragment>
            <h1>Showing {dataModel?.schemaName}.{dataModel?.name}</h1>
            <table className={"table "+styles.tableContainer}>
            <thead>
                <tr>
                    {queryData?.columns.map(colName => (<th>{colName}</th>))}    
                </tr>
            </thead>
            <tfoot>
                <tr>
                    {queryData?.columns.map(colName => (<th>{colName}</th>))}
                </tr>
            </tfoot>
            <tbody>
                {queryData?.rows.map(row => {
                    return <tr>
                        { queryData?.columns.map((colName) => {
                                return <td>{row[colName]}</td>
                            })
                        }
                    </tr>
                })}
                
            </tbody>
            </table>
        </React.Fragment>
    )
}


export default DBShowDataFragment